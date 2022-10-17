package user

import (
	"context"
	"log"

	"github.com/google/uuid"
)

type UserManagementServer struct {
	UnimplementedUserManagementServiceServer
}

// usermngt_grpc.pb.go 에 자동으로 생성된 인터페이스 메소드 구현
func (u *UserManagementServer) CreateUser(ctx context.Context, in *UserRequest) (*UserResponse, error) {
	log.Printf("Create Request Received : %v", in)
	uuid := uuid.New()
	log.Printf("Generated User Id : %s\n", uuid.String())
	return &UserResponse{Name: in.Name, Age: in.Age, Mobile: in.Mobile, Uuid: uuid.String()}, nil
}

func (u *UserManagementServer) DeleteUser(ctx context.Context, in *UserID) (*UserResponse, error) {
	log.Printf("Delete Request Received : %v", in.Uuid)
	//해당 uuid 값으로 db의 값이 있는지 확인후 삭제 처리하고 응답.
	return &UserResponse{Name: "kyungmun", Age: 49, Mobile: "010-1234-1234", Uuid: in.Uuid}, nil
}
