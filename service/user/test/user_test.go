package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
	v1 "user/api/user/v1"
)

var userClient v1.UserClient
var conn *grpc.ClientConn

func init() {
	Init()

}

// Init 初始化 grpc 链接 注意这里链接的 端口
func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	if err != nil {
		panic("grpc link err" + err.Error())
	}
	userClient = v1.NewUserClient(conn)
}

func TestCreateUser(t *testing.T) {

	defer conn.Close()
	rsp, err := userClient.CreateUser(context.Background(), &v1.CreateUserRequest{
		Mobile:   fmt.Sprintf("138888874%d", 1),
		Password: "admin123",
		NickName: fmt.Sprintf("YWWW%d", 1),
	})
	if err != nil {
		panic("grpc 创建用户失败" + err.Error())
	}
	fmt.Println(rsp.Id)
}
