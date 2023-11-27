package database

import (
	"job-portal-api/config"

	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	cfg := config.GetConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisConfig.Addr,
		Password: cfg.RedisConfig.Password, // no password set
		DB:       cfg.RedisConfig.DB,       // use default DB
	})
	return rdb
}
