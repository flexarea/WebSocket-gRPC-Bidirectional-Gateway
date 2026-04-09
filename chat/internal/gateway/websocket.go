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
	grpcClient pb.MessageClient
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

func (g *Gateway) Process(message []byte) (*manager.ClientMessage, error){
	// Setup the context for gRPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Make the gRPC Unary Call

	var WSMsg manager.ClientMessage

        if err := json.Unmarshal(message, &WSMsg); err != nil {
            log.Println("error parsing json payload", err)
        }

	resp, err := g.grpcClient.HandleMessage(ctx, &pb.MessageRequest{
		Content:    WSMsg.Content,
		SrcUserId:  WSMsg.SrcUserId,
		DestUserId: WSMsg.DestUserId,
		Timestamp: WSMsg.Timestamp,
	})

	if err != nil {
		return nil, err // Return the error back to the ReadPump
	}

	//return []byte(resp.Content), nil
	return &manager.ClientMessage{
		Content:    resp.Content,
		SrcUserId:  resp.SrcUserId,
		DestUserId: resp.DestUserId,
		Timestamp: resp.Timestamp,
	}, nil
}

func (g *Gateway) HandleWebSocket(w http.ResponseWriter, r *http.Request){


	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}


	client := &manager.Client{
		Id: uuid.New().String(),
		Conn: conn,
		Send: make(chan *manager.ClientMessage, 256),
	}

	// log new client
	fmt.Println("New client ID: ", client.Id)

	// register to hub
	
	g.hub.Register <- client

	go client.WritePump()
	client.ReadPump(g.hub, g)

}

