package main

import (
	pb "go-grpc-test/proto/tserver"
	"context"
	"fmt"
	"ofordcode.ofo.so/baseservice/osf-go.git"
)

type Hi struct { }

func (h Hi) SayHello(ctx context.Context, req *pb.Request) (*pb.Response, error){
	msg := fmt.Sprintf("Hello, %s", req.Name)

	return &pb.Response{
		Message: msg,
	}, nil
	return nil, nil
}

func main(){
	helloService := osf.ServiceConf{
		Name : "ofo.hello.HY.hiService",
		Register:pb.RegisterHiServer,
		Server : Hi{},
		NoRegistration:true,
	}
	osf.NewServer().Register(helloService).Start(8888)
}