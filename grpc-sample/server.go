package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kyungmun/golearning/grpc-sample/chat"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("go gRPC first tutorial")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &chat.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v\n", err)
	}

}
