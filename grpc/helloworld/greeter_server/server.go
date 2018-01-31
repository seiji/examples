package server

import (
	"context"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/seiji/grpc-examples/helloworld/helloworld"
)

var (
	hostName, _ = os.Hostname()
)

type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: hostName + ": Hello " + in.Name}, nil
}

func Start(port string) {
	lis, _ := net.Listen("tcp", port)
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	log.Println(s.Serve(lis))
}
