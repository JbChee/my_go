package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "grpc_demo/proto"

)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, args *pb.HelloRequest) (*pb.HelloReply, error) {
	r := &pb.HelloReply{
		Message: fmt.Sprintf("hello %s!", args.Name),
	}
	return r, nil
}


func main() {
	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
