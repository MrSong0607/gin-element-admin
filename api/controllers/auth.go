package controllers

import (
	"admin/dto"
	"admin/service"

	"github.com/gin-gonic/gin"
)

type Auth struct{}

func init() {
	_router = append(_router, Auth{})
}

func (f Auth) RouterRegist(r *gin.Engine) {
	g := r.Group("auth")
	{
		g.POST("login", ApiWrap(f.Login))

		g.GET("info", ApiAuth(), ApiWrap(f.Info))

		g.POST("info", ApiAuth(), ApiWrap(f.UpdateInfo))
		g.POST("security", ApiAuth(), ApiWrap(f.Security))
	}
}

func (Auth) Login(c *gin.Context) (r any, e error) {
	p := dto.LoginForm{}
	if e := c.ShouldBindJSON(&p); e != nil {
		return nil, e
	}

	val, _, e := new(service.Account).Login(p)
	if e != nil {
		return
	}

	r = val
	return
}

func (Auth) Info(c *gin.Context) (r any, e error) {
	session := GetSessionInfo(c)
	r, e = new(service.Account).GetAccountInfo(session.ID)
	return
}

func (Auth) UpdateInfo(c *gin.Context) (r any, e error) {
	p := dto.UserProfileReq{}
	if e := c.ShouldBindJSON(&p); e != nil {
		return nil, e
	}

	e = new(service.Account).UpdateUserInfo(p, GetSessionInfo(c))
	return
}

func (Auth) Security(c *gin.Context) (r any, e error) {
	p := dto.UserSecurityReq{}
	if e := c.ShouldBindJSON(&p); e != nil {
		return nil, e
	}

	e = new(service.Account).UpdatePasswd(p, GetSessionInfo(c))
	return
}
