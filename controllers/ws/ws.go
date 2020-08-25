package ws

import (
	"github.com/gin-gonic/gin"
	"time"
)

var (
	SocketHub *Hub
)

const (
	WriteWait      = 10 * time.Second
	PongWait       = 60 * time.Second
	PingPeriod     = (PongWait * 9) / 10
	MaxMessageSize = 512
)

type Hub struct {
	DanmuHub *DanmuHub
}

func WebSocket(baseRouter *gin.Engine) {
	go SocketHub.DanmuHub.Run()
	baseRouter.GET("/ws/:uid/:rid", ServeDanmuWs(SocketHub.DanmuHub))
}

func init() {
	SocketHub = &Hub{
		DanmuHub: NewDanmuHub(),
	}
}
