package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	Port int64 `envconfig:"PORT"`
}

var Env *Environment

func init() {
	godotenv.Load()

	err := envconfig.Process("", &Env)
	if err != nil {
		panic(err)
	}
}
