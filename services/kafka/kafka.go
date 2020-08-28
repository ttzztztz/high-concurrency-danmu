package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"os"
)

var (
	Version               sarama.KafkaVersion
	Broker                []string
	CancelCtx, CancelFunc = context.WithCancel(context.Background())
)

func init() {
	devEnv := os.Getenv("DEV")
	if devEnv == "1" {
		Broker = []string{
			"kafka:9092",
		}
	} else {
		Broker = []string{
			"localhost:9092",
			"localhost:9093",
			"localhost:9094",
		}
	}

	if version, err := sarama.ParseKafkaVersion("2.6.0"); err != nil {
		panic(err)
	} else {
		Version = version
	}

	Producer()
}
