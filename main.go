package main

import (
	"books/controller"
	"books/lib"
	"fmt"

	"github.com/gin-contrib/cors"
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

	route := gin.Default()
	route.Use(cors.Default())

	route.POST("/book/create", ctrl.CreateBook)
	route.GET("/books", ctrl.ReadBooks)
	route.GET("/book/:bookId", ctrl.ReadBook)
	route.PUT("/book/update/:bookId", ctrl.UpdateBook)
	route.DELETE("/book/delete/:bookId", ctrl.DeleteBook)

	route.Run(url)
}
