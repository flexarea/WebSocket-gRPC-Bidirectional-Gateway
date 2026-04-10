package message

import (
	"context"
	"fmt"
	pb "xchat.io/proto"
)

type MessageServer struct {
	pb.UnimplementedMessageServer  // required boilerplate
}

func NewMessageServer() *MessageServer {
	return &MessageServer{}
}

func (s *MessageServer) HandleMessage(stream grpc.BidiStreamingServer[pb.ChatMessage, pb.ChatMessage]) error {

	outgoing := make(chan *pb.ChatMessage)

	// RECEIVE loop
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		// send back message
		outgoing <- &pb.ChatMessage{
			Content: "echo: " + msg.Content,
		}
	}

	// SEND loop
	go func(){
		for msg := range outgoing {
			if err := stream.Send(msg); err != nil {
				return
			}
		}

	}{}

rqMessage := req.Content
srcUserId := req.SrcUserId

return &pb.MessageReply{
	Content: fmt.Sprintf("From gRPC Server: received message %s from %d", rqMessage, srcUserId),
	SrcUserId: req.SrcUserId,
	DestUserId: req.DestUserId,
	Timestamp: req.Timestamp,
}, nil
}
