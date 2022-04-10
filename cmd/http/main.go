package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof" //DEBUG INTERFACE
	"simpleApi/db"
	"simpleApi/etc"
	"simpleApi/middleware"
	"simpleApi/router"
)

var config = flag.String("c", "etc/config.yaml", "config path")

func main()  {
	flag.Parse()

	etc.InitConfig(*config)
	etc.LoggerInit()
	db.InitMsql()
	db.InitRedis()

	go monit()

	server := gin.New()

	p := middleware.NewPrometheus("simpleApi")
	p.Use(server)

	r := server.Group("/api/v1")
	r.Use(func(context *gin.Context) {
		logrus.Println("Host", context.Request.Host)
		context.Next()
	})

	r.GET("/health", router.Health)
	r.GET("/user/:id", router.GetUserInfo)
	r.PUT("/user/:id", router.UpdateUserInfo)


	http.ListenAndServe("0.0.0.0:"+etc.Config.Server.HttpPort, server)
}

func monit() {
	network := etc.Config.Server.Host + ":1" + etc.Config.Server.HttpPort
	if err := http.ListenAndServe(network, nil); err != nil {
		panic(err)
	}
}
