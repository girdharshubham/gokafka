package main

import (
	"fmt"
	"gokafka/kafka"
	"gokafka/kafka/session"
	"gokafka/util"
)

func main() {
	kafkaConfig := util.ConfigReader("application.json")
	kafkaConsumer := session.ConnectionStarter(kafkaConfig.Kafka_Broker)
	fmt.Println("================= Consuming =================")
	kafka.Consume(kafkaConsumer,kafkaConfig, -1)

}
