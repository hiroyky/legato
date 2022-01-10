package main

import (
	"fmt"
	"github.com/hiroyky/legato/infrastructure/config"
	"github.com/hiroyky/legato/interface/router"
)

func main() {
	if err := router.Router.Run(fmt.Sprintf(":%d", config.Env.APIPort)); err != nil {
		panic(err)
	}
}
