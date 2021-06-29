// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// FileClient is the client API for File service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileClient interface {
	AddFile(ctx context.Context, in *AddFileReq, opts ...grpc.CallOption) (*AddFileReply, error)
	GetFile(ctx context.Context, in *GetFileReq, opts ...grpc.CallOption) (*GetFileReply, error)
	ListFile(ctx context.Context, in *ListFileReq, opts ...grpc.CallOption) (*ListFileReply, error)
	DeleteFile(ctx context.Context, in *DeleteFileReq, opts ...grpc.CallOption) (*DeleteFileReply, error)
}

type fileClient struct {
	cc grpc.ClientConnInterface
}

func NewFileClient(cc grpc.ClientConnInterface) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) AddFile(ctx context.Context, in *AddFileReq, opts ...grpc.CallOption) (*AddFileReply, error) {
	out := new(AddFileReply)
	err := c.cc.Invoke(ctx, "/File/AddFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) GetFile(ctx context.Context, in *GetFileReq, opts ...grpc.CallOption) (*GetFileReply, error) {
	out := new(GetFileReply)
	err := c.cc.Invoke(ctx, "/File/GetFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) ListFile(ctx context.Context, in *ListFileReq, opts ...grpc.CallOption) (*ListFileReply, error) {
	out := new(ListFileReply)
	err := c.cc.Invoke(ctx, "/File/ListFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) DeleteFile(ctx context.Context, in *DeleteFileReq, opts ...grpc.CallOption) (*DeleteFileReply, error) {
	out := new(DeleteFileReply)
	err := c.cc.Invoke(ctx, "/File/DeleteFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServer is the server API for File service.
// All implementations must embed UnimplementedFileServer
// for forward compatibility
type FileServer interface {
	AddFile(context.Context, *AddFileReq) (*AddFileReply, error)
	GetFile(context.Context, *GetFileReq) (*GetFileReply, error)
	ListFile(context.Context, *ListFileReq) (*ListFileReply, error)
	DeleteFile(context.Context, *DeleteFileReq) (*DeleteFileReply, error)
	mustEmbedUnimplementedFileServer()
}

// UnimplementedFileServer must be embedded to have forward compatible implementations.
type UnimplementedFileServer struct {
}

func (UnimplementedFileServer) AddFile(context.Context, *AddFileReq) (*AddFileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFile not implemented")
}
func (UnimplementedFileServer) GetFile(context.Context, *GetFileReq) (*GetFileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedFileServer) ListFile(context.Context, *ListFileReq) (*ListFileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFile not implemented")
}
func (UnimplementedFileServer) DeleteFile(context.Context, *DeleteFileReq) (*DeleteFileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFile not implemented")
}
func (UnimplementedFileServer) mustEmbedUnimplementedFileServer() {}

// UnsafeFileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServer will
// result in compilation errors.
type UnsafeFileServer interface {
	mustEmbedUnimplementedFileServer()
}

func RegisterFileServer(s grpc.ServiceRegistrar, srv FileServer) {
	s.RegisterService(&File_ServiceDesc, srv)
}

func _File_AddFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).AddFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/File/AddFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).AddFile(ctx, req.(*AddFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/File/GetFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).GetFile(ctx, req.(*GetFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_ListFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).ListFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/File/ListFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).ListFile(ctx, req.(*ListFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_DeleteFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).DeleteFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/File/DeleteFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).DeleteFile(ctx, req.(*DeleteFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

// File_ServiceDesc is the grpc.ServiceDesc for File service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var File_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "File",
	HandlerType: (*FileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFile",
			Handler:    _File_AddFile_Handler,
		},
		{
			MethodName: "GetFile",
			Handler:    _File_GetFile_Handler,
		},
		{
			MethodName: "ListFile",
			Handler:    _File_ListFile_Handler,
		},
		{
			MethodName: "DeleteFile",
			Handler:    _File_DeleteFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/file/service/v1/file.proto",
}