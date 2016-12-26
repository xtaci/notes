package main

import (
	"log"

	. "github.com/Shopify/sarama"
)

func main() {
	producer, err := NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	for i := 0; i < 1024*1024; i++ {
		msg := &ProducerMessage{Topic: "test", Key: StringEncoder("userid-1"), Value: StringEncoder("testing 123")}
		_, _, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("FAILED to send message: %s\n", err)
		} else {
			//		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
		}
	}
}
