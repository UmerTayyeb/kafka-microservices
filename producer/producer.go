package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	defer producer.Close()
	topic := "orders"
	order_no := 1120

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed of order n0: %d\n", order_no)
				} else {
					fmt.Printf("Successful delivery of order no: %d\n", order_no)
				}
			}
		}
	}()

	for v := 1; v < 11; v++ {
		message := "order created: order ID: " + strconv.Itoa(order_no)
		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, nil)
		if err != nil {
			fmt.Printf("error producing message with order ID: %d, error: %v", order_no, err)
		} else {
			fmt.Println("message produced!", message)
		}
		time.Sleep(1 * time.Second)
		order_no += 1
	}
	producer.Flush(5000)
}
