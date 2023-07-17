// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: route.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RouteRequest) Reset() {
	*x = RouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteRequest) ProtoMessage() {}

func (x *RouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteRequest.ProtoReflect.Descriptor instead.
func (*RouteRequest) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{0}
}

func (x *RouteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DestintationAndPolylineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GoogleId     string `protobuf:"bytes,1,opt,name=googleId,proto3" json:"googleId,omitempty"`
	Destination  *Point `protobuf:"bytes,2,opt,name=destination,proto3" json:"destination,omitempty"`
	EncodedRoute string `protobuf:"bytes,3,opt,name=encodedRoute,proto3" json:"encodedRoute,omitempty"`
}

func (x *DestintationAndPolylineRequest) Reset() {
	*x = DestintationAndPolylineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestintationAndPolylineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestintationAndPolylineRequest) ProtoMessage() {}

func (x *DestintationAndPolylineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestintationAndPolylineRequest.ProtoReflect.Descriptor instead.
func (*DestintationAndPolylineRequest) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{1}
}

func (x *DestintationAndPolylineRequest) GetGoogleId() string {
	if x != nil {
		return x.GoogleId
	}
	return ""
}

func (x *DestintationAndPolylineRequest) GetDestination() *Point {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *DestintationAndPolylineRequest) GetEncodedRoute() string {
	if x != nil {
		return x.EncodedRoute
	}
	return ""
}

type DestintationAndPolylineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32                        `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Success    bool                         `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Message    string                       `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Data       *DestintationAndPolylineType `protobuf:"bytes,4,opt,name=data,proto3,oneof" json:"data,omitempty"`
}

func (x *DestintationAndPolylineResponse) Reset() {
	*x = DestintationAndPolylineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestintationAndPolylineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestintationAndPolylineResponse) ProtoMessage() {}

func (x *DestintationAndPolylineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestintationAndPolylineResponse.ProtoReflect.Descriptor instead.
func (*DestintationAndPolylineResponse) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{2}
}

func (x *DestintationAndPolylineResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DestintationAndPolylineResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DestintationAndPolylineResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DestintationAndPolylineResponse) GetData() *DestintationAndPolylineType {
	if x != nil {
		return x.Data
	}
	return nil
}

type DestintationAndPolylineType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EncodedRoute string `protobuf:"bytes,2,opt,name=encodedRoute,proto3" json:"encodedRoute,omitempty"`
	Destination  *Point `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
}

func (x *DestintationAndPolylineType) Reset() {
	*x = DestintationAndPolylineType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DestintationAndPolylineType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DestintationAndPolylineType) ProtoMessage() {}

func (x *DestintationAndPolylineType) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DestintationAndPolylineType.ProtoReflect.Descriptor instead.
func (*DestintationAndPolylineType) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{3}
}

func (x *DestintationAndPolylineType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DestintationAndPolylineType) GetEncodedRoute() string {
	if x != nil {
		return x.EncodedRoute
	}
	return ""
}

func (x *DestintationAndPolylineType) GetDestination() *Point {
	if x != nil {
		return x.Destination
	}
	return nil
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_route_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_route_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_route_proto_rawDescGZIP(), []int{4}
}

func (x *Point) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_route_proto protoreflect.FileDescriptor

var file_route_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x1e, 0x0a, 0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x8d, 0x01, 0x0a, 0x1e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x6e, 0x64, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2b,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x22,
	0xb8, 0x01, 0x0a, 0x1f, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x6e, 0x64, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69,
	0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01,
	0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x7e, 0x0a, 0x1b, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x6c,
	0x79, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x2b, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x41, 0x0a, 0x05, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x32, 0x99, 0x02,
	0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x77, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x6c, 0x79,
	0x6c, 0x69, 0x6e, 0x65, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x6c, 0x79, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2f, 0x67, 0x65, 0x74, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x96, 0x01, 0x0a, 0x1a, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x22, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29,
	0x3a, 0x01, 0x2a, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x73,
	0x65, 0x6e, 0x64, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e,
	0x64, 0x50, 0x6f, 0x6c, 0x79, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x6f, 0x63, 0x65, 0x72, 0x64, 0x69, 0x6b,
	0x69, 0x61, 0x77, 0x61, 0x6e, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x70, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_route_proto_rawDescOnce sync.Once
	file_route_proto_rawDescData = file_route_proto_rawDesc
)

func file_route_proto_rawDescGZIP() []byte {
	file_route_proto_rawDescOnce.Do(func() {
		file_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_route_proto_rawDescData)
	})
	return file_route_proto_rawDescData
}

var file_route_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_route_proto_goTypes = []interface{}{
	(*RouteRequest)(nil),                    // 0: pb.RouteRequest
	(*DestintationAndPolylineRequest)(nil),  // 1: pb.DestintationAndPolylineRequest
	(*DestintationAndPolylineResponse)(nil), // 2: pb.DestintationAndPolylineResponse
	(*DestintationAndPolylineType)(nil),     // 3: pb.DestintationAndPolylineType
	(*Point)(nil),                           // 4: pb.Point
}
var file_route_proto_depIdxs = []int32{
	4, // 0: pb.DestintationAndPolylineRequest.destination:type_name -> pb.Point
	3, // 1: pb.DestintationAndPolylineResponse.data:type_name -> pb.DestintationAndPolylineType
	4, // 2: pb.DestintationAndPolylineType.destination:type_name -> pb.Point
	0, // 3: pb.Route.GetDestinationAndPolyline:input_type -> pb.RouteRequest
	1, // 4: pb.Route.SendDestinationAndPolyline:input_type -> pb.DestintationAndPolylineRequest
	2, // 5: pb.Route.GetDestinationAndPolyline:output_type -> pb.DestintationAndPolylineResponse
	2, // 6: pb.Route.SendDestinationAndPolyline:output_type -> pb.DestintationAndPolylineResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_route_proto_init() }
func file_route_proto_init() {
	if File_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestintationAndPolylineRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestintationAndPolylineResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DestintationAndPolylineType); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_route_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_route_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_route_proto_goTypes,
		DependencyIndexes: file_route_proto_depIdxs,
		MessageInfos:      file_route_proto_msgTypes,
	}.Build()
	File_route_proto = out.File
	file_route_proto_rawDesc = nil
	file_route_proto_goTypes = nil
	file_route_proto_depIdxs = nil
}
