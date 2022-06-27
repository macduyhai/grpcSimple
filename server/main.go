package main

import (
	"fmt"
	"log"
	"net"

	"github.com/macduyhai/grpcSimple/simplepb"
	_ "github.com/macduyhai/grpcSimple/tree/main/simplepb"

	"google.golang.org/grpc"
)

type server struct {
	simplepb.UnimplementedGreeterServer
}

func main() {
	fmt.Println("Service proto simon")
	lis, err := net.Listen("tcp", "0.0.0.0:8899")
	if err != nil {
		log.Fatalf("Error while create listen %v", err)
	}
	s := grpc.NewServer()
	simplepb.RegisterGreeterServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Error while server %v", err)
	}

}
