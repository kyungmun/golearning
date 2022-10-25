package user

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/google/uuid"
)

func NewUserManagementServer() *UserManagementServer {

	return &UserManagementServer{
		user_list: &UserList{},
	}
}

type UserManagementServer struct {
	UnimplementedUserManagementServiceServer
	user_list *UserList
}

func (us *UserManagementServer) Run() error {
	//통신 할 네트워크 준비
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalln("grpc server listen fail")
	}

	//grpc server 생성
	grpcServer := grpc.NewServer()

	RegisterUserManagementServiceServer(grpcServer, us)

	log.Printf("grpc server listen at: %v\n", lis.Addr())

	//grpc server 실행
	return grpcServer.Serve(lis)
}

// usermngt_grpc.pb.go 에 자동으로 생성된 인터페이스 메소드 구현
func (us *UserManagementServer) CreateUser(ctx context.Context, in *UserRequest) (*UserResponse, error) {
	log.Printf("Create Request Received : %v", in)
	uuid := uuid.New()
	log.Printf("Generated User Id : %s\n", uuid.String())
	user := &UserResponse{Name: in.Name, Age: in.Age, Mobile: in.Mobile, Uuid: uuid.String()}

	var user_list *UserList = &UserList{}
	readBytes, err := ioutil.ReadFile("user.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("file not found. user.json file creating")
			user_list.Users = append(user_list.Users, user)
			jsonBytes, err := protojson.Marshal(user_list)
			if err != nil {
				log.Fatalf("json marshaling failed: %v", err)
			}
			if err := ioutil.WriteFile("user.json", jsonBytes, 0644); err != nil {
				log.Fatalf("user.json file create failed: %v", err)
			}
			return user, nil
		} else {
			log.Fatalf("json file read failed : %v", err)
		}
	}
	if err := protojson.Unmarshal(readBytes, user_list); err != nil {
		log.Fatalf("user.json parser failed: %v", err)
	}
	user_list.Users = append(user_list.Users, user)
	jsonBytes, err := protojson.Marshal(user_list)
	if err != nil {
		log.Fatalf("json marshaling failed: %v", err)
	}
	if err := ioutil.WriteFile("user.json", jsonBytes, 0644); err != nil {
		log.Fatalf("user.json file create failed: %v", err)
	}

	return user, nil
}

func (us *UserManagementServer) GetUsers(ctx context.Context, in *GetUserParams) (*UserList, error) {
	readBytes, err := ioutil.ReadFile("user.json")
	if err != nil {
		log.Fatalf("usr.json read failed: %v", err)
	}
	var user_list *UserList = &UserList{}
	err = protojson.Unmarshal(readBytes, user_list)
	if err != nil {
		log.Fatalf("json Marshal failed: %v", err)
	}

	return user_list, nil
}

func (us *UserManagementServer) DeleteUser(ctx context.Context, in *UserID) (*UserResponse, error) {
	log.Printf("Delete Request Received : %v", in.Uuid)
	//해당 uuid 값으로 db의 값이 있는지 확인후 삭제 처리하고 응답.
	return &UserResponse{Name: "kyungmun", Age: 49, Mobile: "010-1234-1234", Uuid: in.Uuid}, nil
}
