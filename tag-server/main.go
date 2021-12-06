package main

import (
	"flag"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"tag-server/internal/middleware"
	pb "tag-server/proto"
	"tag-server/server"
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
	flag.StringVar(&port, "port", "8003", "启动端口")
	flag.Parse()

}

func main() {
	l, err := RunTcpServer(port)
	if err != nil {
		log.Fatalf("err = %v", err)
	}
	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))

	httpL := m.Match(cmux.HTTP1Fast())

	grpcS := RunGrpcServer(port)
	httpS := RunHttpServer(port)
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)
	httpS.Serve(httpL)
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
