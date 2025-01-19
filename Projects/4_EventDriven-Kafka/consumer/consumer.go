package consumer

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func Consumer() {
	topic := "comments"
	worker, err := ConnectConsumer([]string{"localost:29092"})
	if err != nil {
		log.Fatal("Error connecting to consumer : ", err)
		panic(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal("Error consuming partition : ", err)
		panic(err)
	}

	fmt.Println("Consumer started")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	msgCount := 0

	doneChan := make(chan struct{})

	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)

			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Received message : Message count : %d | Topic : (%s) | Message : %s\n", msgCount, string(msg.Topic), string(msg.Value))
			case <-sigChan:
				fmt.Println("Interruption Detected")
				doneChan <- struct{}{}
			}
		}
	}()

	<-doneChan
	fmt.Printf("Processed %d Messages\n", msgCount)
	if err := worker.Close(); err != nil {
		panic(err)
	}
}

func ConnectConsumer(brokersUrl []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()

	config.Consumer.Return.Errors = true
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
