package main

import (
	"log"
	"net"

	pb "github.com/kyungmun/golearning/grpc-sample2/user"
	"google.golang.org/grpc"
)

func main() {

	//통신 할 네트워크 준비
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalln("grpc server listen fail")
	}

	//grpc server 생성
	grpcServer := grpc.NewServer()

	pb.RegisterUserManagementServiceServer(grpcServer, &pb.UserManagementServer{})

	log.Printf("grpc server listen at: %v\n", lis.Addr())

	//grpc server 실행
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpc server start fail : %v\n", err)
	}

}
