package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/coreos/etcd/proxy/grpcproxy"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"strings"
	"github.com/coreos/etcd/clientv3"
	"tag-server/internal/middleware"
	pb "tag-server/proto"
	"tag-server/server"
	"time"
)

//type server struct {
//	pb.UnimplementedTagServiceServer
//	pb.TagServiceServer
//}

var grpcPort string
var httpPort string
var port string

func init() {
	//flag.StringVar(&grpcPort, "grpc_port", "8001", "grpc启动端口")
	//flag.StringVar(&grpcPort, "grpc_port", "8001", "grpc启动端口")
	//flag.StringVar(&port, "port", "8003", "启动端口")
	flag.StringVar(&port, "port", "8004", "启动端口号")
	flag.Parse()

}

func main() {
	/*1.3.0*/
	err := RunServer(port)
	if err != nil {
		log.Fatalf("Run Serve err: %v", err)
	}

	//---------------

	//l, err := RunTcpServer(port)
	//if err != nil {
	//	log.Fatalf("err = %v", err)
	//}
	//m := cmux.New(l)
	//grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	//
	//httpL := m.Match(cmux.HTTP1Fast())
	//
	//grpcS := RunGrpcServer(port)
	//httpS := RunHttpServer(port)
	//go grpcS.Serve(grpcL)
	//go httpS.Serve(httpL)
	//httpS.Serve(httpL)
	//-------------v1.2.0---------
	//errs := make(chan error)
	//go func() {
	//	err := RunHttpServer(httpPort)
	//	if err != nil {
	//		errs <- err
	//	}
	//}()
	//
	//go func() {
	//	err := RunGrpcServer(grpcPort)
	//	if err != nil {
	//		errs <- err
	//	}
	//}()
	//
	//select {
	//case err := <-errs:
	//	log.Fatalf("err = %v", err)
	//}

	//-------------v1.0.0----------
	//s := grpc.NewServer()
	//pb.RegisterTagServiceServer(s, server.NewTagServer())
	//reflection.Register(s)
	//
	//fmt.Println("xxxxxxxxxxxxxxxxxxxxxxxxxx")
	//lis, err := net.Listen("tcp", ":8001")
	//fmt.Println(lis)
	//if err != nil {
	//	log.Fatal("net listin err : %v", err)
	//}
	//err = s.Serve(lis)
	//if err != nil {
	//	log.Fatal("server.server err = %v", err)
	//}
}

/* v1.3.0*/
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func RunServer(port string) error {
	httpMux := runHttpServer()
	grpcS := runGrpcServer()
	gatewayMux := runGrpcGatewayServer()

	httpMux.Handle("/", gatewayMux)

	//增加服务发现
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:[]string{"http://localhost:2379"},DialTimeout:time.Second * 60,
	})
	if err != nil {
		return err
	}
	defer etcdClient.Close()

	target := fmt.Sprintf("/etcdv3://go-programming-tour/grpc/%s", "tar-service")
	grpcproxy.Register(etcdClient, target, ":" + port, 60)

	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
}

func runHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return serveMux
}

func runGrpcServer() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middleware.AccessLog,
			middleware.ErrorLog,
			middleware.Recovery,
			HelloInterceptor,
		)),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func runGrpcGatewayServer() *runtime.ServeMux {
	endpoint := "0.0.0.0:" + port
	runtime.HTTPError = grpcGatewayError
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts)

	return gwmux
}

type httpError struct {
	Code    int32  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func grpcGatewayError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	httpError := httpError{Code: int32(s.Code()), Message: s.Message()}
	details := s.Details()
	for _, detail := range details {
		if v, ok := detail.(*pb.Error); ok {
			httpError.Code = v.Code
			httpError.Message = v.Message
		}
	}

	resp, _ := json.Marshal(httpError)
	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))
	_, _ = w.Write(resp)
}

/*1.3.0*/

func RunHttpServer(port string) *http.Server {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))

	})
	//return http.ListenAndServe(":"+port, serveMux)
	return &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}
}

func RunGrpcServer(port string) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middleware.AccessLog,
			middleware.ErrorLog,
			middleware.Recovery,
		)),
	}

	s := grpc.NewServer(opts...)
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	//lis, err := net.Listen("tcp", ":"+port)
	//if err != nil {
	//	//log.Fatalf("err%v",err)
	//	return err
	//}
	//return s.Serve(lis)
	return s
}

func RunTcpServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

//拦截器
func HelloInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("你好")
	resp, err := handler(ctx, req)
	log.Println("再见")
	return resp, err
}
