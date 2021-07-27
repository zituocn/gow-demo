package rpc_handler

import (
	user "github.com/zituocn/gow-demo/grpc/proto"
	"github.com/zituocn/gow/lib/logy"
	"github.com/zituocn/gow/lib/rpc"
	"google.golang.org/grpc"
)

func InitRPCServer() {
	g, err := rpc.NewServer(10001)
	if err != nil {
		panic(err)
	}
	handler(g.Server)
	g.Run()
	logy.Infof("[GRPC] Listening and serving TCP on %v", g.Port)

}

func handler(g *grpc.Server) {
	user.RegisterUserServer(g, &UserHandler{})
}
