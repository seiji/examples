package main_bm

import (
	"testing"
	"time"

	"golang.org/x/net/context"
	g "google.golang.org/grpc"

	sv "github.com/seiji/grpc-examples/helloworld/greeter_server"
	pb "github.com/seiji/grpc-examples/helloworld/helloworld"
)

const (
	port        = ":50051"
	defaultName = "world"
)

func init() {
	go sv.Start(port)
	time.Sleep(time.Second)
}

func BenchmarkGRPC(b *testing.B) {
	conn, err := g.Dial("127.0.0.1"+port, g.WithInsecure())
	if err != nil {
		b.Fatalf("grpc connection failed: %v", err)
	}
	c := pb.NewGreeterClient(conn)
	for n := 0; n < b.N; n++ {
		doGRPC(c, b)
	}
}

func doGRPC(c pb.GreeterClient, b *testing.B) {
	resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: defaultName})
	if err != nil {
		b.Fatalf("grpc request failed: %v", err)
	}
	if resp == nil || resp.Message == "" {
		b.Fatalf("grpc response is wrong: %v", resp)
	}
}
