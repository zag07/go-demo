package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"127.0.0.1:9093"}, config)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	msg := &sarama.ProducerMessage{Topic: "bkc_zag001", Value: sarama.StringEncoder("testing 111")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
}
