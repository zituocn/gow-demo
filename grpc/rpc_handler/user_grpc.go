package rpc_handler

import (
	"context"
	"fmt"
	user "github.com/zituocn/gow-demo/grpc/proto"
)

type UserHandler struct {
}

func (m *UserHandler) GetUser(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	resp = &user.UserResponse{
		Username: fmt.Sprintf("get %d response: zituocn", req.Uid),
	}
	return
}
