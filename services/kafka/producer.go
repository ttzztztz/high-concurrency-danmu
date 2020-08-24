package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"strconv"
	"time"
)

var (
	AsyncProducer sarama.AsyncProducer
)

func Producer() {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.Return.Errors = true
	producer, err := sarama.NewAsyncProducer(Broker, config)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case err := <-AsyncProducer.Errors():
				fmt.Println("[ERROR] ", err)
			}
		}
	}()

	AsyncProducer = producer
}

func PublishMessage(topic string, buf []byte) error {
	strTime := strconv.Itoa(int(time.Now().Unix()))
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(strTime),
		Value: sarama.StringEncoder(buf),
	}
	AsyncProducer.Input() <- msg

	return nil
}
