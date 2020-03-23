package main

import "github.com/go-redis/redis/v7"

var SubRdb *redis.Client
var PubRdb *redis.Client

func OpenRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
