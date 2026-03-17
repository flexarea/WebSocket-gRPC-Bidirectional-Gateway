package manager

import (
    "fmt"
    "github.com/gorilla/websocket"
)

type Client struct {
    Id   string
    Conn *websocket.Conn
    Send chan []byte
}

type MessageHandler interface{
    Process(message []byte) ([]byte, error)
}

func (c *Client) ReadPump(hub *Hub, handler MessageHandler){

    // Tear down
    defer func(){
        hub.Unregister <- c
        c.Conn.Close()
    }()

    for{
        _, message, err := c.Conn.ReadMessage()

        if err != nil{
            fmt.Println("Client disconnected")
            break
        }


        // 3. Call the interface method!
        resp, err := handler.Process(message)
        if err != nil {
            fmt.Println("Processing error:", err)
            continue
        }

        // Drop the response onto the conveyor belt
        hub.Broadcast <- resp

    }

}

/*
TODO: Is it possible to have a single writePump for all conections?
*/
func (c *Client) WritePump(){
    defer c.Conn.Close()

    // Blocking loop listening to channel 
    for message := range c.Send{
        err := c.Conn.WriteMessage(websocket.TextMessage, message)
        if err != nil { 
            break
        }
    }

    // WS Close signal
    c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
}
