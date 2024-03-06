package redis

import "github.com/go-redis/redis/v8"

var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	//defer Rdb.Close()
}
