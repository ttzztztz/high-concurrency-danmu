package kafka

import (
	"context"
	"github.com/Shopify/sarama"
)

var (
	Version sarama.KafkaVersion
	Broker  = []string{
		"localhost:9092",
		"localhost:9093",
		"localhost:9094",
	}
	CancelCtx, CancelFunc = context.WithCancel(context.Background())
)

func init() {
	if version, err := sarama.ParseKafkaVersion("2.6.0"); err != nil {
		panic(err)
	} else {
		Version = version
	}

	Producer()
}
