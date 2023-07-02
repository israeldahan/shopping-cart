package main

import (
	"fmt"
	"github.com/spf13/viper"

	"shoppingcart.com/config"
	"shoppingcart.com/logger"
	"shoppingcart.com/routers"
)

const (
	AppName string = "shoppingcart-service"
	Version string = "0.0.0" // will be overridden in the build process
)

func main() {

	config.InitConfig()
	logger.Init(AppName)
	
	r := routers.InitRouter()
	serverURL := fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port"))
	logger.Info("started url: " + serverURL + " version " + Version)
	r.Run(serverURL)

}