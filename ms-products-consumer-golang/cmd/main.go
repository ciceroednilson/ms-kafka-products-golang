package main

import (
	"encoding/json"
	"fmt"
	"go/consumer/productconsumer"
	"go/consumer/request"
	"go/core/usercases/productusercase"
	_ "go/docs"
	"go/infrastructure/repositories/productrepository"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	productrepository := productrepository.NewRepository()
	productusercase := productusercase.NewProductUserCase(productrepository)
	productconsumer := productconsumer.NewProductConsumer(productusercase)
	startConsumer(productconsumer)
}

func startConsumer(productconsumer productconsumer.ProductConsumer) {
	host := "localhost:9092"
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": host,
		"group.id":          "consumer group - cicero",
		"auto.offset.reset": "smallest"})
	if err != nil {
		panic("Error to consume")
	}

	topics := []string{"topic_save_product"}
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		panic("Error to consume")
	}

	run := true
	for run {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			value := string(e.Value)
			var request request.ProductRequest
			err := json.Unmarshal([]byte(value), &request)
			if err != nil {
				panic(err)
			}
			productconsumer.Create(request)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}
	consumer.Close()
}
