package main

import (
	"context"
	"log"
	"time"

	pb "github.com/kyungmun/golearning/grpc-sample2/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":9091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect : %v\n", err)
	}
	defer conn.Close()

	client := pb.NewUserManagementServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var newUser = pb.UserRequest{Name: "kyungmun", Age: 49, Mobile: "010-1234-1234"}

	newResponse, err := client.CreateUser(ctx, &newUser)
	if err != nil {
		log.Fatalf("CreateUser : %v", err)
	}
	log.Printf("Success CreateUser : %v\n", newResponse)

	var delUser = pb.UserID{Uuid: newResponse.Uuid}
	delResponse, err := client.DeleteUser(ctx, &delUser)
	if err != nil {
		log.Fatalf("DeleteUser : %v", err)
	}
	log.Printf("Success DeleteUser : %v\n", delResponse)

}
