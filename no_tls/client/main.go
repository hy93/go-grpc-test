package main

import (
	pb "go-grpc-test/proto/hello_http" // 引入proto包

	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloHTTPClient(conn)

	// 调用方法
	reqBody := new(pb.HelloHTTPRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)

	if err != nil {
		grpclog.Fatalln(err)
	}

	fmt.Println(r.Message)
}

// OR: curl -X POST -k https://localhost:50052/example/echo -d '{"name": "gRPC-HTTP is working!"}'
