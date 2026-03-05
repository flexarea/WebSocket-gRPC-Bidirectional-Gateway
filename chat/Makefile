.PHONY: proto server client clean

proto: proto/message.pb.go proto/message_grpc.pb.go

proto/message.pb.go proto/message_grpc.pb.go: proto/message.proto
	protoc --go_out=. --go_opt=paths=source_relative \
	       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	       $<

server: proto
	go run cmd/server/main.go

gateway: proto
	go run cmd/gateway/main.go

clean:
	rm -f proto/*.pb.go
