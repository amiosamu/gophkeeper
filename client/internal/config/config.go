package config

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	TLS          bool
	CAfile       string
	ServerAddres string
	Register     bool
	User         string
	Password     string
	Help         bool
}

const configFile = "./client/internal/config/envs/client.env"

func NewConfig() *Config {
	cfg := Config{}

	cfg.setDefualtValues()
	cfg.loadEnv()
	cfg.getFlagConfig()

	return &cfg
}

func (c *Config) setDefualtValues() {
	c.TLS = false
	c.CAfile = "./client/cert.ca"
	c.ServerAddres = "localhost:50001"
}

func (c *Config) loadEnv() {
	err := godotenv.Load(configFile)
	if err != nil {
		log.Println(err)
	}

	tls, ok := os.LookupEnv("TLS_CLIENT")
	if ok {
		c.TLS, err = strconv.ParseBool(tls)
		if err != nil {
			log.Println("Parse bool error:", err)
		}
	}

	caFile, ok := os.LookupEnv("CA_FILE")
	if ok {
		c.CAfile = caFile
	}

	serverAddress, ok := os.LookupEnv("SERVER")
	if ok {
		c.ServerAddres = serverAddress
	}

}

func (c *Config) getFlagConfig() {
	flag.StringVar(&c.CAfile, "c", c.CAfile, "Certificate file")
	flag.StringVar(&c.ServerAddres, "s", c.ServerAddres, "The server address in the format of host:port")
	flag.BoolVar(&c.TLS, "t", c.TLS, "Connection use TLS")

	flag.Parse()
}
