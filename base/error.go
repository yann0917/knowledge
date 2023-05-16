package base

import "errors"

type Errors struct {
	Code    int
	Message string
}

var Errs = map[string]*Errors{
	"ERR_UNKNOW_ERROR":     {0, "未知系统错误"},
	"ERR_NOERROR":          {1, ""},
	"SUCCESS":              {1, ""},
	"ERR_TOO_MANY_REQUEST": {429, "操作过于频繁, 请稍候再试"},
	"ERR_CANNOT_OPERATION": {430, "不可执行此操作"},

	"ERR_PARAM":         {1000, "参数错误"},
	"ERR_UNLOGIN":       {1001, "请先注册登录"},
	"ERR_INVALID_TOKEN": {1002, "无效的token"},
	"ERR_UNAUTHORIZED":  {1004, "您没有权限访问该数据"},
	"ERR_DATA_DECODE":   {1005, "数据解析失败"},
	"ERR_HTTP_BASEURL":  {1006, "请设置 baseurl"},
	"ERR_TCP_TIMEOUT":   {1504, "TCP接口响应超时"},
	"ERR_HTTP_TIMEOUT":  {1505, "HTTP接口响应超时"},

	"ERR_MYSQL":              {2000, "MySQL错误"},
	"ERR_MYSQL_INSTALL_FAIL": {2001, "数据插入失败"},
	"ERR_MYSQL_DELETE_FAIL":  {2002, "数据删除失败"},
	"ERR_MYSQLPOOL_FAIL":     {2004, "mysql连接池丢失"},
	"ERR_REDIS":              {2100, "Reids错误"},
	"ERR_REDISPOOL_FAIL":     {2104, "redis连接池丢失"},
	"ERR_IDGEN_FAIL":         {2404, "id生成失败"},
}

func Excp(errString string) error {
	return errors.New(errString)
}
