package main

import (
	"danmu/controllers/danmu"
	"danmu/controllers/user"
	"danmu/controllers/ws"
	"danmu/middlewares"
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
	server := gin.Default()
	server.Use(middlewares.Cors())

	user.Router(server)
	danmu.Router(server)
	ws.WebSocket(server)

	go serverFrontend()
	err := server.Run(":8888")
	if err != nil {
		fmt.Println(err)
	}
}
