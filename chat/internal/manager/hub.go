package manager


type Hub struct {
    Clients map[string]*Client
    Register chan *Client
    Unregister chan *Client
    Broadcast chan *ClientMessage
}

func (hub *Hub) Run(){
    for{
        select{
        case c := <-hub.Register:
            hub.Clients[c.Id] = c //register a client | new connection
        case c := <-hub.Unregister:
        if _, ok := hub.Clients[c.Id]; ok {
                    delete(hub.Clients, c.Id) //unregister a client | client disconnect
                    close(c.Send) //close goroutine
            }
        case message := <-hub.Broadcast: //broad cast to all clients
            for _, c := range hub.Clients {
                select{
                case c.Send <- message: // send message
                default:
                    close(c.Send)
                    delete(hub.Clients, c.Id) 
                }
            }
        }
    }
}
