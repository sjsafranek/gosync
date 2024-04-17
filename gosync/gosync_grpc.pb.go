// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: gosync.proto

package gosync

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

// GoSyncServiceClient is the client API for GoSyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoSyncServiceClient interface {
	Authenticate(ctx context.Context, opts ...grpc.CallOption) (GoSyncService_AuthenticateClient, error)
	GetFileDetails(ctx context.Context, in *FilePayload, opts ...grpc.CallOption) (*FilePayload, error)
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (GoSyncService_UploadFileClient, error)
	DownloadFile(ctx context.Context, in *FilePayload, opts ...grpc.CallOption) (GoSyncService_DownloadFileClient, error)
}

type goSyncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoSyncServiceClient(cc grpc.ClientConnInterface) GoSyncServiceClient {
	return &goSyncServiceClient{cc}
}

func (c *goSyncServiceClient) Authenticate(ctx context.Context, opts ...grpc.CallOption) (GoSyncService_AuthenticateClient, error) {
	stream, err := c.cc.NewStream(ctx, &GoSyncService_ServiceDesc.Streams[0], "/GoSyncService/Authenticate", opts...)
	if err != nil {
		return nil, err
	}
	x := &goSyncServiceAuthenticateClient{stream}
	return x, nil
}

type GoSyncService_AuthenticateClient interface {
	Send(*FilePayload) error
	Recv() (*FilePayload, error)
	grpc.ClientStream
}

type goSyncServiceAuthenticateClient struct {
	grpc.ClientStream
}

func (x *goSyncServiceAuthenticateClient) Send(m *FilePayload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *goSyncServiceAuthenticateClient) Recv() (*FilePayload, error) {
	m := new(FilePayload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *goSyncServiceClient) GetFileDetails(ctx context.Context, in *FilePayload, opts ...grpc.CallOption) (*FilePayload, error) {
	out := new(FilePayload)
	err := c.cc.Invoke(ctx, "/GoSyncService/GetFileDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goSyncServiceClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (GoSyncService_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &GoSyncService_ServiceDesc.Streams[1], "/GoSyncService/UploadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &goSyncServiceUploadFileClient{stream}
	return x, nil
}

type GoSyncService_UploadFileClient interface {
	Send(*FilePayload) error
	CloseAndRecv() (*FilePayload, error)
	grpc.ClientStream
}

type goSyncServiceUploadFileClient struct {
	grpc.ClientStream
}

func (x *goSyncServiceUploadFileClient) Send(m *FilePayload) error {
	return x.ClientStream.SendMsg(m)
}

func (x *goSyncServiceUploadFileClient) CloseAndRecv() (*FilePayload, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FilePayload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *goSyncServiceClient) DownloadFile(ctx context.Context, in *FilePayload, opts ...grpc.CallOption) (GoSyncService_DownloadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &GoSyncService_ServiceDesc.Streams[2], "/GoSyncService/DownloadFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &goSyncServiceDownloadFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GoSyncService_DownloadFileClient interface {
	Recv() (*FilePayload, error)
	grpc.ClientStream
}

type goSyncServiceDownloadFileClient struct {
	grpc.ClientStream
}

func (x *goSyncServiceDownloadFileClient) Recv() (*FilePayload, error) {
	m := new(FilePayload)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GoSyncServiceServer is the server API for GoSyncService service.
// All implementations must embed UnimplementedGoSyncServiceServer
// for forward compatibility
type GoSyncServiceServer interface {
	Authenticate(GoSyncService_AuthenticateServer) error
	GetFileDetails(context.Context, *FilePayload) (*FilePayload, error)
	UploadFile(GoSyncService_UploadFileServer) error
	DownloadFile(*FilePayload, GoSyncService_DownloadFileServer) error
	mustEmbedUnimplementedGoSyncServiceServer()
}

// UnimplementedGoSyncServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoSyncServiceServer struct {
}

func (UnimplementedGoSyncServiceServer) Authenticate(GoSyncService_AuthenticateServer) error {
	return status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedGoSyncServiceServer) GetFileDetails(context.Context, *FilePayload) (*FilePayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileDetails not implemented")
}
func (UnimplementedGoSyncServiceServer) UploadFile(GoSyncService_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedGoSyncServiceServer) DownloadFile(*FilePayload, GoSyncService_DownloadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method DownloadFile not implemented")
}
func (UnimplementedGoSyncServiceServer) mustEmbedUnimplementedGoSyncServiceServer() {}

// UnsafeGoSyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoSyncServiceServer will
// result in compilation errors.
type UnsafeGoSyncServiceServer interface {
	mustEmbedUnimplementedGoSyncServiceServer()
}

func RegisterGoSyncServiceServer(s grpc.ServiceRegistrar, srv GoSyncServiceServer) {
	s.RegisterService(&GoSyncService_ServiceDesc, srv)
}

func _GoSyncService_Authenticate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GoSyncServiceServer).Authenticate(&goSyncServiceAuthenticateServer{stream})
}

type GoSyncService_AuthenticateServer interface {
	Send(*FilePayload) error
	Recv() (*FilePayload, error)
	grpc.ServerStream
}

type goSyncServiceAuthenticateServer struct {
	grpc.ServerStream
}

func (x *goSyncServiceAuthenticateServer) Send(m *FilePayload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *goSyncServiceAuthenticateServer) Recv() (*FilePayload, error) {
	m := new(FilePayload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GoSyncService_GetFileDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilePayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoSyncServiceServer).GetFileDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoSyncService/GetFileDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoSyncServiceServer).GetFileDetails(ctx, req.(*FilePayload))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoSyncService_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GoSyncServiceServer).UploadFile(&goSyncServiceUploadFileServer{stream})
}

type GoSyncService_UploadFileServer interface {
	SendAndClose(*FilePayload) error
	Recv() (*FilePayload, error)
	grpc.ServerStream
}

type goSyncServiceUploadFileServer struct {
	grpc.ServerStream
}

func (x *goSyncServiceUploadFileServer) SendAndClose(m *FilePayload) error {
	return x.ServerStream.SendMsg(m)
}

func (x *goSyncServiceUploadFileServer) Recv() (*FilePayload, error) {
	m := new(FilePayload)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GoSyncService_DownloadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FilePayload)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GoSyncServiceServer).DownloadFile(m, &goSyncServiceDownloadFileServer{stream})
}

type GoSyncService_DownloadFileServer interface {
	Send(*FilePayload) error
	grpc.ServerStream
}

type goSyncServiceDownloadFileServer struct {
	grpc.ServerStream
}

func (x *goSyncServiceDownloadFileServer) Send(m *FilePayload) error {
	return x.ServerStream.SendMsg(m)
}

// GoSyncService_ServiceDesc is the grpc.ServiceDesc for GoSyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoSyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GoSyncService",
	HandlerType: (*GoSyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFileDetails",
			Handler:    _GoSyncService_GetFileDetails_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Authenticate",
			Handler:       _GoSyncService_Authenticate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "UploadFile",
			Handler:       _GoSyncService_UploadFile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DownloadFile",
			Handler:       _GoSyncService_DownloadFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "gosync.proto",
}
