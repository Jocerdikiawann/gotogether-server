package services

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/model/proto/auth"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	Repo       design.AuthRepository
	JWTManager *utils.JWTManager
	auth.UnimplementedAuthServer
}

func NewUserService(repo design.AuthRepository, jwtManager *utils.JWTManager) *UserServiceServer {
	return &UserServiceServer{
		Repo:       repo,
		JWTManager: jwtManager,
	}
}

func (s *UserServiceServer) Authentication(context context.Context, input *auth.UserRequest) (*auth.UserResponse, error) {
	result, err := s.Repo.Authentication(context, request.UserRequest{
		GoogleId: input.GoogleId,
		Email:    input.Email,
		Name:     input.Name,
	})

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
