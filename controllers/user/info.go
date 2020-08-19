package user

import (
	UserService "danmu/services/user"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	username := c.Param("username")

	user, err := UserService.OtherInfoByUsername(username)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    500,
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code":    200,
			"message": user,
		})
	}
}
