package services

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"github.com/Jocerdikiawann/shared_proto_share_trip/auth"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	Repo       design.AuthRepository
	JWTManager *utils.JWTManager
	Validator  *validator.Validate
	auth.UnimplementedAuthServer
}

func NewUserService(repo design.AuthRepository, jwtManager *utils.JWTManager, validator *validator.Validate) *UserServiceServer {
	return &UserServiceServer{
		Repo:       repo,
		JWTManager: jwtManager,
		Validator:  validator,
	}
}

func (s *UserServiceServer) SignUp(context context.Context, input *auth.UserRequest) (*auth.UserResponse, error) {
	requestStruct := request.UserRequest{
		GoogleId: input.GoogleId,
		Email:    input.Email,
		Name:     input.Name,
	}
	validate := s.Validator.Struct(requestStruct)
	if validate != nil {
		return nil, status.Errorf(codes.InvalidArgument, validate.Error())
	}
	result, err := s.Repo.SignUp(context, requestStruct)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	token, err := s.JWTManager.Generate(input.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &auth.UserResponse{
		StatusCode: int32(codes.OK),
		Success:    true,
		Message:    "success get data.",
		Data: &auth.UserType{
			Id:       result.Id,
			GoogleId: result.GoogleId,
			Email:    result.Email,
			Name:     result.Name,
		},
		Token: &token,
	}, err
}
