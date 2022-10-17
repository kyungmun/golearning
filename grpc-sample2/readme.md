#### proto 파일에서 go_package 옵션 설정
주의) 하나의 . 또는 / 값이 포함된 경로여야 한다.

> option go_package = "github.com/kyungmun/golearning/grpc-sample2/user";

#### 프로젝트 디렉토리에서 아래 명령 수행해서 grpc 파일을 자동 생성한다.
> protoc --go_out=. \
       --go_opt=paths=source_relative \
       --go-grpc_out=. \
       --go-grpc_opt=paths=source_relative \
       user/usermngt.proto

proto 파일 내용이 갱신 된다면 위 명령 재실행으로 생성된 소스를 재생성 하면 된다.

#### go mode init 으로 프로젝트 모듈 설정과 필요한 모듈 정리
> go mod init github.com/kyungmun/golearning/grpc-sample2

> go mod tidy

#### 메소드 구현
proto 파일에서 정의하고 grpc.pb.go 에 생성된 서비스 인터페이스 메소드를 구현한다.