package util

import "github.com/go-redis/redis/v8"

var rd *redis.Client

func GetRedis() *redis.Client {
	if rd == nil {
		rd = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379", // Redis server address
			Password: "",               // no password set
			DB:       0,                // use default DB
		})
	}

	return rd
}
