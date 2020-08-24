package db

import (
	"danmu/protobuf"
	"danmu/services/rpc"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func PersistDanmu(buf []byte) {
	message := &protobuf.DanmuInternalMessage{}
	err := proto.Unmarshal(buf, message)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("will persist danmu %+v \n", message)

	danmuRequest := protobuf.DanmuRequest{
		Uid:     message.Uid,
		RoomId:  message.Rid,
		Visible: true,
		Content: message.Content,
		Color:   message.Color,
	}

	ds := &rpc.DanmuService{}
	res := &protobuf.DanmuChangeResponse{}
	err = ds.AddDanmu(danmuRequest, res)
	if err != nil {
		fmt.Println(err)
		return
	}
}
