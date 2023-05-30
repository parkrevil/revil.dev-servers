package lib

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
	Env            string
	GatewayServer  GatewayServerConfig
	UserGrpcServer UserGrpcServerConfig
	MongoDb        MongoDBConfig
	Redis          RedisConfig
}

type GatewayServerConfig struct {
	Host string
	Port int
}

type UserGrpcServerConfig struct {
	Host string
	Port int
}

type MongoDBConfig struct {
	Uri      string
	Database string
}

type RedisConfig struct {
	Host      string
	Port      int
	Password  string
	LimiterDb int
	AuthDb    int
	CacheDb   int
}

func NewConfig() (*Config, error) {
	var env string

	flag.StringVar(&env, "env", "", "local, production")
	flag.Parse()

	if env != Local && env != Production {
		return nil, errors.New("Invalid argument env")
	}

	if err := godotenv.Load("../.env." + env); err != nil {
		return nil, errors.New("Failed to load env file")
	}

	gatewayServerPort, _ := strconv.Atoi(os.Getenv("GATEWAY_SERVER_PORT"))
	userGrpcServerPort, _ := strconv.Atoi(os.Getenv("USER_GRPC_SERVER_PORT"))
	redisPort, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
	limiterDb, _ := strconv.Atoi(os.Getenv("REDIS_DB_LIMITER"))
	cacheDb, _ := strconv.Atoi(os.Getenv("REDIS_DB_CACHE"))
	authDb, _ := strconv.Atoi(os.Getenv("REDIS_DB_AUTH"))

	return &Config{
		Env: env,
		GatewayServer: GatewayServerConfig{
			Host: os.Getenv("GATEWAY_SERVER_HOST"),
			Port: gatewayServerPort,
		},
		UserGrpcServer: UserGrpcServerConfig{
			Host: os.Getenv("USER_GRPC_SERVER_HOST"),
			Port: userGrpcServerPort,
		},
		MongoDb: MongoDBConfig{
			Uri:      os.Getenv("MONGODB_URI"),
			Database: os.Getenv("MONGODB_DATABASE"),
		},
		Redis: RedisConfig{
			Host:      os.Getenv("REDIS_HOST"),
			Port:      redisPort,
			Password:  os.Getenv("REDIS_PASSWORD"),
			LimiterDb: limiterDb,
			AuthDb:    authDb,
			CacheDb:   cacheDb,
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
