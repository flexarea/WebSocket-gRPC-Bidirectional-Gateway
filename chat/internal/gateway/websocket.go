package gateway

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"context"
	"time"
	pb "xchat.io/proto"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true},
}

type Gateway struct {
	grpcClient pb.MessageClient
}

func NewGateway (client pb.MessageClient) *Gateway{
	return &Gateway{
		grpcClient: client,
	}
}

func (g *Gateway) HandleWebSocket(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")
	conn.WriteMessage(websocket.TextMessage, []byte("Connected to echo server!"))

	// read message from Websocket
	for {

		messageType, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Client disconnected")
			break
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		defer cancel()

		resp, err := g.grpcClient.HandleMessage(ctx, &pb.MessageRequest{
			Content: string(data),	
			SrcUserId: 214,	
			DestUserId: 040,	
		})

		if err != nil { 
			log.Fatal(err) 
		}
	
		fmt.Println(resp.Content) //gRPC server output

		conn.WriteMessage(messageType, []byte("Echo: "+string(data)))
	}
}

