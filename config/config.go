package config 

import (
	"fmt"

	"github.com/JIeeiroSst/go-app/repositories/mysql"
	"github.com/JIeeiroSst/go-app/repositories/mongo"
	"github.com/JIeeiroSst/go-app/log"
	"github.com/kelseyhightower/envconfig"
)

type WebConfig struct {
	PORT string     			`envconfig:"WEB_PORT"`
	MysqlConfig mysql.Config 	`envconfig:"WEB_MYSQL"`
	MongoCofig mongo.Mongoconn 	`envconfig:"WEB_MONGO"`
	NATS string 				`envconfig:"WEB_NATS"`
}

var Config WebConfig

func init(){
	err:=envconfig.Process("",&Config)
	if err!=nil{
			log.InitZapLog().Error("no config setup")
	}
	log.InitZapLog().Info(Config.PORT)
	fmt.Println(&Config)
}