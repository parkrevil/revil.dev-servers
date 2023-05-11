package main

import (
	"crypto/rand"
	"errors"
	"flag"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	Local      string = "local"
	Production string = "production"
)

type Config struct {
	env    string
	server ServerConfig
	redis  RedisConfig
}

type ServerConfig struct {
	host string
	port int
}

type RedisConfig struct {
	host     string
	port     int
	password string
}

func NewConfig() (*Config, error) {
	var env string

	flag.StringVar(&env, "env", "", "local, production")
	flag.Parse()

	if env != Local && env != Production {
		return nil, errors.New("Invalid argument env")
	}

	if err := godotenv.Load(".env." + env); err != nil {
		return nil, errors.New("Failed to load env file")
	}

	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, err
	}

	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		return nil, err
	}

	return &Config{
		env: env,
		server: ServerConfig{
			host: os.Getenv("SERVER_HOST"),
			port: serverPort,
		},
		redis: RedisConfig{
			host:     os.Getenv("REDIS_HOST"),
			port:     redisPort,
			password: os.Getenv("REDIS_PASSWORD"),
		},
	}, nil
}

func (c *Config) print() {
	n, err := rand.Int(rand.Reader, big.NewInt(1000))
	if err != nil {
		panic(err)
	}

	log.Print(n)
}
