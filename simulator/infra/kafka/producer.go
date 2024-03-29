package kafka

import (
	"log"

	"github.com/LeonardoMuller13/geomap/config"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func NewKafkaProducer(cfg config.Kafka) *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": cfg.BootstrapServers,
	}
	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}
	err := producer.Produce(message, nil)
	if err != nil {
		return err
	}
	return nil
}
