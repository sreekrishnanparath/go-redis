package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Message struct {
	Key  string `json:"key"`
	Data string `json:"data"`
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Define bulk messages
	messages := []Message{
		{"key1", "data1"},
		{"key2", "data2"},
		{"key3", "data3"},
	}


	var wg sync.WaitGroup
	wg.Add(len(messages))

	go func() {
		for _, message := range messages {
			publish(rdb, message, &wg)
		}
	}()

	wg.Wait()
	fmt.Println("All messages published.")
}

func publish(rdb *redis.Client, message Message, wg *sync.WaitGroup) {
	defer wg.Done()

	msg, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Failed to marshal message: %v", err)
	}

	err = rdb.Publish(ctx, "my_channel", msg).Err()
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	fmt.Printf("Published message: %s\n", msg)
}
