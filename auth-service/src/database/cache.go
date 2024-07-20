package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var cache *redis.Client
var ctx = context.Background()

func ConnectCache() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cacheHost := os.Getenv("CACHE_HOST")
	cachePort := os.Getenv("CACHE_PORT")
	cachePassword := os.Getenv("CACHE_PASSWORD")

	cache = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cacheHost, cachePort),
		Password: cachePassword,
		DB:       0,
	})

	fmt.Println("AUTH Service Successfully connected to cache!")
}

func SetCache(key string, val string) error {
	return cache.Set(ctx, key, val, 0).Err()
}

func GetCache(key string) (string, error) {
	return cache.Get(ctx, key).Result()
}

func DeleteCache(key string) (int64, error) {
	return cache.Del(ctx, key).Result()
}