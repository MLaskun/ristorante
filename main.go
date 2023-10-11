package main

import (
	"flag"
	"log"

	"github.com/MLaskun/ristorante/api"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "the server address")
	flag.Parse()

	server := api.NewServer(*listenAddr)
	log.Println("Server running on port", *listenAddr)

	log.Fatal(server.Run())
}
