package message

import (
	"io"
	"log"
	pb "xchat.io/proto"
)

type MessageServer struct {
	pb.UnimplementedMessageServer  // required boilerplate
}

func NewMessageServer() *MessageServer {
	return &MessageServer{}
}

func (s *MessageServer) HandleMessage(stream pb.Message_HandleMessageServer) error {

	outgoing := make(chan *pb.ChatMessage)

	// SEND loop
	go func(){

		for msg := range outgoing {
			if err := stream.Send(msg); err != nil {
				log.Println("Something broke", err)
				break
			}
		}

	}()

	defer close(outgoing)

	// RECEIVE loop
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			log.Println("The server closed the stream.")
			return nil
		}

		if err != nil {
			log.Println("Something broke", err)
			return err
		}


		// send back message
		outgoing <- &pb.ChatMessage{
			Content: "echo: " + msg.Content,
			SrcUserId: msg.SrcUserId,
			DestUserId: msg.DestUserId,
			Timestamp: msg.Timestamp,
			GRPCTime: msg.GRPCTime,
		}
	}


}
