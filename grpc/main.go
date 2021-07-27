package main

import (
	"github.com/zituocn/gow"
	"github.com/zituocn/gow-demo/grpc/rpc_handler"
)

func init() {
	rpc_handler.InitRPCServer()
}

func main() {
	r := gow.Default()
	r.Run()
}
