package services

import (
	"context"
	"io"
	"sync"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/proto/route"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RouteServiceServer struct {
	Repo      design.RouteRepository
	Validator *validator.Validate
	route.UnimplementedRouteServer
}

func NewRouteService(repo design.RouteRepository, validator *validator.Validate) *RouteServiceServer {
	return &RouteServiceServer{
		Repo:      repo,
		Validator: validator,
	}
}

func (s *RouteServiceServer) WatchLocation(input *route.WatchRequest, stream route.Route_WatchLocationServer) error {
	cursor, err := s.Repo.WatchLocation(input.GetGoogleId())

	waitGroup := sync.WaitGroup{}
	dataChan := make(chan *route.LocationResponse)

	if err != nil {
		return status.Errorf(codes.Internal, "failed to watch location: %v", err)
	}

	routineCtx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	waitGroup.Add(1)

	go func() {
		defer cursor.Close(routineCtx)
		defer waitGroup.Done()

		for cursor.Next(routineCtx) {
			var data bson.M
			if err := cursor.Decode(&data); err != nil {
				utils.CheckError(err)
			}
			fullDocument, _ := data["fullDocument"].(bson.M)
			id := fullDocument["_id"].(primitive.ObjectID).Hex()
			point := fullDocument["point"].(bson.M)
			latitude := point["latitude"].(float64)
			longitude := point["longitude"].(float64)
			dataChan <- &route.LocationResponse{
				StatusCode: int32(codes.OK),
				Success:    true,
				Message:    "success get data",
				Data: &route.LocationType{
					Id: id,
					Point: &route.Point{
						Latitude:  latitude,
						Longitude: longitude,
					},
				},
			}
		}
	}()

	for {
		select {
		case <-stream.Context().Done():
			defer close(dataChan)
			waitGroup.Wait()
			return nil
		case data := <-dataChan:
			if err := stream.Send(data); err != nil {
				utils.CheckError(err)
			}
		}
	}
}

func (s *RouteServiceServer) GetDestination(context context.Context, request *route.RouteRequest) (*route.DestintationAndPolylineResponse, error) {
	result, err := s.Repo.GetDestinationAndPolyline(context, request.GetId())

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	points := make([]*route.Point, 0, len(result.Polyline))

	for _, p := range result.Polyline {
		points = append(points, &route.Point{Latitude: p.Latitude, Longitude: p.Longitude})
	}

	return &route.DestintationAndPolylineResponse{
		StatusCode: int32(codes.OK),
		Success:    true,
		Message:    "success get data.",
		Data: &route.DestintationAndPolylineType{
			Id:            result.Id.Hex(),
			RoutePolyline: &route.RoutePolyline{Points: points},
			Destination:   &route.Point{Latitude: result.DestinationLatLng.Latitude, Longitude: result.DestinationLatLng.Longitude},
		},
	}, nil
}

func (s *RouteServiceServer) SendLocation(stream route.Route_SendLocationServer) error {
	for {
		in, err := stream.Recv()

		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if err == io.EOF {
			break
		}
		structRequest := request.LocationRequest{
			GoogleId: in.GetGoogleId(),
			Point: request.Point{
				Latitude:  in.GetPoint().Latitude,
				Longitude: in.GetPoint().Longitude,
			},
		}

		errorValidate := s.Validator.Struct(structRequest)
		if errorValidate != nil {
			return status.Error(codes.InvalidArgument, err.Error())
		}

		id, errData := s.Repo.SendLocation(stream.Context(), structRequest)

		if errData != nil {
			return errData
		}

		errSending := stream.Send(
			&route.LocationResponse{
				StatusCode: int32(codes.OK),
				Success:    true,
				Message:    "success get data.",
				Data: &route.LocationType{
					Point: in.GetPoint(),
					Id:    id,
				},
			},
		)

		if errSending != nil {
			return status.Error(codes.Internal, errSending.Error())
		}
	}
	return nil
}

func (s *RouteServiceServer) SendDestinationAndPolyline(context context.Context, req *route.DestintationAndPolylineRequest) (*route.DestintationAndPolylineResponse, error) {
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
		return nil, status.Error(codes.InvalidArgument, errorValidate.Error())
	}

	result, err := s.Repo.SendDestinationAndPolyline(context, structRequest)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &route.DestintationAndPolylineResponse{
		StatusCode: int32(codes.OK),
		Success:    true,
		Message:    "success get data.",
		Data: &route.DestintationAndPolylineType{
			Destination:   req.GetDestination(),
			Id:            result,
			RoutePolyline: req.GetRoutePolyline(),
		},
	}, nil
}
