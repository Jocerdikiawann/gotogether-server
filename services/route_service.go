package services

// type RouteServices struct {
// 	auth.UnimplementedShareTripServer
// }

// func NewService() auth.ShareTripServer {
// 	return &RouteServices{}
// }

// func (s *RouteServices) GetDestination(
// 	stream auth.ShareTrip_GetDestinationServer,
// ) error {
// 	for {
// 		in, err := stream.Recv()
// 		if err != nil {
// 			return err
// 		}
// 		err = stream.Send(
// 			&auth.RoutePolyline{
// 				Points:  in.Points,
// 				Message: in.Message,
// 			},
// 		)
// 		if err != nil {
// 			return err
// 		}
// 	}
// }

// func (s *RouteServices) StreamLocation(
// 	stream model.ShareTrip_StreamLocationServer,
// ) error {
// 	for {
// 		point, err := stream.Recv()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			return err
// 		}

// 		err = stream.Send(&model.RouteTrip{
// 			Location: point,
// 			Message:  "Success",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 	}
// }
