package main

import (
	"context"
	"log"
	"time"

	"github.com/kyungmun/golearning/grpc-sample/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn

	// Set up a connection to the server.
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &chat.Message{Body: "Hello from Client"})
	if err != nil {
		log.Fatalf("could not server response Sayhello: %v\n", err)
	}
	log.Printf("Response from Server: %s\n", r.Body)
}
