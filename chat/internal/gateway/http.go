package gateway


import (
	"fmt"
	"net/http"
)

func StartServer(addr string) error{
	http.HandleFunc("/ws", HandleWebSocket)
	fmt.Println("WebSocket server running on", addr)
	return http.ListenAndServe(addr, nil)
}
