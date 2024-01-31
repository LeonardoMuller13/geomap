package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/LeonardoMuller13/geomap/config"
	route2 "github.com/LeonardoMuller13/geomap/simulator/app/route"
	"github.com/LeonardoMuller13/geomap/simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Produce(msg *ckafka.Message, cfg config.Kafka) {
	producer := kafka.NewKafkaProducer(cfg)
	route := route2.NewRoute()
	json.Unmarshal(msg.Value, &route)
	err := route.LoadPositions()
	if err != nil {
		fmt.Println("Error load positions: " + err.Error())
	}

	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, cfg.ProduceTopic, producer)
		time.Sleep(time.Millisecond * 500)
	}
}
