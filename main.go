package main

import (
	"fmt"
	"github.com/legato/infrastructure/config"
	"github.com/legato/interface/router"
)

func main() {
	if err := router.Router.Run(fmt.Sprintf(":%d", config.Env.Port)); err != nil {
		panic(err)
	}
}
