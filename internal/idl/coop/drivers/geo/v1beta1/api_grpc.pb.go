// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1beta1

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

// GeoServiceClient is the client API for GeoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GeoServiceClient interface {
	// Convert coordinates to Place ID
	GetPlace(ctx context.Context, in *GetPlaceRequest, opts ...grpc.CallOption) (*GetPlaceResponse, error)
	// Get distance between two places
	GetDistance(ctx context.Context, in *GetDistanceRequest, opts ...grpc.CallOption) (*GetDistanceResponse, error)
}

type geoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGeoServiceClient(cc grpc.ClientConnInterface) GeoServiceClient {
	return &geoServiceClient{cc}
}

func (c *geoServiceClient) GetPlace(ctx context.Context, in *GetPlaceRequest, opts ...grpc.CallOption) (*GetPlaceResponse, error) {
	out := new(GetPlaceResponse)
	err := c.cc.Invoke(ctx, "/coop.drivers.geo.v1beta1.GeoService/GetPlace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *geoServiceClient) GetDistance(ctx context.Context, in *GetDistanceRequest, opts ...grpc.CallOption) (*GetDistanceResponse, error) {
	out := new(GetDistanceResponse)
	err := c.cc.Invoke(ctx, "/coop.drivers.geo.v1beta1.GeoService/GetDistance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GeoServiceServer is the server API for GeoService service.
// All implementations should embed UnimplementedGeoServiceServer
// for forward compatibility
type GeoServiceServer interface {
	// Convert coordinates to Place ID
	GetPlace(context.Context, *GetPlaceRequest) (*GetPlaceResponse, error)
	// Get distance between two places
	GetDistance(context.Context, *GetDistanceRequest) (*GetDistanceResponse, error)
}

// UnimplementedGeoServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGeoServiceServer struct {
}

func (UnimplementedGeoServiceServer) GetPlace(context.Context, *GetPlaceRequest) (*GetPlaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlace not implemented")
}
func (UnimplementedGeoServiceServer) GetDistance(context.Context, *GetDistanceRequest) (*GetDistanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDistance not implemented")
}

// UnsafeGeoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GeoServiceServer will
// result in compilation errors.
type UnsafeGeoServiceServer interface {
	mustEmbedUnimplementedGeoServiceServer()
}

func RegisterGeoServiceServer(s grpc.ServiceRegistrar, srv GeoServiceServer) {
	s.RegisterService(&GeoService_ServiceDesc, srv)
}

func _GeoService_GetPlace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServiceServer).GetPlace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coop.drivers.geo.v1beta1.GeoService/GetPlace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServiceServer).GetPlace(ctx, req.(*GetPlaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GeoService_GetDistance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDistanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoServiceServer).GetDistance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coop.drivers.geo.v1beta1.GeoService/GetDistance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoServiceServer).GetDistance(ctx, req.(*GetDistanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GeoService_ServiceDesc is the grpc.ServiceDesc for GeoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GeoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coop.drivers.geo.v1beta1.GeoService",
	HandlerType: (*GeoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlace",
			Handler:    _GeoService_GetPlace_Handler,
		},
		{
			MethodName: "GetDistance",
			Handler:    _GeoService_GetDistance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/coop/drivers/geo/v1beta1/api.proto",
}