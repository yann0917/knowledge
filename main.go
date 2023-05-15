package main

import (
	"log"

	_ "github.com/yann0917/knowledge/base"
	"github.com/yann0917/knowledge/config"
)

func main() {
	// v, err := config.InitConfig()
	// if err != nil {
	// 	log.Printf("read remote error:%+v\n")
	// }
	log.Println(config.Conf.App.Cookie)
	log.Println(config.DB)

	log.Printf("remote read app.mode=%+v\n", config.Viper.GetString("app.env"))
	log.Printf("remote read db.mysql.url=%+v\n", config.Viper.GetString("mysql.host"))

}
