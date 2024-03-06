package redis

import "github.com/go-redis/redis/v8"

var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "10.2.171.84:6379",
		DB:   0,
	})
	//defer Rdb.Close()
}
