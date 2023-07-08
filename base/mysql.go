package base

import (
	"bytes"
	"log"
	"os"
	"time"

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
	if conf.Dump {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: false,
				ParameterizedQueries:      false,
				Colorful:                  true,
			},
		)
		gConf.Logger = newLogger
	}
	orm, err = gorm.Open(mysql.Open(dsn), &gConf)

	if err != nil {
	}
	return
}
