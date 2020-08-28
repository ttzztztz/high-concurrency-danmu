package rpc

import (
	"danmu/models"
	"danmu/protobuf"
	"danmu/services/db"
	"github.com/golang/protobuf/proto"
	"log"
	"time"
)

func PersistDanmu(buf []byte) {
	message := &protobuf.DanmuInternalMessage{}
	err := proto.Unmarshal(buf, message)
	if err != nil {
		log.Println(err)
		return
	}

	danmu := models.Danmu{
		Uid:     message.Uid,
		Rid:     message.Rid,
		Visible: true,
		Content: message.Content,
		Created: time.Now(),
	}

	if _, err = db.DB.Table("danmu").Insert(&danmu); err != nil {
		log.Println(err)
		return
	}
}
