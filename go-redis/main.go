package main

import (
	"github.com/go-redis/redis"
)

func GetRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:1090",
		Password: "",
		DB:       0,
	})
	return client
}

func main() {

}
