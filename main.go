package main

import (
	"fmt"
	"github.com/mind1949/bingo1/router"
	"log"
	"net/http"
)

func main() {

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router.New(),
	}

	fmt.Printf("starting bingo\naddr: %s\n\n", server.Addr)

	log.Fatal(server.ListenAndServe())
}
