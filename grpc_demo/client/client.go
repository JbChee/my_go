package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"log"

	pb "grpc_demo/proto"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

func main() {
	conn, err := grpc.Dial(":"+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	err = SayHello(c)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

}

func SayHello(client pb.GreeterClient) error {
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "xxxx"})
	if err != nil {
		return err
	}

	log.Println("client.sayhello resp : %s", resp.Message)
	return nil

}
