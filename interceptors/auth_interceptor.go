package interceptors

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	JWTManager *utils.JWTManager
	Repo       design.AuthRepository
}

func NewAuthInterceptor(jwtManager *utils.JWTManager, repo design.AuthRepository) *AuthInterceptor {
	return &AuthInterceptor{JWTManager: jwtManager, Repo: repo}
}

func (interceptor *AuthInterceptor) Authorize(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Error(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := interceptor.JWTManager.VerifyAccessToken(accessToken)
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	isValid, errValid := interceptor.Repo.CheckIsValidEmail(ctx, claims.Email)
	if isValid {
		return nil
	}
	if errValid != nil {
		return status.Error(codes.Unauthenticated, errValid.Error())
	}
	return nil
}
