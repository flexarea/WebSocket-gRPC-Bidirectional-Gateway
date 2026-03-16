package manager

type Hub struct {
    Clients map[string]*Client
    Register chan *Client
    Unregister chan *Client
    Broadcast chan []byte
}

func (hub *Hub) Run(){
    for{
        select{
        case client := <-hub.Register:
            h.Clients[client.id] = client //register a client | new connection
        case client := <-hub.Unregister:
            delete(h.Clients, client.id] //unregister a client | client disconnect
            close(client.send) //close goroutine
        case message := <-hub.Broadcast: //broad cast to all clients
            for _, c := range h.Clients {
                    select{
                        case c.Send <- message: // send message
                        default:
                    }
                }
        }
    }
}
