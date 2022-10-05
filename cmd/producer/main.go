package main

import (
	"context"
	"go-rabbit-demo/internal/rabbitmq"
	"time"
)

func main() {
	conn := rabbitmq.CreateConnection("amqp://guest:guest@localhost:5672/")

	ch := rabbitmq.CreateChannel(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rabbitmq.PublishMessage(ch, ctx, "exchange1", "bye", "Hello World!")
}
