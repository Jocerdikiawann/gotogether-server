package interceptors

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	JWTManager *utils.JWTManager
	Repo       design.AuthRepository
	Logger     *logrus.Logger
}

func NewAuthInterceptor(jwtManager *utils.JWTManager, repo design.AuthRepository, logger *logrus.Logger) *AuthInterceptor {
	return &AuthInterceptor{JWTManager: jwtManager, Repo: repo, Logger: logger}
}

func (interceptor *AuthInterceptor) Authorize(ctx context.Context) (*utils.UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, status.Error(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := interceptor.JWTManager.VerifyAccessToken(accessToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	errValid := interceptor.Repo.CheckIsValidEmail(ctx, claims.Email)
	if errValid != nil {
		return nil, status.Error(codes.Unauthenticated, errValid.Error())
	}
	return claims, nil
}
