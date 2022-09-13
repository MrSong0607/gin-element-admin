package controllers

import (
	"errors"
	"net/http"

	"admin/dto"
	"admin/models"
	"admin/service"
	"admin/utils"

	"github.com/gin-gonic/gin"
)

const SessionKey = "X-SESSION"

var _router = []Controller{}

type Controller interface {
	RouterRegist(r *gin.Engine)
}

type ActionHandler func(*gin.Context) (any, error)

type RequestContext struct {
	ctxID string
}

func RouterRegist(engine *gin.Engine) {
	engine.Use(func(ctx *gin.Context) {
		ctx.Set(string(utils.TraceID), NewRequestContext(ctx))
	})

	for i := range _router {
		_router[i].RouterRegist(engine)
	}
}

func NewRequestContext(ctx *gin.Context) (c *RequestContext) {
	c = &RequestContext{ctxID: utils.Hex(utils.GetUUID()[:4])}
	return
}

func GetRequestContext(ctx *gin.Context) *RequestContext {
	return ctx.MustGet(string(utils.TraceID)).(*RequestContext)
}

func GetSessionInfo(ctx *gin.Context) *dto.SessionInfo {
	return ctx.MustGet(SessionKey).(*dto.SessionInfo)
}

func ApiAuth(userType ...models.UserType) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, e := buildSession(ctx)
		if e != nil {
			ctx.AbortWithStatusJSON(http.StatusOK, dto.ApiResponse{Code: dto.Unauthorized, Message: e.Error()})
			return
		}

		if session.UserType == models.UserTypeAdmin {
			return
		}

		if len(userType) > 0 && !utils.Contains(session.UserType, userType...) { //判断角色权限
			ctx.AbortWithStatusJSON(http.StatusOK, dto.ApiResponse{Code: dto.Forbidden, Message: "权限不足!"})
		}
	}
}

func buildSession(ctx *gin.Context) (session *dto.SessionInfo, e error) {
	val, ok := ctx.Get(SessionKey)
	if ok { //已经加载过session
		session, ok = val.(*dto.SessionInfo)
		if !ok {
			e = service.ErrLoginExpire
		}
	} else if session, e = new(service.Account).LoadSession(ctx.GetHeader(SessionKey)); e == nil { //第一次读取session
		ctx.Set(SessionKey, session)
	}

	return
}

func Otp(ctx *gin.Context) {
	session, e := buildSession(ctx)
	if e != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, dto.ApiResponse{Code: dto.Unauthorized, Message: "登录信息已失效"})
		return
	}

	if !session.Otp() {
		ctx.AbortWithStatusJSON(http.StatusOK, dto.ApiResponse{Code: dto.Forbidden, Message: "请先绑定OTP验证码"})
	}
}

func ApiWrap(do ActionHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, e := do(ctx)
		if e == nil {
			ctx.JSON(http.StatusOK, dto.ApiResponse{Data: res, Code: dto.Success, Message: "success"})
			return
		}

		utils.LogInfo(e.Error())

		if errors.Is(e, dto.ErrAuth) {
			ctx.JSON(http.StatusForbidden, dto.ApiResponse{Code: dto.Forbidden, Message: e.Error()})
		} else {
			ctx.JSON(http.StatusOK, dto.ApiResponse{Code: dto.Unkown, Message: e.Error()})
		}
	}
}

// 访问频率限制
func QpsBlock(key string, count, sec int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, e := service.QpsLimit(key+"_"+c.ClientIP(), count, sec); e == service.ErrFreq {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, dto.ApiResponse{Code: dto.Unkown, Message: "访问速度太快了,请稍后重试!"})
		}
	}
}
