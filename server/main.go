package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	fmt.Println("Service proto simon")
	lis, err := net.Listen("tcp", "0.0.0.0:8899")
	if err != nil {
		log.Fatalf("Error while create listen %v", err)
	}
	s := grpc.NewServer()

}
