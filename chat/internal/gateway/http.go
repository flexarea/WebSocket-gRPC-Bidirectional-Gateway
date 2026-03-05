/*
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
*/
package gateway

import (
	"net/http"
)

func NewRouter(gw *Gateway) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", gw.HandleWebSocket)

	return mux
}
