package models

import (
	"admin/utils"
	"errors"
)

const (
	AccountNormal  AccountStatus = "NORMAL"
	AccountDisable AccountStatus = "DISABLE"
)

type AccountStatus string

type Account struct {
	ID            uint64        `json:"id" gorm:"column:id;primary_key;autoIncrement"`                         //用户ID
	Acct          string        `json:"acct" gorm:"column:acct;type:VARCHAR(64);uniqueIndex:acct;not null"`    //账号
	PassWord      string        `json:"-" gorm:"column:passwd;type:VARCHAR(128);not null"`                     //密码
	OtpSecret     string        `json:"otp_secret" gorm:"column:otp_secret;type:VARCHAR(64);not null"`         //谷歌验证码
	AccountStatus AccountStatus `json:"account_status" gorm:"column:account_status;type:VARCHAR(64);not null"` //账号状态
	CreatedAt     int64         `json:"create_at" gorm:"column:create_at;autoCreateTime;not null"`             //创建时间
	CreatedBy     uint64        `json:"create_by" gorm:"column:create_by;not null"`                            //创建人
}

type AccountInfo struct {
	ID            uint64        `json:"id" gorm:"column:id;"`
	Acct          string        `json:"acct" gorm:"column:acct;"`
	AccountStatus AccountStatus `json:"account_status" gorm:"column:account_status;"`
	Alias         string        `json:"alias" gorm:"column:alias;"`
	UserType      UserType      `json:"user_type" gorm:"column:user_type;"`
}

func (Account) TableName() string {
	return "account"
}

func (a Account) IsNormal() bool {
	return a.AccountStatus == AccountNormal
}

func (a Account) ValidOtp(code string) bool {
	return utils.ValidOtpCode(a.OtpSecret, code)
}

func (s AccountStatus) Valid() error {
	switch s {
	case AccountNormal, AccountDisable:
		return nil
	default:
		return errors.New("无效账户状态")
	}
}
