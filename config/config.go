package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Environment string

const (
	EnvTest       Environment = "test"
	EnvLocal      Environment = "local"
	EnvProduction Environment = "production"
)

type Config struct {
	Environment Environment `required:"true" envconfig:"ENVIRONMENT"`
	Development bool        `required:"true" envconfig:"DEVELOPMENT"`

	// Simulator
	Simulator Simulator

	// Infra
	Kafka Kafka
}

type Kafka struct {
	ReadTopic        string `required:"true" envconfig:"KAFKA_READ_TOPIC"`
	ProduceTopic     string `required:"true" envconfig:"KAFKA_PRODUCE_TOPIC"`
	BootstrapServers string `required:"true" envconfig:"KAFKA_BOOTSTRAP_SERVERS"`
	Consumer         string `required:"true" envconfig:"KAFKA_CONSUMER_SIMULATOR"`
}

type Simulator struct {
	Status bool `required:"true" envconfig:"SIMULATOR_STATUS"`
}

func New() (Config, error) {
	const operation = "Config.New"

	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("%s -> %w", operation, err)
	}

	return cfg, nil
}
