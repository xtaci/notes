package main

import (
	"log"
	"os"
	"os/signal"

	. "github.com/Shopify/sarama"
)

func main() {
	producer, err := NewAsyncProducer([]string{"localhost:9092"}, nil)
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

	msg := &ProducerMessage{Topic: "test", Key: StringEncoder("userid-1"), Value: StringEncoder("testing 123")}
	var enqueued, errors int
ProducerLoop:
	for {
		select {
		case producer.Input() <- msg:
			enqueued++
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
			errors++
		case <-signals:
			break ProducerLoop
		}
	}

	log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}
