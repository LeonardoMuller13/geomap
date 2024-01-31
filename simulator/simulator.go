package simulator

import (
	"fmt"

	"github.com/LeonardoMuller13/geomap/config"
	kafka2 "github.com/LeonardoMuller13/geomap/simulator/app/kafka"
	"github.com/LeonardoMuller13/geomap/simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func New(cfg config.Kafka) {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume(cfg)

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg, cfg)
	}
}
