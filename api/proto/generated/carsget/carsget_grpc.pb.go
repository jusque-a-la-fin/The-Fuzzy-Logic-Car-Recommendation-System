// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: carsget.proto

package carsget

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Cars_Select_FullMethodName = "/carsget.Cars/Select"
	Cars_Search_FullMethodName = "/carsget.Cars/Search"
)

// CarsClient is the client API for Cars service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarsClient interface {
	// Select собирает информацию об автомобилях из базы данных для сервиса,
	// осуществляющего подбор автомобилей с использованием нечеткого алгоритма
	Select(ctx context.Context, in *CarsSelectionRequest, opts ...grpc.CallOption) (*CarsSelectionResponse, error)
	// Search собирает информацию об автомобилях из базы данных для сервиса,
	// осуществляющего поиск автомобилей с помощью фильтров
	Search(ctx context.Context, in *CarsSearchRequest, opts ...grpc.CallOption) (*CarsSearchResponse, error)
}

type carsClient struct {
	cc grpc.ClientConnInterface
}

func NewCarsClient(cc grpc.ClientConnInterface) CarsClient {
	return &carsClient{cc}
}

func (c *carsClient) Select(ctx context.Context, in *CarsSelectionRequest, opts ...grpc.CallOption) (*CarsSelectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CarsSelectionResponse)
	err := c.cc.Invoke(ctx, Cars_Select_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carsClient) Search(ctx context.Context, in *CarsSearchRequest, opts ...grpc.CallOption) (*CarsSearchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CarsSearchResponse)
	err := c.cc.Invoke(ctx, Cars_Search_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarsServer is the server API for Cars service.
// All implementations must embed UnimplementedCarsServer
// for forward compatibility.
type CarsServer interface {
	// Select собирает информацию об автомобилях из базы данных для сервиса,
	// осуществляющего подбор автомобилей с использованием нечеткого алгоритма
	Select(context.Context, *CarsSelectionRequest) (*CarsSelectionResponse, error)
	// Search собирает информацию об автомобилях из базы данных для сервиса,
	// осуществляющего поиск автомобилей с помощью фильтров
	Search(context.Context, *CarsSearchRequest) (*CarsSearchResponse, error)
	mustEmbedUnimplementedCarsServer()
}

// UnimplementedCarsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCarsServer struct{}

func (UnimplementedCarsServer) Select(context.Context, *CarsSelectionRequest) (*CarsSelectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Select not implemented")
}
func (UnimplementedCarsServer) Search(context.Context, *CarsSearchRequest) (*CarsSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedCarsServer) mustEmbedUnimplementedCarsServer() {}
func (UnimplementedCarsServer) testEmbeddedByValue()              {}

// UnsafeCarsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarsServer will
// result in compilation errors.
type UnsafeCarsServer interface {
	mustEmbedUnimplementedCarsServer()
}

func RegisterCarsServer(s grpc.ServiceRegistrar, srv CarsServer) {
	// If the following call pancis, it indicates UnimplementedCarsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Cars_ServiceDesc, srv)
}

func _Cars_Select_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CarsSelectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarsServer).Select(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cars_Select_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarsServer).Select(ctx, req.(*CarsSelectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cars_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CarsSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarsServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Cars_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarsServer).Search(ctx, req.(*CarsSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cars_ServiceDesc is the grpc.ServiceDesc for Cars service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cars_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "carsget.Cars",
	HandlerType: (*CarsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Select",
			Handler:    _Cars_Select_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _Cars_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "carsget.proto",
}
