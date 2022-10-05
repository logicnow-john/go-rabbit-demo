package server

import (
	"fmt"
	"go-rabbit-demo/internal/config"
	"go-rabbit-demo/internal/rabbitmq"
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello\n")

	_, _ = fmt.Fprintf(w, "Here are your request headers:\n")
	for name, headers := range req.Header {
		for _, h := range headers {
			_, _ = fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func Result(w http.ResponseWriter, req *http.Request) {
	configObj, err := config.LoadConfig("./cmd/server")
	if err != nil {
		log.Fatal("cannot load configObj:", err)
	}

	typeRequest := req.URL.Query().Get("type")
	routingKey := req.URL.Query().Get("routingKey")

	conn := rabbitmq.CreateConnection(configObj.RabbitUrl)

	ch := rabbitmq.CreateChannel(conn)

	rabbitmq.PublishMessage(ch, req.Context(), configObj.Exchange, routingKey, typeRequest)

	_, _ = fmt.Fprintf(w, "Published: %s\n", typeRequest)
}
