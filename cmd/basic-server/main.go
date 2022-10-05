package main

import (
	"go-rabbit-demo/internal/server"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.Hello)

	addr := ":8090"

	log.Println("Server starting, listening on: ", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
