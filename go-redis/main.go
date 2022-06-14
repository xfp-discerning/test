package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisdb *redis.Client

func initRedis() (err error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_, err = redisdb.Ping().Result()
	return
}

func main() {
	err := initRedis()
	if err != nil {
		fmt.Printf("connect redis fialed , err :%v\n", err)
		return
	}
	fmt.Printf("connect redis successful\n")

	//zset
	key := "rank"
	items := []redis.Z{
		{Score: 90, Member: "java"},
		{Score: 91, Member: "python"},
		{Score: 92, Member: "C++"},
		{Score: 93, Member: "PHP"},
		{Score: 94, Member: "golang"},
	}
	redisdb.ZAdd(key, items...)

	//the key element will be added score
	f, err := redisdb.ZIncrBy(key, 10, "golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed ,err:%v\n",err)
		return
	}
	fmt.Printf("current score of golang is %f\n",f)
}
