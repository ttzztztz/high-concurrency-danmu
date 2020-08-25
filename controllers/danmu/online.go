package danmu

import (
	"danmu/controllers/ws"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Online(c *gin.Context) {
	rid, err := strconv.ParseUint(c.Param("rid"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	online := len(ws.SocketHub.DanmuHub.Rooms[uint32(rid)])
	c.JSON(200, gin.H{
		"code":    200,
		"message": online,
	})
}
