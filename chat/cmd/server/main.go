package main

import (
	"log"
	"net"
	"xchat.io/internal/message"
	pb "xchat.io/proto"
	"google.golang.org/grpc"
)


func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	// create grpc server
	grpcServer := grpc.NewServer()

	// instantiate message service
	msgServer := message.NewMessageServer()

	// Register service with gRPC server
	pb.RegisterMessageServer(grpcServer, msgServer)

	log.Printf("gRPC server listening at %v", lis.Addr())


	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to server: %v", err)
	}
}
