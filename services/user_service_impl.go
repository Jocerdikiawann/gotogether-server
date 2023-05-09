package services

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/model/proto/auth"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	Repo design.AuthRepository
	auth.UnimplementedAuthServer
}

func NewUserService(repo design.AuthRepository) *UserServiceServer {
	return &UserServiceServer{
		Repo: repo,
	}
}

func (s *UserServiceServer) Authentication(context context.Context, input *auth.UserRequest) (data auth.UserResponse, err error) {
	result, err := s.Repo.Authentication(context, request.UserRequest{
		GoogleId: input.GoogleId,
		Email:    input.Email,
		Name:     *input.Name,
	})

	if err != nil {
		err = status.Errorf(codes.Internal, err.Error())
	}

	data = auth.UserResponse{
		Id:       result.Id,
		GoogleId: result.GoogleId,
		Email:    result.Email,
		Name:     &result.Name,
	}
	return
}
