package main

import (
	"danmu/controllers/user"
	"danmu/middlewares"
	"danmu/services/initialize"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	initialize.Config()

	server := gin.Default()
	server.Use(middlewares.Cors())
	user.Router(server)

	err := server.Run(":8888")
	if err != nil {
		fmt.Println(err)
	}
}
