package dto

import (
	"admin/models"
	"errors"
)

type UserQueryReq struct {
	QueryBase
	Acct     string `form:"acct"`
	UserType models.UserType
}

type UserCreateReq struct {
	Acct     string          `json:"acct"`
	Alias    string          `json:"alias"`
	PassWord string          `json:"pwd"`
	UserType models.UserType `json:"user_type" `
}

func (_u UserCreateReq) Valid() error {
	if _u.Acct == "" {
		return errors.New("账号不能为空")
	}
	if _u.Alias == "" {
		return errors.New("名称不能为空")
	}
	if _u.PassWord == "" {
		return errors.New("密码不能为空")
	}
	return nil
}
