package main

import (
	"go-rabbit-demo/internal/rabbitmq"
	"log"
)

func main() {
	conn := rabbitmq.CreateConnection("amqp://guest:guest@172.17.0.2:5672/")

	ch := rabbitmq.CreateChannel(conn)

	q := rabbitmq.CreateQueue(ch, "testQueue1")

	rabbitmq.BindQueue(ch, q, "routingKey1", "exchange1")

	msgs, err := ch.Consume(
		q.Name,           // queue
		"go-rabbit-demo", // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		log.Panic("Failed to register a consumer", err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
