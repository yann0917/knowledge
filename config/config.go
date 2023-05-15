package config

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type (
	// redis struct {
	// 	Host     string `default:"192.168.3.3"`
	// 	Port     string `default:"6379"`
	// 	Select   int    `default:"0"`
	// 	Password string
	// 	Prefix   string
	// }

	mysql struct {
		Port     string
		Host     string
		Database string
		Username string
		Password string
		Dump     string
	}
)

type Config struct {
	App struct {
		Name   string
		Mode   string
		Cookie string
	}
	Mysql mysql
}

func init() {
	Viper, _ = InitConfig()
}

func InitConfig() (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	v.SetConfigName("config.yaml")
	if err := v.ReadInConfig(); err == nil {
		log.Printf("use config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&Conf); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&Conf); err != nil {
		fmt.Println(err)
	}
	return v, nil
}
