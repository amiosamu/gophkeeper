package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	KafkaURL      string
	DBPostgresURL string
}

const configFile = "./command-consumer-service/internal/config/envs/command-consumer.env"

func NewConfig() *Config {
	cfg := Config{}

	cfg.setDefualtValues()
	cfg.loadEnv()
	cfg.getFlagConfig()

	return &cfg
}

func (c *Config) setDefualtValues() {
	c.Port = ":50041"
	c.KafkaURL = ""
	c.DBPostgresURL = "postgres://<USER>:<PASSWORD>@<HOST>:<PORT>/gophkeeper_db"
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

	dbPostgres, ok := os.LookupEnv("DB_GOPHKEEPER_URL")
	if ok {
		c.DBPostgresURL = dbPostgres
	}
}

func (c *Config) getFlagConfig() {
	flag.StringVar(&c.Port, "p", c.Port, "Command producer service port")
	flag.StringVar(&c.KafkaURL, "k", c.KafkaURL, "Kafka message broker URL")
	flag.StringVar(&c.DBPostgresURL, "d", c.DBPostgresURL, "DB URL, example postgres://<USER>:<PASSWORD>@<HOST>:<PORT>/auth_db")
	flag.Parse()
}
