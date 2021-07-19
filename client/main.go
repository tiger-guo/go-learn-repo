package main

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" // 引入grpc认证包

	"go-learn-repo/common"
	pb "go-learn-repo/proto/hello"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	// TLS连接
	creds, err := credentials.NewClientTLSFromFile("../keys/server.pem", "server")
	if err != nil {
		common.CheckErr(err)
	}

	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		common.CheckErr(err)
	}
	defer conn.Close()
	// 初始化客户端
	c := pb.NewHelloClient(conn)
	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		common.CheckErr(err)
	}
	fmt.Println(res.Message)
}
