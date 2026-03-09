package registry

type Client struct {
    id   string
    conn *websocket.Conn
    send chan []byte
}

type Hub struct {
    clients map[string]*Client
    register chan *Client
    unregister chan *Client
    broadcast chan []byte
}

func NewHub() *Hub {
    return &Hub{
        clients:    make(map[string]*Client),
        register:   make(chan *Client),
        unregister: make(chan *Client),
        broadcast:  make(chan []byte),
    }
}

func (h *Hub) run() {
    for {
        select {
        case client := <-h.register:
            h.clients[client.id] = client
        case client := <-h.unregister:
            delete(h.clients, client.id)
            close(client.send)
        case message := <-h.broadcast:
            for _, c := range h.clients {
                select {
                case c.send <- message:
                default:
                    // drop if client send buffer full
                }
            }
        }
    }
}
