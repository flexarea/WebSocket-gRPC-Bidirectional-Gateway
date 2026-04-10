package gateway

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"context"
	"time"
	"encoding/json"
	"xchat.io/internal/manager"
	"github.com/google/uuid" // Don't forget this import!
	pb "xchat.io/proto"
)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true},
}

type Gateway struct {
	grpcClient pb.Message_HandleMessageClient
	hub *manager.Hub
}

func NewGateway (client pb.MessageClient, hub *manager.Hub) *Gateway{
	return &Gateway{
		grpcClient: client,
		hub: hub,
	}
}

/*
* Note: This function will create a new goroutine for every incoming connection 
*/


func (g *Gateway) HandleWebSocket(w http.ResponseWriter, r *http.Request){


	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}


	stream, err := g.grpcClient.HandleMessage(context.Background())

	if err != nil {
		log.Println("Failed to create stream:", err)
		return err
	}

	client := &manager.Client{
		Id: uuid.New().String(),
		Conn: conn,
		Send: make(chan *manager.ClientMessage, 256),
		Stream: stream,
	}

	// log new client
	fmt.Println("New client ID: ", client.Id)

	// register to hub
	g.hub.Register <- client

	// spin goroutines
	go client.WritePump()
	go client.ListenStream()
	client.ReadPump(g.hub)

}

