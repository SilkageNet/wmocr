package main

import (
	"context"
	"github.com/SilkageNet/wmocr"
	"github.com/SilkageNet/wmocr/example/pb"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	pb.UnimplementedOCRServer
	ocr *wmocr.WmOcr
}

func NewServer(ocr *wmocr.WmOcr) *Server {
	return &Server{ocr: ocr}
}

func (s *Server) Serve(addr string) error {
	var ln, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	var rpcSrv = grpc.NewServer()
	pb.RegisterOCRServer(rpcSrv, s)
	return rpcSrv.Serve(ln)
}

func (s *Server) Recognize(_ context.Context, req *pb.Request) (*pb.Response, error) {
	var ret, err = s.ocr.Recognize(req.Data)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Ret: ret}, nil
}
