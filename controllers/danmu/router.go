package danmu

import (
	"github.com/gin-gonic/gin"
)

func Router(baseRouter *gin.Engine) {
	danmuRouter := baseRouter.Group("/danmu")

	danmuRouter.POST("/Send", Send)
}
