package main

import (
	"fmt"

	"github.com/devilzs1/go-kafka/consumer"
	"github.com/devilzs1/go-kafka/producer"
)

func main() {
	fmt.Println("Learning event driven programming using Kafka & Golang")
	producer.Producer()
	consumer.Consumer()
}