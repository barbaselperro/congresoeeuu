package cacheRedis

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

func RedisDB(senadores interface{}) {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.91:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	valor, _ := json.Marshal(senadores)
	jsn := string(valor)
	err := client.Set(ctx, "senadores", jsn, 0).Err()
	if err != nil {
		panic(err)
	}
}

func Consulta(keys string) string {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.0.91:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	val, _ := client.Get(ctx, keys).Result()
	/*if err != nil {
		panic(err)
	}*/
	return val
}
