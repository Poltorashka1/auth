package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Load(name string) {
	err := godotenv.Load(name)
	if err != nil {
		log.Fatal("error loading .env file")
	}
}
func GetEnv(key string) (string, error) {
	val := os.Getenv(key)
	if len(val) == 0 {
		return "", fmt.Errorf("%s not found, check .env file", key)
	}
	return val, nil
}
