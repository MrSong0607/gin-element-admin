package service

import (
	"admin/dao"
	"admin/dto"
	"admin/models"
	"admin/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/gomodule/redigo/redis"
)

type Account struct{}

func (_a Account) Login(p dto.LoginForm) (sessionVal string, session *dto.SessionInfo, e error) {
	num, e := QpsLimit("LOGIN_"+p.Acct, 5, 10*60)
	if e != nil {
		return
	}
	defer func() {
		if e != nil {
			e = fmt.Errorf("%w,剩余:%d次", e, 5-num)
		}
	}()

	_aDao := dao.NewAccount()
	acct, e := _aDao.FindByAcct(p.Acct)
	if e != nil {
		if e != dao.ErrRecordNotFound {
			utils.LogInfo(e)
		} else {
			e = ErrAuthFail
		}
		return
	}

	pwd, e := EncryptPassword(p.Passwd)
	if e != nil {
		return "", nil, ErrAuthFail
	}

	if utils.Hex(pwd) != acct.PassWord {
		return "", nil, ErrAuthFail
	}

	user, e := dao.NewUser().GetByID(acct.ID)
	if e != nil {
		return
	}

	return _a.SaveSession(&acct, &user)
}

func (_a Account) SaveSession(acct *models.Account, user *models.User) (sessionVal string, session *dto.SessionInfo, e error) {
	token := utils.Hex(utils.GetUUID())
	session = &dto.SessionInfo{
		ID:       acct.ID,
		Acct:     acct.Acct,
		Alias:    user.Alias,
		UserType: user.UserType,
		Token:    token,
	}

	str, e := json.MarshalToString(session)
	if e != nil {
		return
	}

	conn := dao.Cache().Get()
	defer conn.Close()

	id := strconv.FormatUint(acct.ID, 10)
	sessionVal = id + ":" + token

	_, e = conn.Do("SET", "SESSION_"+id, str, "EX", 3600*12)

	return
}

func (_a Account) LoadSession(token string) (session *dto.SessionInfo, e error) {
	if len(token) > 100 {
		e = ErrLoginExpire
		return
	}

	index := strings.IndexRune(token, ':')
	if index == len(token)-1 || index < 0 {
		e = ErrLoginExpire
		return
	}

	id := token[:index]
	token = token[index+1:]

	conn := dao.Cache().Get()
	defer conn.Close()

	str, e := redis.String(conn.Do("GET", "SESSION_"+id))
	if e != nil {
		return
	}

	session = &dto.SessionInfo{}
	if e = json.UnmarshalFromString(str, session); e != nil {
		return
	}

	if session.Token != token {
		e = ErrLoginExpire
	}

	return
}

func EncryptPassword(passwd string) (pwdb []byte, e error) {
	pwdb, e = utils.HmacSHA256(SystemSecretKey, []byte(passwd))
	return
}

func (_a Account) GetAccountInfo(id uint64) (models.AccountInfo, error) {
	return dao.NewAccount().GetAccountInfo(id)
}

func (_a Account) UpdateUserInfo(q dto.UserProfileReq, session *dto.SessionInfo) (e error) {
	q.ID = session.ID
	e = dao.NewUser().Update(q)
	return
}

func (_a Account) UpdatePasswd(q dto.UserSecurityReq, session *dto.SessionInfo) (e error) {
	q.ID = session.ID
	adao := dao.NewAccount()

	b, e := EncryptPassword(q.Passwd)
	if e != nil {
		return
	}

	ob, e := EncryptPassword(q.OldPasswd)
	if e != nil {
		return
	}

	info, e := adao.GetByID(session.ID)
	if e != nil {
		return
	}

	if utils.Hex(ob) != info.PassWord {
		e = errors.New("旧密码错误")
		return
	}

	q.Passwd = utils.Hex(b)
	e = adao.UpdateSecurity(q)
	return
}
