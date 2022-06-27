package main

import (
	"context"
	pb "grpcSimple/simplepb"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	name = "Anna"
)

func main() {
	log.Println("Client proto starting...")
	// flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial("0.0.0.0:8899", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	} else {
		log.Println("Connection done")
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// // Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
