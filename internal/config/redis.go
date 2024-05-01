package config

import (
	"fmt"
	"log"
	"strconv"
)

type RedisConfig interface {
	Address() string
	Password() string
	DB() int
}

type redisConfig struct {
	Host     string
	Port     string
	password string
	dB       int
}

func NewRedisConfig() RedisConfig {
	host, err := GetEnv("REDIS_HOST")
	if err != nil {
		log.Fatal(err)
	}

	port, err := GetEnv("REDIS_PORT")
	if err != nil {
		log.Fatal(err)
	}

	password, err := GetEnv("REDIS_PASSWORD")
	if err != nil {
		log.Fatal(err)
	}

	db, err := GetEnv("REDIS_DB")
	if err != nil {
		log.Fatal(err)
	}
	DB, err := strconv.Atoi(db)
	if err != nil {
		log.Fatal(err)
	}

	return &redisConfig{
		Host:     host,
		Port:     port,
		password: password,
		dB:       DB,
	}
}

func (r *redisConfig) Address() string {
	return fmt.Sprintf("%s:%s", r.Host, r.Port)
}

func (r *redisConfig) Password() string {
	return r.password
}

func (r *redisConfig) DB() int {
	return r.dB
}
