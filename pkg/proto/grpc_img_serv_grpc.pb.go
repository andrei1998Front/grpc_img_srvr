// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0--rc2
// source: grpc_img_serv.proto

package grpc_img_server

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

// ImgServiceClient is the client API for ImgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImgServiceClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (ImgService_UploadClient, error)
}

type imgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImgServiceClient(cc grpc.ClientConnInterface) ImgServiceClient {
	return &imgServiceClient{cc}
}

func (c *imgServiceClient) Upload(ctx context.Context, opts ...grpc.CallOption) (ImgService_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &ImgService_ServiceDesc.Streams[0], "/grpc_img_server.ImgService/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &imgServiceUploadClient{stream}
	return x, nil
}

type ImgService_UploadClient interface {
	Send(*ImgUploadRequest) error
	CloseAndRecv() (*ImgInfo, error)
	grpc.ClientStream
}

type imgServiceUploadClient struct {
	grpc.ClientStream
}

func (x *imgServiceUploadClient) Send(m *ImgUploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *imgServiceUploadClient) CloseAndRecv() (*ImgInfo, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ImgInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImgServiceServer is the server API for ImgService service.
// All implementations must embed UnimplementedImgServiceServer
// for forward compatibility
type ImgServiceServer interface {
	Upload(ImgService_UploadServer) error
	mustEmbedUnimplementedImgServiceServer()
}

// UnimplementedImgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImgServiceServer struct {
}

func (UnimplementedImgServiceServer) Upload(ImgService_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedImgServiceServer) mustEmbedUnimplementedImgServiceServer() {}

// UnsafeImgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImgServiceServer will
// result in compilation errors.
type UnsafeImgServiceServer interface {
	mustEmbedUnimplementedImgServiceServer()
}

func RegisterImgServiceServer(s grpc.ServiceRegistrar, srv ImgServiceServer) {
	s.RegisterService(&ImgService_ServiceDesc, srv)
}

func _ImgService_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImgServiceServer).Upload(&imgServiceUploadServer{stream})
}

type ImgService_UploadServer interface {
	SendAndClose(*ImgInfo) error
	Recv() (*ImgUploadRequest, error)
	grpc.ServerStream
}

type imgServiceUploadServer struct {
	grpc.ServerStream
}

func (x *imgServiceUploadServer) SendAndClose(m *ImgInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *imgServiceUploadServer) Recv() (*ImgUploadRequest, error) {
	m := new(ImgUploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ImgService_ServiceDesc is the grpc.ServiceDesc for ImgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_img_server.ImgService",
	HandlerType: (*ImgServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _ImgService_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "grpc_img_serv.proto",
}
