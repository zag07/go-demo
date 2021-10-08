package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	ByManual()
}

func ByManual() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_1_0
	config.Producer.Partitioner = sarama.NewManualPartitioner

	producer, err := sarama.NewSyncProducer([]string{"localhost:9093"}, config)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Println("Failed to close producer:", err)
		}
	}()

	// Now, we set the Partition field of the ProducerMessage struct.
	msg := &sarama.ProducerMessage{Topic: "test", Partition: 6, Value: sarama.StringEncoder("test")}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalln("Failed to produce message to kafka cluster.")
	}

	if partition != 6 {
		log.Fatal("Message should have been produced to partition 6!")
	}

	log.Printf("Produced message to partition %d with offset %d", partition, offset)
}

func ByPerTopic() {
	config := sarama.NewConfig()
	config.Producer.Partitioner = func(topic string) sarama.Partitioner {
		switch topic {
		case "access_log", "error_log":
			return sarama.NewRandomPartitioner(topic)

		default:
			return sarama.NewHashPartitioner(topic)
		}
	}

	// ...
}

func ByRandom() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_1_0
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Println("Failed to close producer:", err)
		}
	}()

	msg := &sarama.ProducerMessage{Topic: "test", Key: sarama.StringEncoder("key is set"), Value: sarama.StringEncoder("test")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalln("Failed to produce message to kafka cluster.")
	}

	log.Printf("Produced message to partition %d with offset %d", partition, offset)
}
