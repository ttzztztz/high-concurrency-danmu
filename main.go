package main

import (
	"danmu/controllers/danmu"
	"danmu/controllers/user"
	"danmu/controllers/ws"
	"danmu/middlewares"
	"danmu/services/kafka"
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
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.Use(middlewares.Cors())

	user.Router(server)
	danmu.Router(server)
	ws.WebSocket(server)

	go serverFrontend()
	kafka.CreateMessageConsumer([]string{"danmu"}, "1")
	err := server.Run(":8888")
	if err != nil {
		fmt.Println(err)
	}
}
