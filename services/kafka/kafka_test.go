package kafka

import (
	"danmu/protobuf"
	"fmt"
	"github.com/golang/protobuf/proto"
	"testing"
	"time"
)

func TestProducer(t *testing.T) {
	CreateMessageConsumer([]string{"danmu"}, "1")
	CreateMessageConsumer([]string{"danmu"}, "2")
	CreateMessageConsumer([]string{"danmu"}, "3")

	CreateDatabaseConsumer([]string{"danmu"})
	CreateDatabaseConsumer([]string{"danmu"})

	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("publishing msg")

			testMsg := &protobuf.DanmuInternalMessage{
				Uid:     2,
				Rid:     3,
				Content: "22222",
				Color:   "red",
			}
			buf, _ := proto.Marshal(testMsg)

			if err := PublishMessage("danmu", buf); err != nil {
				fmt.Println(err)
			}

			<-time.After(5 * time.Second)
		}

		done <- true
	}()

	<-done
}
