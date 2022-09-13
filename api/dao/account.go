package dao

import (
	"admin/dto"
	"admin/models"

	"gorm.io/gorm"
)

type Account struct {
	DB *gorm.DB
}

func NewAccount() *Account {
	return &Account{DB: Db()}
}

func (_a *Account) BeginTransaction(trans *gorm.DB) {
	_a.DB = trans
}

func (_a *Account) FindByAcct(acct string) (account models.Account, e error) {
	e = _a.DB.Where("acct = ?", acct).First(&account).Error
	return
}

func (_a *Account) GetByID(id uint64) (account models.Account, e error) {
	e = _a.DB.Where("id = ?", id).First(&account).Error
	return
}

func (_a *Account) Create(acct *models.Account) (e error) {
	return _a.DB.Create(acct).Error
}

func (_a *Account) GetAccountInfo(id uint64) (info models.AccountInfo, e error) {
	e = _a.DB.Table("account").Joins("left join user on user.id = account.id").
		Where("account.id = ?", id).Select("account.*", "user.alias", "user.user_type", "user.chn_code").First(&info).Error
	return
}

func (_a *Account) UpdateSecurity(q dto.UserSecurityReq) (e error) {
	e = _a.DB.Where("id = ?", q.ID).Select("passwd").Updates(&models.Account{PassWord: q.Passwd}).Error
	return
}
