package main

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Message struct {
	UserID    string `json:"user_id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

func hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}
	defer p.Close()

	topic := "go-kafka-topic"

	for i := 0; i < 10; i++ {
		userID := fmt.Sprintf("user-%d", i)
		msg := Message{
			UserID:    userID,
			Message:   fmt.Sprintf("This is message %d", i),
			Timestamp: time.Now().Format(time.RFC3339),
		}

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}

		partition := hash(userID) % 3

		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: int32(partition)},
			Value:          msgBytes,
		}, nil)

		if err != nil {
			panic(err)
		}
	}

	p.Flush(15 * 1000)
}
