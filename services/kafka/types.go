package kafka

import (
	"danmu/controllers/ws"
	"danmu/services/rpc"
	"fmt"
	"github.com/Shopify/sarama"
)

type MessageConsumer struct {
	ready chan bool
}

func (consumer *MessageConsumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *MessageConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *MessageConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("%+v \n", message)
		go ws.PublishDanmu(message.Value)
		session.MarkMessage(message, "")
	}

	return nil
}

type DatabaseConsumer struct {
	ready chan bool
}

func (consumer *DatabaseConsumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *DatabaseConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *DatabaseConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		go rpc.PersistDanmu(message.Value)
		session.MarkMessage(message, "")
	}

	return nil
}
