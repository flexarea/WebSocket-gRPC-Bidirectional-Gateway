package manager

type Hub struct {
    clients map[string]*Client
    register chan *Client
    unregister chan *Client
    broadcast chan []byte
}

func (hub *Hub) run(){
    for{
        select{
        case client := <-hub.register:
            h.clients[client.id] = client //register a client | new connection
        case client := <-hub.unregister:
            delete(h.clients, client.id] //unregister a client | client disconnect
            close(client.send) //close goroutine
        case message := <-hub.broadcast: //broad cast to all clients
            for _, c := range h.clients {
                    select{
                        case c.send <- message: // send message
                        default:
                    }
                }
        }
    }
}
