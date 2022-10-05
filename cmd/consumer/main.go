package main

import (
	"fmt"
	"go-rabbit-demo/internal/config"
	"go-rabbit-demo/internal/rabbitmq"
	"log"
	"os"
)

func main() {
	configObj, err := config.LoadConfig("./cmd/consumer")
	if err != nil {
		log.Fatal("cannot load configObj:", err)
	}

	conn := rabbitmq.CreateConnection(configObj.RabbitUrl)

	ch := rabbitmq.CreateChannel(conn)

	q := rabbitmq.CreateQueue(ch, configObj.QueueName)

	rabbitmq.BindQueue(ch, q, configObj.RoutingKey, configObj.Exchange)

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

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("%s Received a message: %s", hostname, d.Body)
		}
	}()

	log.Printf(" [%s] Waiting for messages. To exit press CTRL+C", hostname)
	<-forever
}
