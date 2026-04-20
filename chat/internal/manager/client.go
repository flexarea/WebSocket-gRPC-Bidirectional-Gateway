package manager

import (
    "fmt"
    "io"
    "log"
    "time"
    "github.com/gorilla/websocket"
    "encoding/json"
    pb "xchat.io/proto"
)

type ClientMessage struct{
    Content string	`json:"content"`
    SrcUserId int32 `json:"src_user_id"`
    DestUserId int32 `json:"dest_user_id"`
    Timestamp int64 `json:"timestamp"`
    GRPCTime int64 `json:"grpcTime"`
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
        var WSMsg ClientMessage

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
            GRPCTime: time.Now().UnixNano(),
        })


        if err != nil {
            log.Println("Send failed. Stopping ReadPump:", err)
            break
        }

        //c.Send <- resp
    }

}

func (c *Client) ListenGrpcStream(){

    defer close(c.Send)

    for{

        msg, err := c.Stream.Recv()

        if err == io.EOF {
            log.Println("The server closed the stream.")
            return 
        }

        if err != nil {
            log.Println("Error processing gRPC reply:", err)
            return
        }

        //log.Printf("Received from gRPC: %+v", msg)

        grpcLatencyMs := time.Now().UnixNano()-msg.GRPCTime

        c.Send <- &ClientMessage{
            Content: msg.Content,
            SrcUserId: msg.SrcUserId,
            DestUserId: msg.DestUserId,
            Timestamp: msg.Timestamp,
	    GRPCTime: grpcLatencyMs,
        }
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
