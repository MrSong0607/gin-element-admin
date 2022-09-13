package dto

import (
	"errors"
)

const (
	Success      ApiResponseCode = 2000
	Unauthorized ApiResponseCode = 4001 //未登录
	Forbidden    ApiResponseCode = 4003 //无权限
	Unkown       ApiResponseCode = 5000
)

var (
	ErrAuth        = errors.New("权限错误")
	ErrParam       = errors.New("参数错误")
	ErrLoginExpire = errors.New("登录信息已过期")
)

type ApiResponseCode int

type H map[string]any

type ApiResponse struct {
	Code    ApiResponseCode `json:"code"`
	Message string          `json:"msg"`
	Data    interface{}     `json:"data"`
}

type ArrayResult[T any] struct {
	Count int64 `json:"count"`
	List  []T   `json:"list"`
}

type QueryBase struct {
	Page int `json:"page" form:"page"`
	Size int `json:"size" form:"size"`
}

type QueryTimeBase struct {
	Start int64 `json:"start" form:"start"`
	End   int64 `json:"end" form:"end"`
}

type EntityBase struct {
	ID uint64 `json:"id" form:"id"`
}
