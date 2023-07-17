package services

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/pb"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	Repo       design.AuthRepository
	JWTManager *utils.JWTManager
	Validator  *validator.Validate
	pb.UnimplementedAuthServer
	Logger *logrus.Logger
}

func NewUserService(repo design.AuthRepository, jwtManager *utils.JWTManager, validator *validator.Validate, logrus *logrus.Logger) *UserServiceServer {
	return &UserServiceServer{
		Repo:       repo,
		JWTManager: jwtManager,
		Validator:  validator,
		Logger:     logrus,
	}
}

func (s *UserServiceServer) SignUp(context context.Context, input *pb.UserRequest) (*pb.UserResponse, error) {
	var result entity.Auth
	var authErr error
	requestStruct := request.UserRequest{
		GoogleId: input.GoogleId,
		Email:    input.Email,
		Name:     input.Name,
	}
	validate := s.Validator.Struct(requestStruct)
	if validate != nil {
		s.Logger.Error(validate.Error())
		return nil, status.Errorf(codes.InvalidArgument, "Bad request")
	}

	isUserAlreadyRegister := s.Repo.CheckIsValidEmail(context, input.Email)

	if isUserAlreadyRegister != nil {
		result, authErr = s.Repo.SignUp(context, requestStruct)
	} else {
		result, authErr = s.Repo.UpdateUser(context, requestStruct)
	}

	if authErr != nil {
		s.Logger.Error(authErr.Error())
		return nil, status.Errorf(codes.Internal, authErr.Error())
	}

	token, err := s.JWTManager.Generate(input.Email)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.UserResponse{
		StatusCode: int32(codes.OK),
		Success:    true,
		Message:    "success get data.",
		Data: &pb.UserType{
			Id:       result.Id,
			GoogleId: result.GoogleId,
			Email:    result.Email,
			Name:     result.Name,
		},
		Token: &token,
	}, err
}
