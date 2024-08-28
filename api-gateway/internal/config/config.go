package config

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port              string
	AuthServiceURL    string
	QueryServiceURL   string
	CommandServiceURL string
	TLS               bool
}

const configFile = "./api-gateway/internal/config/envs/dev.env"

func NewConfig() *Config {
	cfg := Config{}

	cfg.setDefualtValues()
	cfg.loadEnv()
	cfg.getFlagConfig()

	return &cfg
}

func (c *Config) setDefualtValues() {
	c.Port = ":30001"
	c.AuthServiceURL = "localhost:30002"
	c.QueryServiceURL = "localhost:30003"
	c.CommandServiceURL = "localhost:30004"
	c.TLS = false
}

func (c *Config) loadEnv() {
	err := godotenv.Load(configFile)
	if err != nil {
		log.Println(err)
	}

	port, ok := os.LookupEnv("PORT")
	if ok {
		c.Port = port
	}

	authService, ok := os.LookupEnv("AUTH_SERVICE_URL")
	if ok {
		c.AuthServiceURL = authService
	}

	queryService, ok := os.LookupEnv("QUERY_SERVICE_URL")
	if ok {
		c.QueryServiceURL = queryService
	}

	commadService, ok := os.LookupEnv("COMMAND_SERVICE_URL")
	if ok {
		c.CommandServiceURL = commadService
	}

	tls, ok := os.LookupEnv("TLS")
	if ok {
		c.TLS, err = strconv.ParseBool(tls)
		if err != nil {
			log.Println("Parse bool error", err)
		}
	}
}

func (c *Config) getFlagConfig() {
	flag.StringVar(&c.Port, "p", c.Port, "API Gateway service port, example :30001")
	flag.StringVar(&c.AuthServiceURL, "a", c.AuthServiceURL, "Auth service address, example localhost:30002")
	flag.StringVar(&c.AuthServiceURL, "q", c.AuthServiceURL, "Query service address, example localhost:30003")
	flag.StringVar(&c.AuthServiceURL, "c", c.AuthServiceURL, "Command service address, example localhost:30004")
	flag.BoolVar(&c.TLS, "t", c.TLS, "Enable TLS gRPC")
	flag.Parse()
}
