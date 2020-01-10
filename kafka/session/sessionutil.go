package session

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func ConnectionStarter(bootstrapServer string) (consumer *kafka.Consumer) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServer,
		"group.id":          "gokafka",
		"auto.offset.reset": "latest",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(consumer)
	return
}
