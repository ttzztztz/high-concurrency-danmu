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
		log.Println(err)
		return
	}

	SocketHub.DanmuHub.Broadcast <- &HubBroadcast{
		Content: message.Content,
		Color:   message.Color,
		Uid:     message.Uid,
		Rid:     message.Rid,
	}
}
