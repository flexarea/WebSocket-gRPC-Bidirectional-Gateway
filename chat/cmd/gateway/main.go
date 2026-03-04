package main


import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true},
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")
	conn.WriteMessage(websocket.TextMessage, []byte("Connected to echo server!"))

	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Client disconnected")
			break
		}

		fmt.Println("Received:", string(data))
		conn.WriteMessage(messageType, []byte("Echo: "+string(data)))
	}
}


func main(){
	http.HandleFunc("/ws", handler)
	fmt.Println("Websocket server running on ws: 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
