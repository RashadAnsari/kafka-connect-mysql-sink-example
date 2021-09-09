package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

type (
	SayHello struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	KafkaMessage struct {
		Schema  json.RawMessage `json:"schema"`
		Payload SayHello        `json:"payload"`
	}
)

func main() {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	brokers := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(brokers, cfg)
	if err != nil {
		log.Fatalf("failed to create producr: %s", err.Error())
	}

	schema, err := os.ReadFile("schema.json")
	if err != nil {
		log.Fatalf("failed to load json schema: %s", err.Error())
	}

	for i := 1; i <= 10000; i++ {
		kafkaMsg := KafkaMessage{
			Schema: json.RawMessage(schema),
			Payload: SayHello{
				ID:   int64(i),
				Name: fmt.Sprintf("Rashad %d", i),
			},
		}

		bytes, err := json.Marshal(&kafkaMsg)
		if err != nil {
			log.Fatalf("failed to marshal json: %s", err.Error())
		}

		msg := &sarama.ProducerMessage{
			Topic: "say-hello",
			Key:   sarama.StringEncoder(fmt.Sprintf("%d", i)),
			Value: sarama.ByteEncoder(bytes),
		}

		if _, _, err := producer.SendMessage(msg); err != nil {
			log.Fatalf("failed to publish message: %s", err.Error())
		}
	}

	log.Println("done!")
}
