package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"

	"github.com/Shopify/sarama"
)

func main() {
	BySelect()
}

func ByGoroutines() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer([]string{"localhost:9093"}, config)
	if err != nil {
		panic(err)
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg                                  sync.WaitGroup
		enqueued, successes, producerErrors int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			producerErrors++
		}
	}()

ProducerLoop:
	for i := 0; i < 10; i++ {
		val := "zag13 testing " + strconv.Itoa(i)
		message := &sarama.ProducerMessage{Topic: "Shakespeare", Value: sarama.StringEncoder(val)}
		select {
		case producer.Input() <- message:
			enqueued++

		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}

	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d\n", successes, producerErrors)
}

func BySelect() {
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9093"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var enqueued, producerErrors int
ProducerLoop:
	for i := 0; i < 10; i++ {
		msg := "testing" + string(i)
		select {
		case producer.Input() <- &sarama.ProducerMessage{Topic: "Shakespeare", Key: nil, Value: sarama.StringEncoder(msg)}:
			enqueued++
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
			producerErrors++
		case <-signals:
			break ProducerLoop
		}
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, producerErrors)

}
