package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	tcpAddr := "0.0.0.0:8080"
	listen := os.Getenv("HTTP_FORWARDER_LISTEN_ADDR")
	n := strings.Count(listen, ":")
	if n == 1 {
		tcpAddr = listen
	}

	httpForwarder := NewAsyncHttpForwarder()

	log.Println("HTTP forwarder: https://github.com/andrewbreksa/httpforwarder")
	log.Printf("Listening on tcp address %s", tcpAddr)
	log.Fatal(http.ListenAndServe(tcpAddr, httpForwarder))
}
