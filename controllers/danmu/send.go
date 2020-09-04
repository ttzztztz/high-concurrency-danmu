package danmu

import (
	"danmu/models/forms"
	"danmu/protobuf"
	"danmu/services/kafka"
	"danmu/services/sensitive"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

const openSensitiveWordFilter = false

func logStaticData(timestamp int64) {
	diff := time.Now().UnixNano() / 1e6 - timestamp
	log.Printf("[PERFORMANCE] reveived danmu %dms \n", diff)
}

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

	if sendDanmuForm.Time > 0 {
		go logStaticData(sendDanmuForm.Time)
	}

	message := &protobuf.DanmuInternalMessage{
		Uid:     sendDanmuForm.Uid,
		Rid:     sendDanmuForm.Rid,
		Content: sendDanmuForm.Content,
		Color:   sendDanmuForm.Color,
	}

	buf, err := proto.Marshal(message)
	if err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err,
		})

		return
	}

	if err := kafka.PublishMessage("danmu", buf); err != nil {
		c.JSON(400, gin.H{
			"code":    400,
			"message": err,
		})

		return
	}

	c.JSON(200, gin.H{
		"code": 200,
	})
}
