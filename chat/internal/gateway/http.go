package gateway

import (
	"net/http"
)

func NewRouter(gw *Gateway) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", gw.HandleWebSocket)

	return mux
}
