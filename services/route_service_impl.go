package services

import (
	"context"
	"io"
	"sync"

	"github.com/Jocerdikiawann/server_share_trip/interceptors"
	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/pb"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func NewRouteService(
	repo design.RouteRepository, validator *validator.Validate,
	Interceptor *interceptors.AuthInterceptor, Logger *logrus.Logger,
) *RouteServiceServer {
	return &RouteServiceServer{
		Repo:        repo,
		Validator:   validator,
		Interceptor: Interceptor,
		Logger:      Logger,
	}
}

func (s *RouteServiceServer) WatchLocation(input *pb.WatchRequest, stream pb.Route_WatchLocationServer) error {
	cursor, err := s.Repo.WatchLocation(input.GetGoogleId())
	if err != nil {
		s.Logger.Error(err)
		return status.Errorf(codes.Internal, "failed to watch location: %v", err)
	}

	waitGroup := sync.WaitGroup{}
	dataChan := make(chan *pb.LocationResponse)

	routineCtx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	waitGroup.Add(1)

	go func() {
		defer func() {
			cursor.Close(routineCtx)
			waitGroup.Done()
		}()

		for cursor.Next(routineCtx) {
			var data bson.M
			if err := cursor.Decode(&data); err != nil {
				s.Logger.Error(err)
				continue
			}

			fullDocument, _ := data["fullDocument"].(bson.M)
			id := fullDocument["_id"].(primitive.ObjectID).Hex()
			point := fullDocument["point"].(bson.M)
			latitude := point["latitude"].(float64)
			longitude := point["longitude"].(float64)
			isFinished, _ := point["isFinished"].(bool)

			dataChan <- &pb.LocationResponse{
				StatusCode: int32(codes.OK),
				Success:    true,
				Message:    "success get data",
				Data: &pb.LocationType{
					Id: id,
					Point: &pb.Point{
						Latitude:  latitude,
						Longitude: longitude,
					},
				},
				IsFinish: &isFinished,
			}
		}
	}()

	for {
		select {
		case <-stream.Context().Done():
			cancelFn()
			waitGroup.Done()
			return status.New(codes.OK, "Stream closed.").Err()
		case data, ok := <-dataChan:
			if !ok {
				waitGroup.Done()
				return status.New(codes.OK, "Stream closed.").Err()
			}
			if data.IsFinish != nil && *data.IsFinish {
				cancelFn()
				waitGroup.Done()
				return status.New(codes.OK, "Stream closed.").Err()
			}
			if err := stream.Send(data); err != nil {
				cancelFn()
				waitGroup.Done()
				s.Logger.Error(err)
				return status.New(codes.Internal, "Stream internal error").Err()
			}
		}
	}
}

func (s *RouteServiceServer) GetDestination(context context.Context, request *pb.RouteRequest) (*pb.DestintationAndPolylineResponse, error) {
	result, err := s.Repo.GetDestinationAndPolyline(context, request.GetId())

	if err != nil {
		s.Logger.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	points := make([]*pb.Point, 0, len(result.Polyline))

	for _, p := range result.Polyline {
		points = append(points, &pb.Point{Latitude: p.Latitude, Longitude: p.Longitude})
	}

	return &pb.DestintationAndPolylineResponse{
		StatusCode: int32(codes.OK),
		Success:    true,
		Message:    "success get data.",
		Data: &pb.DestintationAndPolylineType{
			Id:            result.Id.Hex(),
			RoutePolyline: &pb.RoutePolyline{Points: points},
			Destination:   &pb.Point{Latitude: result.DestinationLatLng.Latitude, Longitude: result.DestinationLatLng.Longitude},
		},
	}, nil
}

func (s *RouteServiceServer) SendLocation(stream pb.Route_SendLocationServer) error {
	if err := s.Interceptor.Authorize(stream.Context()); err != nil {
		return status.Errorf(codes.Unauthenticated, "err : %v", err)
	}
	for {
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			s.Logger.Error(err)
			return status.Error(codes.Internal, err.Error())
		}

		structRequest := request.LocationRequest{
			GoogleId: in.GetGoogleId(),
			Point: request.Point{
				Latitude:  in.GetPoint().Latitude,
				Longitude: in.GetPoint().Longitude,
			},
			IsFinish: in.GetIsFinish(),
		}

		id, errData := s.Repo.SendLocation(stream.Context(), structRequest)
		if errData != nil {
			s.Logger.Error(err)
			return status.Error(codes.Internal, errData.Error())
		}

		errSending := stream.Send(&pb.LocationResponse{
			StatusCode: int32(codes.OK),
			Success:    true,
			Message:    "success get data.",
			Data: &pb.LocationType{
				Point: in.GetPoint(),
				Id:    id,
			},
			IsFinish: &in.IsFinish,
		})
		if errSending != nil {
			s.Logger.Error(err)
			return status.Error(codes.Internal, errSending.Error())
		}
	}

	return nil
}

func (s *RouteServiceServer) SendDestinationAndPolyline(context context.Context, req *pb.DestintationAndPolylineRequest) (*pb.DestintationAndPolylineResponse, error) {
	if err := s.Interceptor.Authorize(context); err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "err : %v", err)
	}
	points := make([]entity.Point, 0, len(req.GetRoutePolyline().Points))

	for _, p := range req.GetRoutePolyline().GetPoints() {
		points = append(points, entity.Point{Latitude: p.GetLatitude(), Longitude: p.GetLongitude()})
	}

	structRequest := request.DestinationAndPolylineRequest{
		GoogleId: req.GoogleId,
		Destination: entity.Point{
			Latitude:  req.GetDestination().Latitude,
			Longitude: req.GetDestination().Longitude,
		},
		Polyline: points,
	}

	errorValidate := s.Validator.Struct(
		structRequest,
	)

	if errorValidate != nil {
		s.Logger.Error(errorValidate)
		return nil, status.Error(codes.InvalidArgument, errorValidate.Error())
	}

	result, err := s.Repo.SendDestinationAndPolyline(context, structRequest)

	if err != nil {
		s.Logger.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.DestintationAndPolylineResponse{
		StatusCode: int32(codes.OK),
		Success:    true,
		Message:    "success get data.",
		Data: &pb.DestintationAndPolylineType{
			Destination:   req.GetDestination(),
			Id:            result,
			RoutePolyline: req.GetRoutePolyline(),
		},
	}, nil
}
