package ws

import (
	"danmu/protobuf"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func PublishDanmu(buf []byte) {
	message := &protobuf.DanmuInternalMessage{}
	err := proto.Unmarshal(buf, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v will publish danmu \n", message)
	SocketHub.DanmuHub.Broadcast <- &HubBroadcast{
		Content: message.Content,
		Color:   message.Color,
		Uid:     message.Uid,
		Rid:     message.Rid,
	}
}
