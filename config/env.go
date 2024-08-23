package config

import (
	"github.com/lpernett/godotenv"
	"log"
	"os"
)

var Envs = initConfig()

type Config struct {
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBAddress:  os.Getenv("DB_ADDRESS"),
		DBName:     os.Getenv("DB_NAME"),
	}
}
