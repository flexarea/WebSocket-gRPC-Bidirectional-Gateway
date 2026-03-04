package main

import (
	"log"
	"xchat.io/internal/gateway"
)

func main() {
	err := gateway.StartServer(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
