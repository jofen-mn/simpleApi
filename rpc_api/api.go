package rpc_api

import (
	"context"
	"github.com/sirupsen/logrus"
	"simpleApi/db"
	simple "simpleApi/protos"
)

type Api struct {
}

func NewApi() *Api {
	return &Api{
	}
}

func (api *Api) GetUserInfo(ctx context.Context, in *simple.UserRequest) (*simple.UserResponse, error) {
	logrus.Printf("UserRequest: %+v", in)
	res := &simple.UserResponse{}
	user, err := db.GetUserByID(int(in.UserId))
	if err != nil {
		return res, err
	}
	resetUserResponse(user, res)

	return res, nil
}