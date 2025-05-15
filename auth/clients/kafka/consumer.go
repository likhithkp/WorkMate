package kafka

import (
	"log"
	"os"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	consumer *kafka.Consumer
	conce    sync.Once
)

func CreateConsumer(host string, groupId string) *kafka.Consumer {
	conce.Do(func() {
		c, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": host,
			"group.id":          groupId,
			"auto.offset.reset": "earliest"})

		if err != nil {
			log.Printf("Failed to create consumer: %s", err)
			os.Exit(1)
		}
		consumer = c
	})

	return consumer
}
