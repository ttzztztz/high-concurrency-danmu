package db

import (
	"danmu/protobuf"
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
}
