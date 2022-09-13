package dto

import (
	"admin/models"
	"admin/utils"
)

type LoginForm struct {
	Acct   string `json:"acct"`
	Passwd string `json:"passwd"`
	Code   string `json:"code"`
}

type SessionInfo struct {
	ID        uint64          `json:"id"`
	Acct      string          `json:"acct"`
	Alias     string          `json:"alias"`
	UserType  models.UserType `json:"user_type"`
	OtpSecret string          `json:"otp_secret"`
	Token     string          `json:"token"`
}

type UserProfileReq struct {
	ID    uint64
	Alias string `json:"alias"`
}

type UserSecurityReq struct {
	ID        uint64
	OldPasswd string `json:"old_passwd"`
	Passwd    string `json:"passwd"`
}

func (_s *SessionInfo) Otp() bool {
	return _s.OtpSecret != ""
}

func (_s *SessionInfo) IsAdmin() bool {
	return utils.Contains(_s.UserType, models.UserTypeAdmin)
}
