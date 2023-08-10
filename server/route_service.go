package server

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/pb"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func getAccessToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return "", status.Error(codes.Unauthenticated, "authorization token is not provided")
	}
	return values[0], nil
}

type RouteServiceServer struct {
	RouteRepo  RouteRepository
	UserRepo   AuthRepository
	JWTManager *utils.JWTManager
	Validator  *validator.Validate
	pb.UnimplementedRouteServer
	Logger *logrus.Logger
}

func NewRouteService(repo RouteRepository, validator *validator.Validate, JWTManager *utils.JWTManager, UserRepo AuthRepository, logrus *logrus.Logger) *RouteServiceServer {
	return &RouteServiceServer{
		RouteRepo:  repo,
		Validator:  validator,
		Logger:     logrus,
		UserRepo:   UserRepo,
		JWTManager: JWTManager,
	}
}

func (s *RouteServiceServer) GetDestinationAndPolyline(context context.Context, request *pb.RouteRequest) (*pb.DestintationAndPolylineResponse, error) {
	result, err := s.RouteRepo.GetDestinationAndPolyline(context, request.GetId())

	if err != nil {
		s.Logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DestintationAndPolylineResponse{
		StatusCode: int32(codes.OK),
		Success:    true,
		Message:    "success get data.",
		Data: &pb.DestintationAndPolylineType{
			Id:              result.Id.Hex(),
			Destination:     &pb.Point{Latitude: result.DestinationLatLng.Latitude, Longitude: result.DestinationLatLng.Longitude},
			EncodedRoute:    result.EncodedRoute,
			InitialLocation: &pb.Point{Latitude: result.InitialLocation.Latitude, Longitude: result.InitialLocation.Longitude},
		},
	}, nil
}

func (s *RouteServiceServer) SendDestinationAndPolyline(context context.Context, req *pb.DestintationAndPolylineRequest) (*pb.DestintationAndPolylineResponse, error) {
	token, err := getAccessToken(context)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	claims, err := s.JWTManager.VerifyAccessToken(token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	errValid := s.UserRepo.CheckIsValidEmail(context, claims.Email)
	if errValid != nil {
		return nil, status.Error(codes.Unauthenticated, errValid.Error())
	}

	structRequest := &request.DestinationAndPolylineRequest{
		GoogleId: claims.GoogleId,
		Destination: &entity.Point{
			Latitude:  req.Destination.Latitude,
			Longitude: req.Destination.Longitude,
		},
		EncodedRoute: req.EncodedRoute,
		InitialLocation: &entity.Point{
			Latitude:  req.InitialLocation.Latitude,
			Longitude: req.InitialLocation.Longitude,
		},
	}
	errorValidate := s.Validator.Struct(structRequest)

	if errorValidate != nil {
		s.Logger.Error(errorValidate.Error())
		return nil, status.Error(codes.InvalidArgument, errorValidate.Error())
	}

	result, err := s.RouteRepo.SendDestinationAndPolyline(context, *structRequest)

	if err != nil {
		s.Logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DestintationAndPolylineResponse{
		StatusCode: int32(codes.OK),
		Success:    true,
		Message:    "success get data.",
		Data: &pb.DestintationAndPolylineType{
			Destination:     req.GetDestination(),
			Id:              result,
			EncodedRoute:    req.EncodedRoute,
			InitialLocation: req.InitialLocation,
		},
	}, nil
}
