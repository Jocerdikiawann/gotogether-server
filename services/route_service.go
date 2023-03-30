package services

import (
	"io"

	"github.com/Jocerdikiawann/server_share_trip/model"
)

type RouteServices struct {
	model.UnimplementedShareTripServer
}

func NewService() model.ShareTripServer {
	return &RouteServices{}
}

func (s *RouteServices) GetDestination(
	stream model.ShareTrip_GetDestinationServer,
) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			return err
		}
		err = stream.Send(
			&model.RoutePolyline{
				Points:  in.Points,
				Message: in.Message,
			},
		)
		if err != nil {
			return err
		}
	}
}

func (s *RouteServices) StreamLocation(
	stream model.ShareTrip_StreamLocationServer,
) error {
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		err = stream.Send(&model.RouteTrip{
			Location: point,
			Message:  "Success",
		})
		if err != nil {
			return err
		}
	}
}
