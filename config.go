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
	env     string
	server  ServerConfig
	mongodb MongoDBConfig
	redis   RedisConfig
}

type ServerConfig struct {
	host string
	port int
}

type MongoDBConfig struct {
	uri      string
	database string
}

type RedisConfig struct {
	host      string
	port      int
	password  string
	limiterDb int
	authDb    int
	cacheDb   int
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

	serverPort, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	redisPort, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	limiterDb, _ := strconv.Atoi(os.Getenv("REDIS_DB_LIMITER"))
	cacheDb, _ := strconv.Atoi(os.Getenv("REDIS_DB_CACHE"))
	authDb, _ := strconv.Atoi(os.Getenv("REDIS_DB_AUTH"))

	return &Config{
		env: env,
		server: ServerConfig{
			host: os.Getenv("SERVER_HOST"),
			port: serverPort,
		},
		mongodb: MongoDBConfig{
			uri:      os.Getenv("MONGODB_URI"),
			database: os.Getenv("MONGODB_DATABASE"),
		},
		redis: RedisConfig{
			host:      os.Getenv("REDIS_HOST"),
			port:      redisPort,
			password:  os.Getenv("REDIS_PASSWORD"),
			limiterDb: limiterDb,
			authDb:    authDb,
			cacheDb:   cacheDb,
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
