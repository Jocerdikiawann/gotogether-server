// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: route.proto

package pb

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
	Route_GetDestinationAndPolyline_FullMethodName  = "/pb.Route/GetDestinationAndPolyline"
	Route_SendDestinationAndPolyline_FullMethodName = "/pb.Route/SendDestinationAndPolyline"
)

// RouteClient is the client API for Route service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteClient interface {
	GetDestinationAndPolyline(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*DestintationAndPolylineResponse, error)
	SendDestinationAndPolyline(ctx context.Context, in *DestintationAndPolylineRequest, opts ...grpc.CallOption) (*DestintationAndPolylineResponse, error)
}

type routeClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteClient(cc grpc.ClientConnInterface) RouteClient {
	return &routeClient{cc}
}

func (c *routeClient) GetDestinationAndPolyline(ctx context.Context, in *RouteRequest, opts ...grpc.CallOption) (*DestintationAndPolylineResponse, error) {
	out := new(DestintationAndPolylineResponse)
	err := c.cc.Invoke(ctx, Route_GetDestinationAndPolyline_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	GetDestinationAndPolyline(context.Context, *RouteRequest) (*DestintationAndPolylineResponse, error)
	SendDestinationAndPolyline(context.Context, *DestintationAndPolylineRequest) (*DestintationAndPolylineResponse, error)
	mustEmbedUnimplementedRouteServer()
}

// UnimplementedRouteServer must be embedded to have forward compatible implementations.
type UnimplementedRouteServer struct {
}

func (UnimplementedRouteServer) GetDestinationAndPolyline(context.Context, *RouteRequest) (*DestintationAndPolylineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDestinationAndPolyline not implemented")
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
	ServiceName: "pb.Route",
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
	Streams:  []grpc.StreamDesc{},
	Metadata: "route.proto",
}