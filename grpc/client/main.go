package main

import (
	"context"
	"github.com/zituocn/gow"
	user "github.com/zituocn/gow-demo/grpc/proto"
	"github.com/zituocn/gow/lib/rpc"
)

var (
	addr = "grpc-test-service:10001"
)

func main() {
	r := gow.Default()
	r.GET("/", grpc2http)
	r.Run(9090)
}

func grpc2http(c *gow.Context) {
	uid, _ := c.GetInt64("uid", 0)
	conn, err := rpc.NewClient(addr)
	if err != nil {
		c.DataJSON(1, err.Error())
		return
	}
	defer conn.Close()

	cc := user.NewUserClient(conn)

	req := &user.UserRequest{
		Uid: uid,
	}
	resp, err := cc.GetUser(context.Background(), req)
	if err != nil {
		c.DataJSON(1, err.Error())
		return
	}
	c.DataJSON(resp)
}
