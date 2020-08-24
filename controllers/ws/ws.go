package ws

import "github.com/gin-gonic/gin"

var (
	SocketHub *Hub
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
