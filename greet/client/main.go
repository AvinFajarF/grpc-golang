package client

import (
	"google.golang.org/grpc"
	"log"
)

func main() {
	var address string = "localhost:50051"

	conn, err := grpc.Dial(address)

	if err != nil {
		log.Fatalf("Failed to connection %v \n", err.Error())
	}

	defer conn.Close()

}
