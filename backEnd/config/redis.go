package config

import (
	"github.com/go-redis/redis"
	"go_code/ginStudy/gindemo/backEnd/global"
	"log"
)

func initRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("database config failed: %v", err)
	}

	global.Rdb = RedisClient
}
