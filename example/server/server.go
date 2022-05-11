//go:build wiondows
// +build wiondows

package server

import (
	"context"
	"github.com/SilkageNet/wmocr"
	"github.com/SilkageNet/wmocr/example/server/pb"
	"google.golang.org/grpc"
	"log"
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

func (s *Server) Recognize(_ context.Context, req *pb.RecognizeReq) (*pb.RecognizeRsp, error) {
	var ret, err = s.ocr.Recognize(req.Data)
	if err != nil {
		return nil, err
	}

	log.Println("New recognize result: ", ret)

	return &pb.RecognizeRsp{Ret: ret}, nil
}
