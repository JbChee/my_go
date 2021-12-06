package server

import (
	//"bolg/pkg/errcode"
	//"bolg/pkg/errcode"
	"context"
	"encoding/json"
	"tag-server/pkg/errcode"

	"tag-server/pkg/bapi"
	//"tag-server/pkg/errcode"

	//"encoding/json"

	//"github.com/go-programming-tour-book/tag-service/pkg/bapi"

	//"google.golang.org/grpc"

	//"github.com/go-programming-tour-book/tag-service/pkg/errcode"
	pb "tag-server/proto"
)

type TagServer struct {
	pb.UnimplementedTagServiceServer
}



func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		//return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
		return nil, err
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
		//return nil, err
	}
	return &tagList, nil
}

//func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
//	opts = append(opts, grpc.WithInsecure())
//	return grpc.DialContext(ctx, target, opts...)
//}