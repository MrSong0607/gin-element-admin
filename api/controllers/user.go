package controllers

import (
	"admin/dto"
	"admin/models"
	"admin/service"

	"github.com/gin-gonic/gin"
)

type User struct{}

func init() {
	_router = append(_router, User{})
}

func (f User) RouterRegist(r *gin.Engine) {
	g := r.Group("user", ApiAuth(models.UserTypeAdmin))
	{
		g.GET("list", ApiWrap(f.List))

		g.POST("new", ApiWrap(f.Create))
	}
}

func (User) List(c *gin.Context) (r any, e error) {
	p := dto.UserQueryReq{}
	if e = c.ShouldBindQuery(&p); e != nil {
		return
	}

	dt, ct, e := new(service.User).List(p, GetSessionInfo(c))
	if e != nil {
		return
	}

	r = dto.ArrayResult[models.UserInfo]{
		Count: ct,
		List:  dt,
	}
	return
}

func (User) Create(c *gin.Context) (r any, e error) {
	q := dto.UserCreateReq{}
	if e = c.ShouldBindJSON(&q); e != nil {
		return
	}

	if e = q.Valid(); e != nil {
		return
	}

	e = new(service.User).Create(q)
	if e != nil {
		return
	}

	return
}
