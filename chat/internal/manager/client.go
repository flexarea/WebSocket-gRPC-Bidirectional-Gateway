package manager

import (
    "sync"
    "fmt"
    "github.com/google/uuid"
    "github.com/gorilla/websocket"
)

type Client struct {
    Id   string
    Conn *websocket.Conn
    Send chan []byte
}

func (c *Client) ReadPump(hub *Hub, g*Gateway){

    // Tear down
    defer func(){
        hub.unregister <- c
        c.Conn.Close()
    }()

    for{
        _, message, err := c.Conn.ReadMessage()

        if err != nil{
            fmt.Println("Client disconnected")
            break
        }

        ctx, cancel := context.WithTimeout(context.Background(), time.second)


        resp, err := g.grpcClient.HandleMessage(ctx, &pb.MessageRequest{
            Content: string(message),	
            SrcUserId: 214,	
            DestUserId: 40,	
        })

        if err != nil { 
            fmt.Println("Error from gRPC server: ", err)
            cancel()
            continue
        }

        cancel()

        c.Send <- []byte(resp.Content)

    }

}

/*
    TODO: Is it possible to have a single writePump for all conections?
*/
func (c *Client) WritePump(){
    defer c.Conn.Close()

    for message := range c.send{
        err := c.Conn.WriteMessage(websocket.TextMessage, message)
        if err != nil { 
            break
        }
    }

    c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
}
