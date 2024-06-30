package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Message struct {
	Key  []byte `json:"response"`
	Data string `json:"requestID"`
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	pubsub := rdb.Subscribe(ctx, "my_channel")
	defer pubsub.Close()

	ch := pubsub.Channel()

	//targetKeys := []string{"response", "requestID", "key3"}

	fmt.Println("Subscribed to channel 'my_channel'. %v", ch)

	for _ = range time.Tick(5 * time.Second) {
		select {
		case msg := <-ch:
			var message Message
			err := json.Unmarshal([]byte(msg.Payload), &message)
			if err != nil {
				fmt.Printf("Failed to unmarshal message: %v", err)
				continue
			}
			fmt.Println("messages '. %s: %s", string(message.Key), message.Data)
			//fmt.Println("messages '. %v", string([]byte(msg.Payload)))
			// for _, key := range targetKeys {
			// 	if message.Key == key {
			// 		fmt.Printf("Processing message: %s\n", message.Data)
			// 	}
			// }
			
			// fmt.Print("Enter input to exit: ")
			// var input string
			// fmt.Scanln(&input)
			// if input == message.Key {
			// 	fmt.Printf("Processing user input message: %s\n", message.Data)
			// }
		}

	}
}
