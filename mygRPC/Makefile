.PHONY: proto server client clean

proto: proto/hello.pb.go proto/hello_grpc.pb.go

proto/hello.pb.go proto/hello_grpc.pb.go: proto/hello.proto
	protoc --go_out=. --go_opt=paths=source_relative \
	       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	       $<

server: proto
	go run server/main.go

client: proto
	go run client/main.go

clean:
	rm -f proto/*.pb.go
