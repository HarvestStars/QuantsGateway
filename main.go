package main

import (
	"fmt"

	"github.com/HarvestStars/QuantsGateway/config"
	"github.com/HarvestStars/QuantsGateway/handlerfactory"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("WELCOME TO QUANTS GATEWAY")

	// config server
	config.Setup()

	// init factory
	handlerfactory.Init()

	// register gin server
	r := gin.Default()

	// for gateway post action
	r.POST("/api/v1/account/post", handlerfactory.FactoryImport)

	// server run!
	r.Run(config.ServerSetting.Host)
}
