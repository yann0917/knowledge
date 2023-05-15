package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	Conf  Config
	Viper *viper.Viper
	DB    *gorm.DB
)

const (
	ConfigEnv         = "develop"
	ConfigDefaultFile = "config.yaml"
	ConfigTestFile    = "config.test.yaml"
	ConfigDebugFile   = "config.debug.yaml"
	ConfigReleaseFile = "config.release.yaml"
)
