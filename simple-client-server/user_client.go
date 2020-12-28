package simple_client_server

import (
	"sync"

	"google.golang.org/grpc"

	userpb "github.com/dojinkimm/go-grpc-example/protos/user"
)

var (
	once sync.Once
	cli  userpb.UserClient
)

func GetUserClient(serviceHost string) userpb.UserClient {
	once.Do(func() {
		conn, _ := grpc.Dial(serviceHost,
			grpc.WithInsecure(),
			grpc.WithBlock(),
		)

		cli = userpb.NewUserClient(conn)
	})

	return cli
}
