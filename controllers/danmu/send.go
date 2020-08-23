package danmu

import (
	"danmu/models/forms"
	"danmu/services/sensitive"
	"github.com/gin-gonic/gin"
)

const openSensitiveWordFilter = false

func Send(c *gin.Context) {
	sendDanmuForm := forms.SendDanmuForm{}

	if err := c.BindJSON(&sendDanmuForm); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	if openSensitiveWordFilter && !sensitive.AllowPublish(sendDanmuForm.Content) {
		c.JSON(400, gin.H{
			"code":    400,
			"message": "含有敏感词语！",
		})

		return
	}

	// do something...

	c.JSON(200, gin.H{
		"code": 200,
	})
}
