package main

import (
	"danmu/controllers/user"
	"danmu/middlewares"
	"danmu/services/initialize"
	"fmt"

	"github.com/gin-gonic/gin"
)

func serverFrontend() {
	fed := gin.Default()
	fed.Static("/", "./frontend/build")
	err := fed.Run(":8889")
	if err != nil {
		fmt.Println("err")
	}
}

func main() {
	initialize.Config()

	server := gin.Default()
	server.Use(middlewares.Cors())
	user.Router(server)

	go serverFrontend()
	err := server.Run(":8888")
	if err != nil {
		fmt.Println(err)
	}
}
