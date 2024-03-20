package productrepository

import (
	"encoding/json"
	"fmt"
	"go/core/ports"
	"go/infrastructure/entities"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Repository struct {
}

func NewRepository() ports.ProductRepositoryPort {
	return &Repository{}
}

func buildProducer() (*kafka.Producer, error) {
	host := "localhost:9092"
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": host})
	return producer, err
}

func (r *Repository) Save(product entities.Product) error {
	producer, err := buildProducer()
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	json, err := json.Marshal(product)
	if err != nil {
		fmt.Printf("Failed to create json: %s\n", err)
		return err
	}
	topic := "topic_save_product"
	topicPartition := kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}
	message := &kafka.Message{TopicPartition: topicPartition, Value: json}
	delivery_chan := make(chan kafka.Event, 10000)
	err = producer.Produce(message, delivery_chan)
	return err
}
