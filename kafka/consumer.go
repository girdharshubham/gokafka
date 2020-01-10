package kafka

import (
	"fmt"
	"gokafka/util"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func Consume(kafkaConsumer *kafka.Consumer, kafkaConfig *util.Config, timeout int8) {
	kafkaConsumer.Subscribe(kafkaConfig.Kafka_Topic_Consumer, nil)

	for {
		message, err := kafkaConsumer.ReadMessage(-1)
		if err != nil {
			panic("Something Went Wrong")
		}

		fmt.Println(message.Timestamp.String() + ": " + string(message.Value))
	}
}
