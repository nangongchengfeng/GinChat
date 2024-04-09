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
	// 接受订阅
	pubsub := client.Subscribe("chat")
	defer pubsub.Close()
	for msg := range pubsub.Channel() {
		fmt.Printf("channel=%s message=%s\n", msg.Channel, msg.Payload)
	}
}
