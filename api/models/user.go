package models

const (
	UserCustomer  UserType = "customer" //客户
	UserTypeAdmin UserType = "admin"    //管理员
)

type UserType string

type User struct {
	ID       uint64   `json:"id" gorm:"column:id;primary_key;autoIncrement:false"`
	Alias    string   `json:"alias" gorm:"column:alias;type:VARCHAR(64);index:alias;not null"`
	UserType UserType `json:"user_type" gorm:"column:user_type;type:VARCHAR(64);index:user_type;not null"`
	UpdateAt int64    `json:"update_at" gorm:"column:update_at;not null"`
	UpdateBy uint64   `json:"update_by" gorm:"column:update_by;not null"`
}

func (User) TableName() string {
	return "user"
}

type UserInfo struct {
	User
	Acct          string        `json:"acct" gorm:"column:acct"`                     //账号
	AccountStatus AccountStatus `json:"account_status" gorm:"column:account_status"` //账号状态
	CreatedAt     int64         `json:"create_at" gorm:"column:create_at"`           //创建时间
}
