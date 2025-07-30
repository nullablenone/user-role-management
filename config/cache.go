package config

import (
	"context"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	db, err := strconv.Atoi(Env.RedisDB)
	if err != nil {
		log.Printf("Warning: REDIS_DB is invalid, using default DB 0. Error: %v", err)
		db = 0
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     Env.RedisAddr,
		Password: Env.RedisPassword,
		DB:       db,
	})

	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Println("Redis connected successfully!")

	return redisClient
}
