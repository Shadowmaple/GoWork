package main

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

//type RdbClient struct {
//	Self *redis.Client
//}

var Rdb *redis.Client

func main() {
	fmt.Println("sub start")
	Rdb = OpenRedisClient()
	defer Rdb.Close()

	sub := Rdb.Subscribe("channel")

	ch := sub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}

func OpenRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
