package manager

import (
    "sync"
    "fmt"
)

type Client struct {
    id   string
    conn *websocket.Conn
    send chan []byte
}

func (client *Client) ReadPump(hub *Hub){
    defer func(){
        hub.unregister <- c
        c.conn.close()
    }

    /*TODO: Implement the rest here*/

}
 
