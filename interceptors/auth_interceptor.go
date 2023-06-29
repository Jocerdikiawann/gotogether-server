package interceptors

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func accessibleRoutes() map[string]bool {
	const path = "/pb.Route"

	return map[string]bool{
		path + "/GetDestination":             false,
		path + "/SendLocation":               true,
		path + "/SendDestinationAndPolyline": true,
		path + "/WatchLocation":              false,
	}
}

type AuthInterceptor struct {
	JWTManager *utils.JWTManager
	Repo       design.AuthRepository
	Logger     *logrus.Logger
}

func NewAuthInterceptor(jwtManager *utils.JWTManager, repo design.AuthRepository, logger *logrus.Logger) *AuthInterceptor {
	return &AuthInterceptor{JWTManager: jwtManager, Repo: repo, Logger: logger}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			interceptor.Logger.Error(err.Error())
			return nil, err
		}
		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		err := interceptor.authorize(ss.Context(), info.FullMethod)
		if err != nil {
			interceptor.Logger.Error(err.Error())
			return err
		}
		return handler(srv, ss)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	accessibleRoutes, ok := accessibleRoutes()[method]

	if !ok || !accessibleRoutes {
		return nil
	}

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
