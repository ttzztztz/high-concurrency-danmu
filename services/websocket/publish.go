package websocket

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
}
