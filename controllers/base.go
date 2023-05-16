package controllers

import (
	"github.com/yann0917/knowledge/config"
	"github.com/yann0917/knowledge/service"
)

var svc *service.Service

func init() {
	cookie := config.Conf.App.Cookie
	co := &service.CookieOptions{}
	service.ParseCookies(cookie, &co)
	svc = service.NewService(co)
}
