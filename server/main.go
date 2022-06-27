package main

import (
	"context"
	"log"
	"net"

	"grpcSimple/simplepb"

	"google.golang.org/grpc"
)

type server struct {
	simplepb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *simplepb.HelloRequest) (*simplepb.HelloReply, error) {
	log.Printf("%v say hello \n", req.GetName())
	resp := &simplepb.HelloReply{
		Message: "I'm Simon",
	}
	return resp, nil

}

func main() {
	log.Println("Server proto is running ....")
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
