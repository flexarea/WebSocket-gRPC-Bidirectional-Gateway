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

func (s *MessageServer) HandleMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageReply, error) {

	rqMessage := req.Content
	srcUserId := req.SrcUserId

	return &pb.MessageReply{
		Content: fmt.Sprintf("From gRPC Server: received message %s from %d", rqMessage, srcUserId),
		SrcUserId: req.SrcUserId,
		DestUserId: req.DestUserId,
		Timestamp: req.Timestamp,
	}, nil
}
