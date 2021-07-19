package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go-learn-repo/common"
	pb "go-learn-repo/proto/hello"
)

const (
	Address = "127.0.0.1:50052"
)

type helloService struct{}

var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s", in.Name)
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		common.CheckErr(err)
	}

	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("../keys/server.pem", "../keys/server.key")
	if err != nil {
		common.CheckErr(err)
	}

	// new Grpc Server,and start tls
	server := grpc.NewServer(grpc.Creds(creds))

	// register hello Server
	pb.RegisterHelloServer(server, HelloService)

	fmt.Printf("Listen on " + Address + " with TLS")

	server.Serve(listen)
}
