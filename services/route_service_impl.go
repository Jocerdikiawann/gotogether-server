package services

import (
	"context"
	"io"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/proto/route"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RouteServiceServer struct {
	Repo design.RouteRepository
	route.UnimplementedRouteServer
}

func (s *RouteServiceServer) WatchLocation(input *route.RouteRequest, stream route.Route_WatchLocationServer) error {
	cursor, err := s.Repo.WatchLocation(input.GetId())

	if err != nil {
		return status.Errorf(codes.Internal, "failed to watch location: %v", err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var data entity.Destination
		if err := cursor.Decode(&data); err != nil {
			return status.Errorf(codes.InvalidArgument, "failed decode data: %v", err)
		}
		err := stream.Send(&route.LocationResponse{
			Point: &route.Point{
				Latitude:  data.DestinationLatLng.Latitude,
				Longitude: data.DestinationLatLng.Longitude,
			},
		})
		if err != nil {
			return status.Errorf(codes.Internal, "failed to send data: %v", err)
		}
	}

	return err
}

func (s *RouteServiceServer) GetDestination(context context.Context, request *route.RouteRequest) (data *route.DestintationAndPolylineResponse, err error) {
	result, err := s.Repo.GetDestinationAndPolyline(context, request.GetId())

	if err != nil {
		err = status.Errorf(codes.Internal, "internal error: %v", err.Error())
	}

	points := make([]*route.Point, 0, len(result.Polyline))

	defer func() {
		for i := range points {
			points[i] = nil
		}
	}()

	for _, p := range result.Polyline {
		points = append(points, &route.Point{Latitude: p.Latitude, Longitude: p.Longitude})
	}

	data = &route.DestintationAndPolylineResponse{
		Data:        &route.RoutePolyline{Points: points},
		Destination: &route.Point{Latitude: result.DestinationLatLng.Latitude, Longitude: result.DestinationLatLng.Longitude},
	}

	return
}

func (s *RouteServiceServer) SendLocation(stream route.Route_SendLocationServer) error {
	for {
		in, err := stream.Recv()

		if err != nil {
			return status.Errorf(codes.Internal, "failed")
		}

		if err == io.EOF {
			return nil
		}

		id, err := s.Repo.SendLocation(stream.Context(), request.LocationRequest{
			GoogleId:  in.GetGoogleId(),
			Latitude:  in.GetPoint().Latitude,
			Longitude: in.GetPoint().Longitude,
		})

		err = stream.Send(
			&route.LocationResponse{
				Point: in.GetPoint(),
				Id:    &id,
			},
		)

		return nil
	}
}

func (s *RouteServiceServer) SendDestinationAndPolyline(context context.Context, req *route.DestintationAndPolylineRequest) (data *route.DestintationAndPolylineResponse, err error) {
	points := make([]entity.Point, 0, len(req.GetPolyline().Points))

	for _, p := range req.GetPolyline().GetPoints() {
		points = append(points, entity.Point{Latitude: p.GetLatitude(), Longitude: p.GetLongitude()})
	}

	result, err := s.Repo.SendDestinationAndPolyline(context, request.DestinationAndPolylineRequest{
		GoogleId: req.GoogleId,
		Destination: entity.Point{
			Latitude:  req.GetDestination().Latitude,
			Longitude: req.GetDestination().Longitude,
		},
		Polyline: points,
	})

	if err != nil {
		err = status.Errorf(codes.Internal, "internal error : %v", err)
	}

	data = &route.DestintationAndPolylineResponse{
		Data:        req.GetPolyline(),
		Destination: req.GetDestination(),
		Id:          result,
	}
	return
}
