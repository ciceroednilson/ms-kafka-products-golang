package main

import (
	"encoding/json"
	"fmt"

	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"module.mod/consumer/productconsumer"
	"module.mod/consumer/request"
	"module.mod/core/usecases/productusecase"
	"module.mod/infrastructure/repositories/productrepository"
)

func main() {
	productrepository := productrepository.NewRepository()
	productusecase := productusecase.NewProductUseCase(productrepository)
	productconsumer := productconsumer.NewProductConsumer(productusecase)
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
