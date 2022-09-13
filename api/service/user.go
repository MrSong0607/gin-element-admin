package service

import (
	"admin/dao"
	"admin/dto"
	"admin/models"
	"admin/utils"
	"errors"
)

type User struct{}

func (User) List(q dto.UserQueryReq, session *dto.SessionInfo) (rst []models.UserInfo, ct int64, e error) {
	return dao.NewUser().List(q)
}

func (User) Create(q dto.UserCreateReq) (e error) {
	_a, _u := dao.NewAccount(), dao.NewUser()

	if _, e = _a.FindByAcct(q.Acct); e != dao.ErrRecordNotFound {
		if e != nil {
			return e
		} else {
			return errors.New(q.Acct + "账号已存在")
		}
	}

	trans := dao.BeginTransaction(_a, _u)
	defer trans.Rollback()

	acct := &models.Account{Acct: q.Acct, AccountStatus: models.AccountNormal, CreatedBy: 1}
	pwd, e := EncryptPassword(q.PassWord)
	if e != nil {
		return
	}

	acct.PassWord = utils.Hex(pwd)
	if e = _a.Create(acct); e != nil {
		return
	}

	user := &models.User{ID: acct.ID, Alias: q.Alias, UserType: q.UserType}
	if e = _u.Create(user); e != nil {
		return
	}

	e = trans.Commit().Error
	return
}
