package main

import (
	"fmt"
	"golang.org/x/net/context"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	pb "go-grpc-test/proto/tserver"
	"net"
	"net/http"
)

type hiGrpcService struct{}

var hiService = hiGrpcService{}

func (h hiGrpcService) Hello(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	resp := new(pb.Response)
	resp.Message = "Hello Grpc :" + in.Name + "."
	return resp, nil
}

func main() {
	endpoint := "127.0.0.1:10006"
	conn, err := net.Listen("tcp", endpoint)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterHiServer(grpcServer, hiService)

	srv := &http.Server{
		Addr:    endpoint,
		Handler: myHandler(grpcServer),
	}

	s2 := &http2.Server{}
	http2.ConfigureServer(srv, s2)

	for {
		rwc, err := conn.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go s2.ServeConn(rwc, &http2.ServeConnOpts{BaseConfig: srv})
	}
}

func myHandler(server *grpc.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.ServeHTTP(w, r)
	})
}
