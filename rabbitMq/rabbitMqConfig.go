package rabbitMq

import (
	"log"
	"github.com/streadway/amqp"
)

type RabbitMQConfig struct {
}

func (cho *RabbitMQConfig) ConfigureHost() (<-chan amqp.Delivery, error) {
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	cho.failOnError(err, "Erro ao conectar no RabbitMQ")
	defer conn.Close()

	// // 📡 Canal
	ch, err := conn.Channel()
	cho.failOnError(err, "Erro ao abrir canal")
	defer ch.Close()

	// // 📬 Declarar fila
	q, err := ch.QueueDeclare(
		"LLM_QUEUE", // nome
		true,        // durável
		false,       // auto-delete
		false,       // exclusiva
		false,       // no-wait
		nil,         // args
	)
	cho.failOnError(err, "Erro ao declarar fila")

	// 👂 Consumir mensagens
	msgs, err := ch.Consume(
		q.Name,
		"",    // consumer
		false, // auto-ack (false = manual)
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	cho.failOnError(err, "Erro ao registrar consumer")
	return msgs, nil
}
func (FoE *RabbitMQConfig) failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// func (hm *RabbitMQConfig) HandleMessage(msgs <-chan *amqp.Delivery) {
// 	for d := range msgs {
// 		log.Printf("📥 Mensagem recebida: %s", d.Body)

// 		factory := &abstractFactory.GetLLMFactory{}
// 		Ifactory := factory.CreateFactory(d.Body)
// 		strategy := Ifactory.CreateStrategy()
// 		strategy.Start(d.Body)

// 		// ⚙️ Processamento da mensagem
// 		err := hm.processMessage(factory, d.Body)
// 		if err != nil {
// 			log.Printf("❌ Erro ao processar: %s", err)
// 			//d.Nack(false, true) // requeue
// 			d.Ack(false)
// 			continue
// 		}

// 		// ✅ Confirma processamento
// 		d.Ack(false)
// 	}

}
