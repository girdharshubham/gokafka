package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Kafka_Broker string
	Kafka_Topic_Consumer string
}

func ConfigReader(fileName string) (kafkaConfig *Config) {
	jsonFile, _ := os.Open(fileName)
	defer jsonFile.Close()
	jsonByte, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(jsonByte, &kafkaConfig)
	return
}