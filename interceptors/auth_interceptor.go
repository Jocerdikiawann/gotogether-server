package interceptors

import (
	"context"
	"log"

	"github.com/Jocerdikiawann/server_share_trip/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func accessibleRoutes() map[string]bool {
	const path = "/app.sharetrip.route.Route/"

	return map[string]bool{
		path + "GetDestination":             true,
		path + "SendLocation":               true,
		path + "SendDestinationAndPolyline": true,
		path + "WatchLocation":              false,
	}
}

type AuthInterceptor struct {
	*utils.JWTManager
}

func NewAuthInterceptor(jwtManager *utils.JWTManager) *AuthInterceptor {
	return &AuthInterceptor{JWTManager: jwtManager}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Printf("--> Unary interceptors : %v", info.FullMethod)
		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
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
		log.Printf("--> Stream interceptors : %v", info.FullMethod)
		err := interceptor.authorize(ss.Context(), info.FullMethod)
		if err != nil {
			return err
		}
		return handler(srv, ss)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {
	_, ok := accessibleRoutes()[method]

	if !ok {
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
	//TODO: claims email check on db
	_, err := interceptor.JWTManager.VerifyAccessToken(accessToken)
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}
	return nil
}
