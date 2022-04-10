package rpc_api

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"simpleApi/db"
	"simpleApi/etc"
	simple "simpleApi/protos"
	"testing"
)

var (
	conn *grpc.ClientConn
	err error
)

func init() {
	etc.InitConfig("../etc/config.yaml")
	db.InitMsql()
	network := "192.168.145.120:8082"
	conn, err = grpc.Dial(network, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	fmt.Println("connect rpc server success")
}

func TestApi_GetUserInfo(t *testing.T) {
	client := simple.NewSimpleApiClient(conn)
	res, err := client.GetUserInfo(context.Background(), &simple.UserRequest{UserId:1})
	if err != nil {
		panic(err)
	}

	fmt.Printf("user: %+v", res)
}
