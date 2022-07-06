package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V3_1_0_0

	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9093"}, config)
	if err != nil {
		log.Panicf("Error creating consumer client: %v", err)
	}

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("my_topic", 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			log.Printf(string(msg.Value))
			consumed++
		case <-signals:
			break ConsumerLoop
		}
		time.Sleep(time.Second)
	}

	log.Printf("Consumed: %d\n", consumed)
}
