package main

import (
	"context"
	"log"
	"net"
	pb "mygRPC/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer  // required boilerplate
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Println("Server starting on :50051")
	s.Serve(lis)
}
