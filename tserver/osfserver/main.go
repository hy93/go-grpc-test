package main

import (
	"fmt"
	"context"
	pb "go-grpc-test/proto/tserver"
	"ofordcode.ofo.so/baseservice/osf-go.git"
)

type Hi struct{}

func (h Hi) Hello(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	msg := fmt.Sprintf("Hello, %s", req.Name)

	return &pb.Response{
		Message: msg,
	}, nil
	return nil, nil
}

func main() {
	helloService := osf.ServiceConf{
		Name:     "ofo.hello.HY.hiService",
		Register:   pb.RegisterHiServer,
		Server:   Hi{},
		//NoRegistration:true,
	}
	osf.NewServer().Register(helloService).Start(10005)
}
