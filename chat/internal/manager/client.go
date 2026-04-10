package manager

import (
    "fmt"
    "github.com/gorilla/websocket"
    "encoding/json"
)

type ClientMessage struct{
    Content string	`json:"content"`
    SrcUserId int32 `json:"src_user_id"`
    DestUserId int32 `json:"dest_user_id"`
    Timestamp int64 `json:"timestamp"`
}

type Client struct {
    Id   string
    Conn *websocket.Conn
    Send chan *ClientMessage
    Stream pb.Message_HandleMessageClient
}



type MessageHandler interface{
    Process(message []byte) (*ClientMessage, error)
}

func (c *Client) ReadPump(hub *Hub){

    // Tear down
    defer func(){
        hub.Unregister <- c
        c.Conn.Close()
        c.Stream.CloseSend()
    }()


    for{
        _, rawMsg, err := c.Conn.ReadMessage()

        if err != nil{
            fmt.Println("Client disconnected")
            break
        }

        // WS JSON Payload Struct
        var WSMsg manager.ClientMessage

        // parse WS JSON Payload
        if err := json.Unmarshal(rawMsg, &WSMsg); err != nil {
            log.Println("error parsing json payload", err)
        }

        // gRPC client stream
        err = c.Stream.Send(&pb.ChatMessage{
            Content:    WSMsg.Content,
            SrcUserId:  WSMsg.SrcUserId,
            DestUserId: WSMsg.DestUserId,
            Timestamp: WSMsg.Timestamp,
        })


        if err != nil {
            log.Println("Error processing gRPC:", err)
            continue
        }

        //c.Send <- resp
    }

}

func (c *Client) ListenGrpcStream(){

    for{

        msg, err := c.Stream.Recv()

        if err != nil {
            log.Println("Error processing gRPC reply:", err)
        }

        c.Send <- msg

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
