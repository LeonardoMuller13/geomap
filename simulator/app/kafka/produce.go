package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	route2 "github.com/LeonardoMuller13/geomap/simulator/app/route"
	"github.com/LeonardoMuller13/geomap/simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
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
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
