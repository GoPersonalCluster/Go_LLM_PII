package main

import (
	"go_llm_pii/rabbitMq"
	"log"
)

func main() {
	rabbitCfg := &rabbitMq.RabbitMQConfig{}
	msgs, err := rabbitCfg.ConfigureHost()
	if err != nil {
		log.Fatal("Erro ao configurar RabbitMQ:", err)
	}

	forever := make(chan bool)

	rabbitCfg.HandleMessage(msgs)

	<-forever
}
