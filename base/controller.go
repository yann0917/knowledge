package base

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Ctx    *gin.Context
	UserID int64
	// UserInfo        UserInfo
	Permission      bool
	Jwt             string
	Client          string
	ClientID        string
	AppName         string
	AppVersion      string
	AppID           int
	Platform        string
	PhoneModel      string
	PhoneSysVersion string
	UmAliasType     string
	UmDeviceToken   string
}

func NewController(ctx *gin.Context) (ctl *Controller) {
	ctl = &Controller{Ctx: ctx}
	// ctl.getLoginInfo()
	ctl.getHeaders()

	return
}

//
// // 必须通过中间件拿数据
// // 要么后台的用 casbin
// // 要么前台的用 member
// func (ctl *Controller) getLoginInfo() {
// 	c := ctl.Ctx
//
// 	path := c.Request.URL.Path
// 	routeSli := strings.Split(path, "/")
// 	// R(routeSli, "routeSli")
//
// 	// 初始值
// 	ctl.Permission = false
// 	ctl.UserInfo = UserInfo{}
// 	ctl.UserID = 0
//
// 	if len(routeSli) > 2 {
// 		if routeSli[1] == "login" || routeSli[1] == "callback" || routeSli[2] == "login" || routeSli[2] == "callback" {
// 			return
// 		}
// 	} else {
// 		if routeSli[1] == "login" || routeSli[1] == "callback" {
// 			return
// 		}
// 	}
//
// 	if m, ok := c.Get("middleware"); ok && m != nil {
// 		middleware, _ := m.(Middleware)
// 		// R(middleware, "middleware")
// 		ctl.UserID = middleware.UserID
// 		ctl.UserInfo = middleware.UserInfo
// 		ctl.Permission = middleware.Permission
// 	}
// }

func (ctl *Controller) getHeaders() {
	c := ctl.Ctx

	authorization := c.GetHeader("authorization")
	jwt := strings.Replace(authorization, " ", "", -1)
	jwt = strings.TrimPrefix(jwt, "Bearer")

	ctl.Jwt = jwt
	ctl.Client = c.GetHeader("client")
	ctl.ClientID = c.GetHeader("client-id")
	ctl.AppName = c.GetHeader("client-name")
	ctl.AppVersion = c.GetHeader("version")
	ctl.Platform = c.GetHeader("platform")
	ctl.PhoneModel = c.GetHeader("phone-model")
	ctl.PhoneSysVersion = c.GetHeader("phone-sys-version")
	ctl.UmAliasType = c.GetHeader("um-alias-type")
	ctl.UmDeviceToken = c.GetHeader("um-device-token")

	appid := c.GetHeader("appid")
	if appid != "" && appid != "0" {
		ctl.AppID, _ = String2Int(appid)
	} else {
		ctl.AppID = 1
	}
}

// // CheckSecret 检测 ServiceSecret,
// // 内部服务请调用该方法, false 则异常 ERR_UNAUTHORIZED
// func (ctl *Controller) CheckSecret() bool {
// 	c := ctl.Ctx
//
// 	serviceSecret := c.GetHeader("ServiceSecret")
//
// 	if serviceSecret != Config.Service.Secret {
// 		return false
// 	} else {
// 		return true
// 	}
// }

func (ctl *Controller) Success(result interface{}) {
	ctl.Ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "",
		"data":    result,
	})
}

func (ctl *Controller) Error(e error, opt ...interface{}) {
	errors, ok := Errs[e.Error()]
	if !ok {
		errors = &Errors{Code: 0, Message: e.Error()}
	}

	if len(opt) > 0 {
		ctl.Ctx.JSON(http.StatusOK, gin.H{
			"code":    errors.Code,
			"message": errors.Message,
			"data":    opt[0],
		})
	} else {
		ctl.Ctx.JSON(http.StatusOK, gin.H{
			"code":    errors.Code,
			"message": errors.Message,
			"data":    struct{}{},
		})
	}

	ctl.Ctx.Abort()
}
