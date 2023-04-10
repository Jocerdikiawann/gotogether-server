package services

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/Jocerdikiawann/server_share_trip/model/entity"
	"github.com/Jocerdikiawann/server_share_trip/model/proto/route"
	"github.com/Jocerdikiawann/server_share_trip/model/request"
	"github.com/Jocerdikiawann/server_share_trip/repository/design"
	"github.com/Jocerdikiawann/server_share_trip/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RouteServiceServer struct {
	Repo design.RouteRepository
	route.UnimplementedRouteServer
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
			// KODE KU TIDAK BERJALAN DI BAGIAN INI
			var data bson.M
			if err := cursor.Decode(&data); err != nil {
				utils.CheckError(err)
			}
			fmt.Println("Document is here2", data)
			fullDocument, _ := data["fullDocument"].(bson.M)
			id := fullDocument["_id"].(primitive.ObjectID).Hex()
			point := fullDocument["point"].(bson.M)
			latitude := point["latitude"].(float64)
			longitude := point["longitude"].(float64)
			dataChan <- &route.LocationResponse{
				Id: id,
				Point: &route.Point{
					Latitude:  latitude,
					Longitude: longitude,
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

func (s *RouteServiceServer) GetDestination(context context.Context, request *route.RouteRequest) (data *route.DestintationAndPolylineResponse, err error) {
	result, err := s.Repo.GetDestinationAndPolyline(context, request.GetId())

	if err != nil {
		err = status.Errorf(codes.Internal, "internal error: %v", err.Error())
	}

	points := make([]*route.Point, 0, len(result.Polyline))

	for _, p := range result.Polyline {
		points = append(points, &route.Point{Latitude: p.Latitude, Longitude: p.Longitude})
	}

	data = &route.DestintationAndPolylineResponse{
		Id:            result.Id.Hex(),
		RoutePolyline: &route.RoutePolyline{Points: points},
		Destination:   &route.Point{Latitude: result.DestinationLatLng.Latitude, Longitude: result.DestinationLatLng.Longitude},
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
			GoogleId: in.GetGoogleId(),
			Point: request.Point{
				Latitude:  in.GetPoint().Latitude,
				Longitude: in.GetPoint().Longitude,
			},
		})

		err = stream.Send(
			&route.LocationResponse{
				Point: in.GetPoint(),
				Id:    id,
			},
		)

		return nil
	}
}

func (s *RouteServiceServer) SendDestinationAndPolyline(context context.Context, req *route.DestintationAndPolylineRequest) (data *route.DestintationAndPolylineResponse, err error) {
	points := make([]entity.Point, 0, len(req.GetRoutePolyline().Points))

	for _, p := range req.GetRoutePolyline().GetPoints() {
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
		return &route.DestintationAndPolylineResponse{}, status.Errorf(codes.Internal, "internal error : %v", err)
	}

	data = &route.DestintationAndPolylineResponse{
		Destination:   req.GetDestination(),
		Id:            result,
		RoutePolyline: req.GetRoutePolyline(),
	}
	return
}
