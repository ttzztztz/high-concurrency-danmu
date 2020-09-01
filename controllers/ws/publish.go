package ws

import (
	"danmu/protobuf"
	"github.com/golang/protobuf/proto"
	"log"
)

func PublishDanmu(buf []byte) {
	message := &protobuf.DanmuInternalMessage{}
	err := proto.Unmarshal(buf, message)
	if err != nil {
		log.Printf("Error when publish danmu %+v \n", err)
		return
	}

	SocketHub.DanmuHub.Broadcast <- &HubBroadcast{
		Content: message.Content,
		Color:   message.Color,
		Uid:     message.Uid,
		Rid:     message.Rid,
	}
}
