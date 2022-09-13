package service

import (
	"admin/dao"
	"errors"
	"fmt"

	"github.com/gomodule/redigo/redis"
	jsoniter "github.com/json-iterator/go"
)

const SystemSecretKey = "bdf5fcc8eda9458c9e7c203903321887"

var json = jsoniter.ConfigFastest

var (
	ErrFreq        = errors.New("访问过于频繁,请稍后再试")
	ErrAuthFail    = errors.New("账号密码或验证码错误")
	ErrLoginExpire = errors.New("登录信息已失效")
)

func CheckFreq(key string, count, sec int64) (num int64, e error) {
	conn := dao.Cache().Get()
	defer conn.Close()

	num, e = redis.Int64(conn.Do("GET", key))
	if e != nil && e != redis.ErrNil {
		return
	}

	if num >= count {
		return num, ErrFreq
	}

	if num == 0 {
		num++
		if _, e = redis.String(conn.Do("SET", key, num, "EX", sec, "NX")); e == nil {
			return
		}
	}

	num, e = redis.Int64(conn.Do("INCR", key))
	return
}

func QpsLimit(key string, count, sec int64) (num int64, e error) {
	if num, e = CheckFreq("QPS_"+key, count, sec); e != nil && !errors.Is(e, ErrFreq) {
		fmt.Println(e)
	}
	return
}

func QpsClear(key string) (e error) {
	conn := dao.Cache().Get()
	defer conn.Close()

	_, e = conn.Do("DEL", "QPS_"+key)
	return
}
