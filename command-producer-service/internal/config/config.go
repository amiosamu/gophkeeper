package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	KafkaURL           string
	CommandConsumerURL string
}

const configFile = "./command-producer-service/internal/config/envs/command-producer.env"

func NewConfig() *Config {
	cfg := Config{}

	cfg.setDefualtValues()
	cfg.loadEnv()
	cfg.getFlagConfig()

	return &cfg
}

func (c *Config) setDefualtValues() {
	c.Port = ":50031"
	c.KafkaURL = ""
	c.CommandConsumerURL = "localhost:50041"
}

func (c *Config) loadEnv() {
	err := godotenv.Load(configFile)
	if err != nil {
		log.Panic(err)
	}

	port, ok := os.LookupEnv("PORT_COMM_PROD")
	if ok {
		c.Port = port
	}

	kafkaURL, ok := os.LookupEnv("KAFKA_URL")
	if ok {
		c.KafkaURL = kafkaURL
	}

	commandConsumerURL, ok := os.LookupEnv("COMM_CONS_URL")
	if ok {
		c.CommandConsumerURL = commandConsumerURL
	}
}

func (c *Config) getFlagConfig() {
	flag.StringVar(&c.Port, "p", c.Port, "Command producer service port")
	flag.StringVar(&c.KafkaURL, "k", c.KafkaURL, "Kafka message broker URL")
	flag.StringVar(&c.CommandConsumerURL, "c", c.CommandConsumerURL, "Command consumer serivce URL")
	flag.Parse()
}
