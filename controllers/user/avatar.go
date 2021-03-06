package user

import (
	"danmu/controllers/auth"
	"danmu/utils/files"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func Avatar(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("uid"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	avatarPath, _ := files.AvatarPath(uint32(uid))

	c.Writer.WriteHeader(200)
	c.Header("Content-Disposition", "attachment; filename=avatar.png")
	c.Header("Content-Type", "application/octet-stream")

	if _, err := os.Stat(avatarPath); err != nil {
		defaultPath, _ := files.DefaultAvatarPath()
		c.File(defaultPath)
	} else {
		c.File(avatarPath)
	}
}

func UploadAvatar(c *gin.Context) {
	authObject, err := auth.GetAuthObj(c)
	if err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"message": err.Error(),
		})

		return
	}

	uid := authObject.Uid
	_, header, err := c.Request.FormFile("avatar")
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	avatarPath, err := files.AvatarPath(uid)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})
		return
	}

	if err := c.SaveUploadedFile(header, avatarPath); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"code": 200,
	})
}
