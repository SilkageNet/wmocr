// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1--rc1
// source: ocr.proto

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

// OCRClient is the client API for OCR service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OCRClient interface {
	Recognize(ctx context.Context, in *RecognizeReq, opts ...grpc.CallOption) (*RecognizeRsp, error)
}

type oCRClient struct {
	cc grpc.ClientConnInterface
}

func NewOCRClient(cc grpc.ClientConnInterface) OCRClient {
	return &oCRClient{cc}
}

func (c *oCRClient) Recognize(ctx context.Context, in *RecognizeReq, opts ...grpc.CallOption) (*RecognizeRsp, error) {
	out := new(RecognizeRsp)
	err := c.cc.Invoke(ctx, "/OCR/Recognize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OCRServer is the server API for OCR service.
// All implementations must embed UnimplementedOCRServer
// for forward compatibility
type OCRServer interface {
	Recognize(context.Context, *RecognizeReq) (*RecognizeRsp, error)
	mustEmbedUnimplementedOCRServer()
}

// UnimplementedOCRServer must be embedded to have forward compatible implementations.
type UnimplementedOCRServer struct {
}

func (UnimplementedOCRServer) Recognize(context.Context, *RecognizeReq) (*RecognizeRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recognize not implemented")
}
func (UnimplementedOCRServer) mustEmbedUnimplementedOCRServer() {}

// UnsafeOCRServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OCRServer will
// result in compilation errors.
type UnsafeOCRServer interface {
	mustEmbedUnimplementedOCRServer()
}

func RegisterOCRServer(s grpc.ServiceRegistrar, srv OCRServer) {
	s.RegisterService(&OCR_ServiceDesc, srv)
}

func _OCR_Recognize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecognizeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCRServer).Recognize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OCR/Recognize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCRServer).Recognize(ctx, req.(*RecognizeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OCR_ServiceDesc is the grpc.ServiceDesc for OCR service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OCR_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OCR",
	HandlerType: (*OCRServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recognize",
			Handler:    _OCR_Recognize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ocr.proto",
}
