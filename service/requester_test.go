package service

import (
	"fmt"
	"os"
	"testing"

	"github.com/yann0917/knowledge/config"
	_ "github.com/yann0917/knowledge/config"
)

var service *Service

func TestMain(m *testing.M) {
	config.InitConfig()
	// cookie := utils.Get(baseURL)
	cookie := "zsxq_access_token=807D1F00-CA38-CCB3-216A-FDFF6E685088_641F5E91EBDB0A8C; abtest_env=product; zsxqsessionid=a7652281175d8be4d1a2a6375391af63"
	// cookie := ""
	co := &CookieOptions{}
	ParseCookies(cookie, &co)
	fmt.Println(co)
	service = NewService(co)
	exitCode := m.Run()
	// 退出
	os.Exit(exitCode)
}

func TestService_GetUserSelf(t *testing.T) {
	service.GetUserSelf()
}
func TestService_GetGroups(t *testing.T) {
	service.GetGroups()
}

func TestService_GetMenus(t *testing.T) {
	service.GetMenus("4224888828")
}

func TestService_GetTopics(t *testing.T) {
	service.GetTopics("4224888828", "digests", "20", "")
}
