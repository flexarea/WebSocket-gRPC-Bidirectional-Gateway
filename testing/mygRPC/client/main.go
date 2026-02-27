package main

import (
	"context"
	"log"
	"time"
	pb "mygRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil { log.Fatal(err) }
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "World"})
	if err != nil { log.Fatal(err) }
	log.Println(resp.Message)  // "Hello World"
}
