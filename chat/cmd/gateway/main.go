/*
package main

import (
"log"
"xchat.io/internal/gateway"
)

func main() {
err := gateway.StartServer(":8080")
if err != nil {
log.Fatal(err)
}
}
*/

package main

import (
	"log"
	"net/http"

	"xchat.io/internal/gateway"
	"xchat.io/internal/manager"
	pb "xchat.io/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// Dial the gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	// Initialize the client
	client := pb.NewMessageClient(conn)

	// Initialize hub
	hub := &manager.Hub{
		Clients:    make(map[string]*manager.Client),
		Register:   make(chan *manager.Client),
		Unregister: make(chan *manager.Client),
		Broadcast:  make(chan *manager.ClientMessage),
	}

	// Inject the client into your gateway struct
	gw := gateway.NewGateway(client, hub)

	go hub.Run()

	// create router
	router := gateway.NewRouter(gw)

	log.Println("Gateway listening on :8080")

	// Start the server, passing the custom router instead of nil
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("failed to serve http: %v", err)
	}
}
