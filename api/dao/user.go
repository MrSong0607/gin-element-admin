package dao

import (
	"admin/dto"
	"admin/models"

	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser() *User {
	return &User{DB: Db()}
}

func (_u *User) BeginTransaction(trans *gorm.DB) {
	_u.DB = trans
}

func (_u *User) List(q dto.UserQueryReq) (rst []models.UserInfo, ct int64, e error) {
	sql := _u.DB.Table("user").Count(&ct)

	if q.Acct != "" {
		sql = sql.Where("user.id in (?)", _u.DB.Table("account").Where("acct like ?", q.Acct+"%").Select("id"))
	}

	if q.UserType != "" {
		sql = sql.Where("user.user_type = ?", q.UserType)
	}

	e = sql.Scopes(Paginate(q.QueryBase)).
		Joins("left join account a on a.id = user.id").
		Select("user.*,a.acct,a.account_status,a.create_at").Find(&rst).Error

	return
}

func (_u *User) Create(user *models.User) (e error) {
	return _u.DB.Create(user).Error
}

func (_u *User) GetByID(id uint64) (rst models.User, e error) {
	e = _u.DB.First(&rst, "id = ?", id).Error
	return
}

func (_u *User) Update(q dto.UserProfileReq) error {
	return _u.DB.Where("id = ?", q.ID).Select("alias").Updates(models.User{Alias: q.Alias}).Error
}
