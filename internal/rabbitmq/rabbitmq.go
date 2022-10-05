package rabbitmq

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func CreateConnection(url string) *amqp.Connection {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Panic("Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()
	return conn
}

func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		log.Panic("Failed to open a channel", err)
	}
	defer ch.Close()
	return ch
}

func PublishMessage(ch *amqp.Channel, ctx context.Context, exchange string, routingKey string, body string) {
	err := ch.PublishWithContext(ctx,
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Panic("Failed to publish a message", err)
	}

	log.Printf(" [x] Sent %s\n", body)
}

func CreateQueue(ch *amqp.Channel, name string) amqp.Queue {
	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Panic("Failed to declare a queue", err)
	}

	return q
}

func BindQueue(ch *amqp.Channel, q amqp.Queue, routingKey string, exchange string) {
	err := ch.QueueBind(
		q.Name,     // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,
		nil)

	if err != nil {
		log.Panic("Failed to register a consumer", err)
	}
}
