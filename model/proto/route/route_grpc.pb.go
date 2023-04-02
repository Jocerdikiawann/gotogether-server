// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: model/proto/route/route.proto

package route

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Route_GetLocation_FullMethodName                = "/app.sharetrip.route.Route/GetLocation"
	Route_GetDestinationAndPolyline_FullMethodName  = "/app.sharetrip.route.Route/GetDestinationAndPolyline"
	Route_SendLocation_FullMethodName               = "/app.sharetrip.route.Route/SendLocation"
	Route_SendDestinationAndPolyline_FullMethodName = "/app.sharetrip.route.Route/SendDestinationAndPolyline"
)

// RouteClient is the client API for Route service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteClient interface {
	GetLocation(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (Route_GetLocationClient, error)
	GetDestinationAndPolyline(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*DestintationAndPolylineResponse, error)
	SendLocation(ctx context.Context, opts ...grpc.CallOption) (Route_SendLocationClient, error)
	SendDestinationAndPolyline(ctx context.Context, in *DestintationAndPolylineRequest, opts ...grpc.CallOption) (*DestintationAndPolylineResponse, error)
}

type routeClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteClient(cc grpc.ClientConnInterface) RouteClient {
	return &routeClient{cc}
}

func (c *routeClient) GetLocation(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (Route_GetLocationClient, error) {
	stream, err := c.cc.NewStream(ctx, &Route_ServiceDesc.Streams[0], Route_GetLocation_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGetLocationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Route_GetLocationClient interface {
	Recv() (*LocationResponse, error)
	grpc.ClientStream
}

type routeGetLocationClient struct {
	grpc.ClientStream
}

func (x *routeGetLocationClient) Recv() (*LocationResponse, error) {
	m := new(LocationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeClient) GetDestinationAndPolyline(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*DestintationAndPolylineResponse, error) {
	out := new(DestintationAndPolylineResponse)
	err := c.cc.Invoke(ctx, Route_GetDestinationAndPolyline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routeClient) SendLocation(ctx context.Context, opts ...grpc.CallOption) (Route_SendLocationClient, error) {
	stream, err := c.cc.NewStream(ctx, &Route_ServiceDesc.Streams[1], Route_SendLocation_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &routeSendLocationClient{stream}
	return x, nil
}

type Route_SendLocationClient interface {
	Send(*LocationRequest) error
	Recv() (*LocationResponse, error)
	grpc.ClientStream
}

type routeSendLocationClient struct {
	grpc.ClientStream
}

func (x *routeSendLocationClient) Send(m *LocationRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeSendLocationClient) Recv() (*LocationResponse, error) {
	m := new(LocationResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeClient) SendDestinationAndPolyline(ctx context.Context, in *DestintationAndPolylineRequest, opts ...grpc.CallOption) (*DestintationAndPolylineResponse, error) {
	out := new(DestintationAndPolylineResponse)
	err := c.cc.Invoke(ctx, Route_SendDestinationAndPolyline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteServer is the server API for Route service.
// All implementations must embed UnimplementedRouteServer
// for forward compatibility
type RouteServer interface {
	GetLocation(*RouteRequest, Route_GetLocationServer) error
	GetDestinationAndPolyline(context.Context, *RouteRequest) (*DestintationAndPolylineResponse, error)
	SendLocation(Route_SendLocationServer) error
	SendDestinationAndPolyline(context.Context, *DestintationAndPolylineRequest) (*DestintationAndPolylineResponse, error)
	mustEmbedUnimplementedRouteServer()
}

// UnimplementedRouteServer must be embedded to have forward compatible implementations.
type UnimplementedRouteServer struct {
}

func (UnimplementedRouteServer) GetLocation(*RouteRequest, Route_GetLocationServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (UnimplementedRouteServer) GetDestinationAndPolyline(context.Context, *RouteRequest) (*DestintationAndPolylineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDestinationAndPolyline not implemented")
}
func (UnimplementedRouteServer) SendLocation(Route_SendLocationServer) error {
	return status.Errorf(codes.Unimplemented, "method SendLocation not implemented")
}
func (UnimplementedRouteServer) SendDestinationAndPolyline(context.Context, *DestintationAndPolylineRequest) (*DestintationAndPolylineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDestinationAndPolyline not implemented")
}
func (UnimplementedRouteServer) mustEmbedUnimplementedRouteServer() {}

// UnsafeRouteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteServer will
// result in compilation errors.
type UnsafeRouteServer interface {
	mustEmbedUnimplementedRouteServer()
}

func RegisterRouteServer(s grpc.ServiceRegistrar, srv RouteServer) {
	s.RegisterService(&Route_ServiceDesc, srv)
}

func _Route_GetLocation_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RouteRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteServer).GetLocation(m, &routeGetLocationServer{stream})
}

type Route_GetLocationServer interface {
	Send(*LocationResponse) error
	grpc.ServerStream
}

type routeGetLocationServer struct {
	grpc.ServerStream
}

func (x *routeGetLocationServer) Send(m *LocationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Route_GetDestinationAndPolyline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServer).GetDestinationAndPolyline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route_GetDestinationAndPolyline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServer).GetDestinationAndPolyline(ctx, req.(*RouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Route_SendLocation_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteServer).SendLocation(&routeSendLocationServer{stream})
}

type Route_SendLocationServer interface {
	Send(*LocationResponse) error
	Recv() (*LocationRequest, error)
	grpc.ServerStream
}

type routeSendLocationServer struct {
	grpc.ServerStream
}

func (x *routeSendLocationServer) Send(m *LocationResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeSendLocationServer) Recv() (*LocationRequest, error) {
	m := new(LocationRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Route_SendDestinationAndPolyline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestintationAndPolylineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServer).SendDestinationAndPolyline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Route_SendDestinationAndPolyline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServer).SendDestinationAndPolyline(ctx, req.(*DestintationAndPolylineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Route_ServiceDesc is the grpc.ServiceDesc for Route service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Route_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.sharetrip.route.Route",
	HandlerType: (*RouteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDestinationAndPolyline",
			Handler:    _Route_GetDestinationAndPolyline_Handler,
		},
		{
			MethodName: "SendDestinationAndPolyline",
			Handler:    _Route_SendDestinationAndPolyline_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLocation",
			Handler:       _Route_GetLocation_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendLocation",
			Handler:       _Route_SendLocation_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "model/proto/route/route.proto",
}
