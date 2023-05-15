package base

import (
	"bytes"

	"github.com/yann0917/knowledge/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	config.DB = newDB()
}

func newDB() (orm *gorm.DB) {
	// var orm *gorm.DB
	var err error

	conf := config.Conf.Mysql
	mysqlLink := bytes.NewBufferString("")

	mysqlLink.WriteString(conf.Username)
	mysqlLink.WriteString(":" + conf.Password + "@tcp")
	mysqlLink.WriteString("(" + conf.Host)
	mysqlLink.WriteString(":" + conf.Port + ")")
	mysqlLink.WriteString("/" + conf.Database)
	mysqlLink.WriteString("?charset=utf8mb4&parseTime=True&loc=Local&timeout=100ms")

	dsn := mysqlLink.String()
	gConf := gorm.Config{
		Logger: nil,
		DryRun: false,
	}
	if config.Conf.Mysql.Dump == "true" {
		gConf.Logger.LogMode(logger.Info)
	}
	orm, err = gorm.Open(mysql.Open(dsn), &gConf)

	if err != nil {
	}
	return
}
