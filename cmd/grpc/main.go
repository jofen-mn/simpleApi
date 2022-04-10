package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"simpleApi/db"
	"simpleApi/etc"
	simple "simpleApi/protos"
	"simpleApi/rpc_api"
)

var config = flag.String("c", "etc/config.yaml", "")

func main()  {
	flag.Parse()

	etc.InitConfig(*config)
	etc.LoggerInit()
	db.InitMsql()

	server := grpc.NewServer()
	network := etc.Config.Server.Host + ":" + etc.Config.Server.GrpcPort
	listener, err := net.Listen("tcp", network)
	if err != nil {
		log.Panicf("listen network error: %s", err.Error())
	}
	api := rpc_api.NewApi()
	simple.RegisterSimpleApiServer(server, api)

	if err = server.Serve(listener); err != nil {
		log.Panicf("start grpc server error: %s", err.Error())
	}

	go monit()

	defer server.Stop()
	defer listener.Close()
}

func monit() {
	network := etc.Config.Server.Host + ":1" + etc.Config.Server.HttpPort
	if err := http.ListenAndServe(network, nil); err != nil {
		panic(err)
	}
}
