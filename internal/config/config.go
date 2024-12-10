package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server  ServerConfig
	MongoDB MongoDBConfig
	JWT     JWTConfig
}

type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MongoDBConfig struct {
	URI      string
	Database string
	Timeout  time.Duration
}

type JWTConfig struct {
	AccessSecret     string
	RefreshSecret    string
	AccessExpiresIn  time.Duration
	RefreshExpiresIn time.Duration
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Using system environment variables.")
	}

	return &Config{
		Server: ServerConfig{
			Port:         getEnv("SERVER_PORT", "8081"),
			ReadTimeout:  time.Second * 15,
			WriteTimeout: time.Second * 15,
		},
		MongoDB: MongoDBConfig{
			URI:      getEnv("MONGODB_URI", "http://localhost:27017"),
			Database: getEnv("MONGODB_DATABASE", "play-to-win-db"),
			Timeout:  time.Second * 10,
		},
		JWT: JWTConfig{
			AccessSecret:     getEnv("ACCESS_SECRET", "access_secret"),
			RefreshSecret:    getEnv("REFRESH_SECRET", "refresh_secret"),
			AccessExpiresIn:  time.Hour * 24,
			RefreshExpiresIn: time.Hour * 24 * 7,
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
