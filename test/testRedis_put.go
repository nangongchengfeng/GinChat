package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.102.20:6379",
		Password: "123456",
		DB:       0,
	})

	// 发布订阅
	n, err := client.Publish("chat", "测阿萨德").Result()

	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Printf("%d clients received the message\n", n)
}
