package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func CreateMessageConsumer(topics []string, group string) {
	config := sarama.NewConfig()
	config.Version = Version
	consumer := MessageConsumer{
		ready: make(chan bool, 0),
	}

	client, err := sarama.NewConsumerGroup(Broker, group, config)
	if err != nil {
		log.Panicf("Error when creating consumer group: %v", err)
		return
	}

	ctx, _ := context.WithCancel(CancelCtx)
	go func() {
		for {
			fmt.Println("consumer group running", group)

			if err := client.Consume(ctx, topics, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}

			if ctx.Err() != nil {
				log.Panicf("Error from ctx: %v", ctx.Err())
				return
			}

			consumer.ready = make(chan bool, 0)
		}
	}()
}

func CreateDatabaseConsumer(topics []string) {
	config := sarama.NewConfig()
	config.Version = Version
	consumer := DatabaseConsumer{
		ready: make(chan bool, 0),
	}

	client, err := sarama.NewConsumerGroup(Broker, "db", config)
	if err != nil {
		log.Panicf("Error when creating consumer group: %v", err)
		return
	}

	ctx, _ := context.WithCancel(CancelCtx)
	go func() {
		for {
			if err := client.Consume(ctx, topics, &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}

			if ctx.Err() != nil {
				log.Panicf("Error from ctx: %v", ctx.Err())
				return
			}

			consumer.ready = make(chan bool, 0)
		}
	}()
}
