package main

import (
	pb "github.com/AvinFajarF/grpc-tutorial/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.GreetServiceServe
}

func main() {

	var address = "0.0.0.0:50051"

	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}

	log.Printf("Connecting to %s \n", address)

	s := grpc.NewServer()

	if err := s.Serve(lis); err != nil {
		log.Printf("Failed to server: %v \n", err)
	}

}
