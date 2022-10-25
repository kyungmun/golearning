package main

import (
	"log"

	pb "github.com/kyungmun/golearning/grpc-sample2/user"
)

func main() {
	var user_mng_server *pb.UserManagementServer = pb.NewUserManagementServer()

	if err := user_mng_server.Run(); err != nil {
		log.Fatalf("grpc server start fail : %v\n", err)
	}

}
