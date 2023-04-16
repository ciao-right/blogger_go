package initialize

import (
	"blogger/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func ConnRedis() {
	redis := redis.NewClient(&redis.Options{
		Addr:     global.Viper.GetString("redis.address"),
		Password: "",
		DB:       0,
	})
	result, err := redis.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("ping err :", err)
		return
	}
	fmt.Println(result)
}
