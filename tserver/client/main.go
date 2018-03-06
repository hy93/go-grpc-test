package client

import (
	pb "go-grpc-test/proto/tserver" // 引入proto包

	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:10006"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHiClient(conn)

	// 调用方法
	reqBody := new(pb.Request)
	reqBody.Name = "gRPC"
	r, err := c.Hello(context.Background(), reqBody)

	if err != nil {
		grpclog.Fatalln(err)
	}

	fmt.Println(r.Message)
}

func GetClient(address string) string {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHiClient(conn)

	// 调用方法
	reqBody := new(pb.Request)
	reqBody.Name = "gRPC"
	r, err := c.Hello(context.Background(), reqBody)

	if err != nil {
		grpclog.Fatalln(err)
	}

	//fmt.Println(r.Message)
	return r.Message
}
