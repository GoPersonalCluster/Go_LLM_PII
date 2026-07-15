package main

import (
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service"
	"github.com/GoPersonalCluster/GO_RabbitMqHandler/app/service/consumer"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/database"
	"github.com/GoPersonalCluster/go_llm_pii/app/internal/pii"
)

func main() {
	database.Connect()
	service := service.RabbitMQConfigComposite{}
	service.ConfigureConnection()

	logCommand := pii.PiiFactory{}

	filterConsumer := consumer.GenericConsumer{}
	config := consumer.ConsumerConfig{}

	config.AbstractFactory = &logCommand
	config.Durable = false
	config.Exclusive = false
	config.AutoDelete = false
	config.NoWait = true
	config.QueueName = "pii_queue"
	config.Args = nil

	filterConsumer.SetConfiguration(&config)

	service.AddConsumer("pii_queue", &filterConsumer)
	service.Start()
}
