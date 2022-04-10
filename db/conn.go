package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"simpleApi/etc"

	//mysql drivers
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	engine      *gorm.DB
	RedisClient *redis.Client
)

func InitMsql() {
	var err error
	url := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=%s&parseTime=True",
		etc.Config.Mysql.User, etc.Config.Mysql.Password, etc.Config.Mysql.IP, etc.Config.Mysql.Port,
		etc.Config.Mysql.Database, etc.Config.Mysql.Charset)
	engine, err = gorm.Open("mysql", url)
	if err != nil {
		logrus.Panicf("connect mysql error: %s", err.Error())
		return
	}

	logrus.Println("init mysql success")
}

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         etc.Config.Redis.Address,
		Password:     etc.Config.Redis.Password,
		DB:           etc.Config.Redis.Db,
		PoolSize:     etc.Config.Redis.PoolSize,
		MinIdleConns: etc.Config.Redis.MinIdleConns,
	})

	pong, err := RedisClient.Ping().Result()
	if err != nil {
		logrus.Panicf("connect redis error: %s", err.Error())
		return
	}

	logrus.Printf("connect redis success ping -> %s", pong)
}
