package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/naming"
	"google.golang.org/grpc"
	"log"
	pb "tag-server/proto"
	"time"
)

func main() {
	ctx := context.Background()
	clientConn, err := GetClientConn(ctx, "tag-service", []grpc.DialOption{grpc.WithBlock()})
	if err != nil {
		log.Fatalf("err:%v", err)
	}
	defer clientConn.Close()
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, err := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{Name: "Go"})
	if err != nil {
		log.Fatalf("err :%v", err)
	}
	log.Printf("resp = %v", resp)

}

func GetClientConn(ctx context.Context, serviceName string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	//增加服务发现
	config := clientv3.Config{
		Endpoints: []string{"http://localhost:2379"}, DialTimeout: time.Second * 60,
	}

	cli, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}

	r := &naming.GRPCResolver{Client: cli}
	target := fmt.Sprintf("/etcdv3://go-programming-tour/grpc/%s", "tag-service")
	opts = append(opts, grpc.WithInsecure(), grpc.WithBlock(), grpc.RoundRobin(r))
	return grpc.DialContext(ctx, target, opts...)
}
