# Kafka Producer and Consumer in Go

## Overview

This project demonstrates a simple implementation of a Kafka producer and consumer using the Go programming language. The producer generates messages representing order creation and sends them to a Kafka topic, while the consumer listens for those messages and processes them accordingly.

## Features

- **Kafka Producer:** Sends messages to a specified Kafka topic.
- **Kafka Consumer:** Consumes messages from the Kafka topic and processes them.
- **Event-Driven Communication:** Utilizes Apache Kafka for efficient message handling.
- **Error Handling:** Provides feedback on message delivery success or failure.

## Technologies Used

- **Programming Language:** Go
- **Message Broker:** Apache Kafka
- **Library:** [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go/kafka)

## Getting Started

### Prerequisites

- Go (version 1.21 or later)
- Apache Kafka
- Docker (optional, for running Kafka locally)

