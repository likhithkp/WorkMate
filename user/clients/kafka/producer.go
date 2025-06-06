package kafka

import (
	"log"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	producer *kafka.Producer
	ponce    sync.Once
)

func CreateProducer(host string) *kafka.Producer {
	ponce.Do(func() {
		p, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": host,
			"acks":              "all",
		})

		if err != nil {
			log.Fatalf("Failed to create new producer: %v", err)
		}

		go func() {
			for e := range p.Events() {
				switch ev := e.(type) {
				case *kafka.Message:
					if ev.TopicPartition.Error != nil {
						log.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
					} else {
						log.Printf("Produced event to topic %s: key = %s value = %s\n",
							*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
					}
				case *kafka.Error:
					log.Println(ev.Error())
				}
			}
		}()

		producer = p
	})

	return producer
}
