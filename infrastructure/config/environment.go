package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type environment struct {
	HTTPProtocol  string `envcofig:"HTTP_PROTOCOL"`
	APIHostName   string `envconfig:"API_HOSTNAME"`
	APIPort       int64  `envconfig:"API_PORT"`
	MySQLHostName string `envconfig:"MYSQL_HOST_NAME"`
	MySQLPort     int64  `envconfig:"MYSQL_PORT"`
	MySQLDatabase string `envconfig:"MYSQL_DATABASE"`
	MySQLUserName string `envconfig:"MYSQL_USER_NAME"`
	MySQLPassword string `envconfig:"MYSQL_PASSWORD"`
}

var Env environment

func init() {
	godotenv.Load()

	err := envconfig.Process("", &Env)
	if err != nil {
		panic(err)
	}
}
