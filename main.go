package main

import (
	"books/controller"
	"books/lib"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config := lib.LoadConfig()
	logging := lib.LoadLogging()
	url := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)

	db, err := lib.LoadDatabase()
	if err != nil {
		panic(err)
	}

	ctrl := controller.New(db)

	logging.Info().Msgf("[CONFIG] %v", config)
	logging.Info().Msgf("[APP] Started at http://%s/", url)

	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()

	route.GET("/books", ctrl.ReadBooks)
	route.POST("/books/create", ctrl.CreateBook)

	route.Run(url)
}
