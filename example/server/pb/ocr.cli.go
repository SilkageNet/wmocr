package pb

import (
	"fmt"
	"google.golang.org/grpc"
)

func NewOCRCli(addr string) (OCRClient, error) {
	var cc, err = grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("grpc.dial.err:%s", err.Error())
	}
	return &oCRClient{cc}, nil
}
