package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dianhadi/roundrobin/internal/handler/app"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Init Handler")
	handlerApp, err := app.New()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handlerApp.Request)

	log.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
