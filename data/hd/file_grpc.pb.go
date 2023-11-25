// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: file.proto

package hd_ticket

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

// FileServiceClient is the client API for FileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileServiceClient interface {
	GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error)
	UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*UpdateFileResponse, error)
	CreateFileTable(ctx context.Context, in *CreateFileTableRequest, opts ...grpc.CallOption) (*CreateFileTableResponse, error)
	CreateNewFile(ctx context.Context, in *CreateNewFileRequest, opts ...grpc.CallOption) (*CreateNewFileResponse, error)
}

type fileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileServiceClient(cc grpc.ClientConnInterface) FileServiceClient {
	return &fileServiceClient{cc}
}

func (c *fileServiceClient) GetFile(ctx context.Context, in *GetFileRequest, opts ...grpc.CallOption) (*GetFileResponse, error) {
	out := new(GetFileResponse)
	err := c.cc.Invoke(ctx, "/FileService/GetFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) UpdateFile(ctx context.Context, in *UpdateFileRequest, opts ...grpc.CallOption) (*UpdateFileResponse, error) {
	out := new(UpdateFileResponse)
	err := c.cc.Invoke(ctx, "/FileService/UpdateFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) CreateFileTable(ctx context.Context, in *CreateFileTableRequest, opts ...grpc.CallOption) (*CreateFileTableResponse, error) {
	out := new(CreateFileTableResponse)
	err := c.cc.Invoke(ctx, "/FileService/CreateFileTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileServiceClient) CreateNewFile(ctx context.Context, in *CreateNewFileRequest, opts ...grpc.CallOption) (*CreateNewFileResponse, error) {
	out := new(CreateNewFileResponse)
	err := c.cc.Invoke(ctx, "/FileService/CreateNewFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServiceServer is the server API for FileService service.
// All implementations must embed UnimplementedFileServiceServer
// for forward compatibility
type FileServiceServer interface {
	GetFile(context.Context, *GetFileRequest) (*GetFileResponse, error)
	UpdateFile(context.Context, *UpdateFileRequest) (*UpdateFileResponse, error)
	CreateFileTable(context.Context, *CreateFileTableRequest) (*CreateFileTableResponse, error)
	CreateNewFile(context.Context, *CreateNewFileRequest) (*CreateNewFileResponse, error)
	mustEmbedUnimplementedFileServiceServer()
}

// UnimplementedFileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileServiceServer struct {
}

func (UnimplementedFileServiceServer) GetFile(context.Context, *GetFileRequest) (*GetFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (UnimplementedFileServiceServer) UpdateFile(context.Context, *UpdateFileRequest) (*UpdateFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFile not implemented")
}
func (UnimplementedFileServiceServer) CreateFileTable(context.Context, *CreateFileTableRequest) (*CreateFileTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFileTable not implemented")
}
func (UnimplementedFileServiceServer) CreateNewFile(context.Context, *CreateNewFileRequest) (*CreateNewFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewFile not implemented")
}
func (UnimplementedFileServiceServer) mustEmbedUnimplementedFileServiceServer() {}

// UnsafeFileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServiceServer will
// result in compilation errors.
type UnsafeFileServiceServer interface {
	mustEmbedUnimplementedFileServiceServer()
}

func RegisterFileServiceServer(s grpc.ServiceRegistrar, srv FileServiceServer) {
	s.RegisterService(&FileService_ServiceDesc, srv)
}

func _FileService_GetFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).GetFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/GetFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).GetFile(ctx, req.(*GetFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_UpdateFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).UpdateFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/UpdateFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).UpdateFile(ctx, req.(*UpdateFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_CreateFileTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFileTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).CreateFileTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/CreateFileTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).CreateFileTable(ctx, req.(*CreateFileTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileService_CreateNewFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServiceServer).CreateNewFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FileService/CreateNewFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServiceServer).CreateNewFile(ctx, req.(*CreateNewFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FileService_ServiceDesc is the grpc.ServiceDesc for FileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FileService",
	HandlerType: (*FileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFile",
			Handler:    _FileService_GetFile_Handler,
		},
		{
			MethodName: "UpdateFile",
			Handler:    _FileService_UpdateFile_Handler,
		},
		{
			MethodName: "CreateFileTable",
			Handler:    _FileService_CreateFileTable_Handler,
		},
		{
			MethodName: "CreateNewFile",
			Handler:    _FileService_CreateNewFile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "file.proto",
}
