package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Configure Kafka consumer
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "order-consumers",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// Subscribe to the "orders" topic
	err = consumer.Subscribe("orders", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Consumer is listening on 'orders' topic...")

	// Poll for new messages
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Error consuming message: %v\n", err)
		}
	}
}
