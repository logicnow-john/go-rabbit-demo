package server

import (
	"fmt"
	"go-rabbit-demo/internal/rabbitmq"
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
	typeRequest := req.URL.Query().Get("type")

	conn := rabbitmq.CreateConnection("amqp://guest:guest@localhost:5672/")

	ch := rabbitmq.CreateChannel(conn)

	rabbitmq.PublishMessage(ch, req.Context(), "exchange1", "bye", typeRequest)

	_, _ = fmt.Fprintf(w, "Published: %s\n", typeRequest)
}
