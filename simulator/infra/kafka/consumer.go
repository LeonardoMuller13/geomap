package kafka

import (
	"fmt"
	"log"

	"github.com/LeonardoMuller13/geomap/config"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

func (k *KafkaConsumer) Consume(cfg config.Kafka) {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": cfg.BootstrapServers,
		"group.id":          cfg.Consumer,
	}
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("Error Consuming Kafka message; " + err.Error())
	}
	topics := []string{cfg.ReadTopic}
	c.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been started")
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}
