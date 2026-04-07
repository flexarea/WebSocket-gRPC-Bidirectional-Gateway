package manager

import (
    "fmt"
    "github.com/gorilla/websocket"
    "encoding/json"
)

type ClientMessage struct{
	Content string	`json:content`
	SrcUserId int32 `json:src_user_id`
	DestUserId int32 `json:dest_user_id`
	Timestamp int64 `json:timestamp`
}

type Client struct {
    Id   string
    Conn *websocket.Conn
    Send chan *ClientMessage
}



type MessageHandler interface{
    Process(message []byte) (*ClientMessage, error)
}

func (c *Client) ReadPump(hub *Hub, handler MessageHandler){

    // Tear down
    defer func(){
        hub.Unregister <- c
        c.Conn.Close()
    }()


    for{
        _, rawMsg, err := c.Conn.ReadMessage()

        if err != nil{
            fmt.Println("Client disconnected")
            break
        }

        // 3. Call the interface method!
        resp, err := handler.Process(rawMsg)
        if err != nil {
            fmt.Println("Processing error:", err)
            continue
        }

        // Drop the response onto the conveyor belt
        hub.Broadcast <- resp

    }

}

func (c *Client) WritePump(){
    defer c.Conn.Close()


    // Blocking loop listening to channel 
    for message := range c.Send{

        JSONdata,_ := json.Marshal(*message)
        err := c.Conn.WriteMessage(websocket.TextMessage, JSONdata)

        if err != nil { 
            break
        }
    }

    // WS Close signal
    c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
}
