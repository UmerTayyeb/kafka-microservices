#!/bin/bash
echo "running consumer" 
go run consumer/consumer.go &

sleep 5

echo "running producer"
go run producer/producer.go

wait
