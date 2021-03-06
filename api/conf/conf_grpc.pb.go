// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package conf

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

// ConfClient is the client API for Conf service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConfClient interface {
	GetConf(ctx context.Context, in *GetConfReq, opts ...grpc.CallOption) (*GetConfReply, error)
}

type confClient struct {
	cc grpc.ClientConnInterface
}

func NewConfClient(cc grpc.ClientConnInterface) ConfClient {
	return &confClient{cc}
}

func (c *confClient) GetConf(ctx context.Context, in *GetConfReq, opts ...grpc.CallOption) (*GetConfReply, error) {
	out := new(GetConfReply)
	err := c.cc.Invoke(ctx, "/Conf/GetConf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfServer is the server API for Conf service.
// All implementations must embed UnimplementedConfServer
// for forward compatibility
type ConfServer interface {
	GetConf(context.Context, *GetConfReq) (*GetConfReply, error)
	mustEmbedUnimplementedConfServer()
}

// UnimplementedConfServer must be embedded to have forward compatible implementations.
type UnimplementedConfServer struct {
}

func (UnimplementedConfServer) GetConf(context.Context, *GetConfReq) (*GetConfReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConf not implemented")
}
func (UnimplementedConfServer) mustEmbedUnimplementedConfServer() {}

// UnsafeConfServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfServer will
// result in compilation errors.
type UnsafeConfServer interface {
	mustEmbedUnimplementedConfServer()
}

func RegisterConfServer(s grpc.ServiceRegistrar, srv ConfServer) {
	s.RegisterService(&Conf_ServiceDesc, srv)
}

func _Conf_GetConf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfServer).GetConf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Conf/GetConf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfServer).GetConf(ctx, req.(*GetConfReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Conf_ServiceDesc is the grpc.ServiceDesc for Conf service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Conf_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Conf",
	HandlerType: (*ConfServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConf",
			Handler:    _Conf_GetConf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/conf/conf.proto",
}
