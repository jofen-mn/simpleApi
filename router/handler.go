package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"simpleApi/db"
	"simpleApi/seriallzer"
	"strconv"
	"time"
)

func Health(ctx *gin.Context) {
	res := make(map[string]interface{})
	res["data"] = "pong"
	res["time"] = time.Now()


	ctx.JSON(200, res)
}

type response struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
	ErrMsg string `json:"err_msg"`
}

func (r *response) reset(code int, errMsg string, data interface{}) {
	r.Data = data
	r.Code = code
	r.ErrMsg = errMsg
}


func GetUserInfo(ctx *gin.Context) {
	res := response{}
	value := ctx.Param("id")
	var err error
	if value == "" {
		res.reset(1000, "param empty", nil)
		logrus.Error(err)
		ctx.JSON(200, res)
		return
	}

	var user *db.User
	// 访问缓存
	key := db.UserCachePrefix + value
	var userValue string
	userValue, err = db.RedisClient.Get(key).Result()
	if err != nil && err != redis.Nil {
		// 访问redis失败
		res.reset(1000, "redis get error", nil)
		logrus.Errorf("redis get error: "+err.Error())
		ctx.JSON(200, res)
		return
	} else if err == redis.Nil {
		id := 0
		if id, err = strconv.Atoi(value); err != nil {
			logrus.Error(err)
			res.reset(1000, "param invalid", nil)
			ctx.JSON(200, res)
			return
		}

		if user, err = db.GetUserByID(id); err != nil {
			logrus.Errorf("get user: %d error: %s", err.Error())
			if err == db.NotFound {
				res.reset(1001, fmt.Sprintf("user %d not found", id), nil)
				ctx.JSON(200, res)
				return
			}

			res.reset(1002, fmt.Sprintf("get user: %d error", id), nil)
			return
		}

		userBytes, _ := json.Marshal(user)
		if _, err = db.RedisClient.Set(key, string(userBytes), 0).Result(); err != nil {
			logrus.Error("redis set key %s error: %s", key, err.Error())
		}
	} else {
		logrus.Infof("From cache, user value: %s", userValue)
		user = &db.User{}
		if err = json.Unmarshal([]byte(userValue), user); err != nil {
			logrus.Errorf("user key: %s\nvalue: %s\nUnmarshal error: %s", key, userValue, err.Error())
			res.reset(1004, "user data invalid", nil)
			ctx.JSON(200, res)
			return
		}
	}

	res.reset(0, "", user)
	ctx.JSON(200, res)
}

func UpdateUserInfo(ctx *gin.Context) {
	res := &response{}
	param := ctx.Param("id")
	var err error
	if param == "" {
		logrus.Error(err)
		res.reset(1000, "param empty", nil)
		ctx.JSON(200, res)
		return
	}

	key := db.UserCachePrefix + param
	// 删除缓存user
	if _, err = db.RedisClient.Del(key).Result(); err != nil {
		logrus.Error("delete user: %s from cache error: %s", key, err.Error())
		res.reset(1000, "update user: %s error", param)
		ctx.JSON(200, res)
		return
	}

	id := 0
	if id, err = strconv.Atoi(param); err != nil {
		logrus.Error(err)
		res.reset(1000, "param invalid", nil)
		ctx.JSON(200, res)
		return
	}
	req := &seriallzer.UserUpdateReq{}
	if err = ctx.ShouldBindJSON(req); err != nil {
		logrus.Error(err)
		res.reset(1000, "param invalid", nil)
		ctx.JSON(200, res)
		return
	}

	// 更新数据库
	if err = db.UpdateUser(id, req); err != nil {
		logrus.Errorf("update user: %d to mysql error: %s", id, err.Error())
		res.reset(1000, fmt.Sprintf("update user: %d error", id), nil)
		ctx.JSON(200, res)
		return
	}

	time.Sleep(100 * time.Millisecond)

	if _, err = db.RedisClient.Del(key).Result(); err != nil {
		logrus.Errorf("delete user: %d from cache error: %s", id, err.Error())
	}

	ctx.JSON(200, res)
}
