package main

import (
	"danmu/controllers/danmu"
	"danmu/controllers/user"
	"danmu/controllers/ws"
	"danmu/middlewares"
	"danmu/services/kafka"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
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

	group, err := os.Hostname()
	if err != nil {
		group = fmt.Sprintf("%d", rand.Int())
	}
	fmt.Printf("Consumer group %s\n", group)

	kafka.CreateMessageConsumer([]string{"danmu"}, group)
	if err := server.Run(":8888"); err != nil {
		fmt.Println(err)
	}
}
