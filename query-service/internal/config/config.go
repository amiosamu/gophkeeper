package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	DBPostgresURL string
}

const configFile = "./query-service/internal/config/envs/query.env"

func NewConfig() *Config {
	cfg := Config{}

	cfg.setDefualtValues()
	cfg.loadEnv()
	cfg.getFlagConfig()

	return &cfg
}

func (c *Config) setDefualtValues() {
	c.Port = ":60002"
	c.DBPostgresURL = "postgres://<USER>:<PASSWORD>@<HOST>:<PORT>/gophkeeper_db"
}

func (c *Config) loadEnv() {
	err := godotenv.Load(configFile)
	if err != nil {
		log.Println(err)
	}

	port, ok := os.LookupEnv("PORT_QUERY")
	if ok {
		c.Port = port
	}

	dbURL, ok := os.LookupEnv("DB_GOPHKEEPER_URL")
	if ok {
		c.DBPostgresURL = dbURL
	}
}

func (c *Config) getFlagConfig() {
	flag.StringVar(&c.Port, "p", c.Port, "Query service port, example :30001")
	flag.StringVar(&c.DBPostgresURL, "db", c.DBPostgresURL, "DB URL, example postgres://<USER>:<PASSWORD>@<HOST>:<PORT>/auth_db")
	flag.Parse()
}
