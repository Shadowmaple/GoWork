package main

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

type RdbClient struct {
	Self *redis.Client
}

var Rdb *RdbClient

func main() {
	fmt.Println("pub start")
	Rdb = &RdbClient{Self: OpenRedisClient()}
	defer Rdb.Self.Close()

	for {
		var s string
		fmt.Scanf("%s", &s)
		fmt.Println(s)
		err := Rdb.Self.Publish("channel", s).Err()
		if err != nil {
			panic(err)
		}
	}

	// err := Rdb.Self.Set("key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := Rdb.Self.Get("key").Result()
	// if err == redis.Nil {
	// 	fmt.Println("missing_key does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key", val)
	// }
}

// 地址的问题，不可用
/*
func (r *RdbClient) Init() {
	x := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	a := &RdbClient{Self: x}
	r = a
	println(a, r)
}
*/

func OpenRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
