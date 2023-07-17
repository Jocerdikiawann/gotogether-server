package services

import (
	"context"

	"github.com/Jocerdikiawann/server_share_trip/interceptors"
	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/pb"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RouteServiceServer struct {
	Repo        design.RouteRepository
	Validator   *validator.Validate
	Interceptor *interceptors.AuthInterceptor
	pb.UnimplementedRouteServer
	Logger *logrus.Logger
}

func NewRouteService(repo design.RouteRepository, validator *validator.Validate, logrus *logrus.Logger, interceptor *interceptors.AuthInterceptor) *RouteServiceServer {
	return &RouteServiceServer{
		Repo:        repo,
		Validator:   validator,
		Logger:      logrus,
		Interceptor: interceptor,
	}
}

func (s *RouteServiceServer) GetDestinationAndPolyline(context context.Context, request *pb.RouteRequest) (*pb.DestintationAndPolylineResponse, error) {
	result, err := s.Repo.GetDestinationAndPolyline(context, request.GetId())

	if err != nil {
		s.Logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DestintationAndPolylineResponse{
		StatusCode: int32(codes.OK),
		Success:    true,
		Message:    "success get data.",
		Data: &pb.DestintationAndPolylineType{
			Id:           result.Id.Hex(),
			Destination:  &pb.Point{Latitude: result.DestinationLatLng.Latitude, Longitude: result.DestinationLatLng.Longitude},
			EncodedRoute: result.EncodedRoute,
		},
	}, nil
}

func (s *RouteServiceServer) SendDestinationAndPolyline(context context.Context, req *pb.DestintationAndPolylineRequest) (*pb.DestintationAndPolylineResponse, error) {
	if err := s.Interceptor.Authorize(context); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "err : %v", err)
	}
	structRequest := request.DestinationAndPolylineRequest{
		GoogleId: req.GoogleId,
		Destination: entity.Point{
			Latitude:  req.GetDestination().Latitude,
			Longitude: req.GetDestination().Longitude,
		},
		EncodedRoute: req.EncodedRoute,
	}

	errorValidate := s.Validator.Struct(
		structRequest,
	)

	if errorValidate != nil {
		s.Logger.Error(errorValidate.Error())
		return nil, status.Error(codes.InvalidArgument, errorValidate.Error())
	}

	result, err := s.Repo.SendDestinationAndPolyline(context, structRequest)

	if err != nil {
		s.Logger.Error(err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DestintationAndPolylineResponse{
		StatusCode: int32(codes.OK),
		Success:    true,
		Message:    "success get data.",
		Data: &pb.DestintationAndPolylineType{
			Destination:  req.GetDestination(),
			Id:           result,
			EncodedRoute: req.EncodedRoute,
		},
	}, nil
}
