package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	Ctx         context.Context
	RedisClient *redis.Client
)

func InitRedis() {
	Ctx = context.Background()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if err := RedisClient.Ping(Ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis :%v", err)
	}
	log.Println("Connected to Redis")
}
