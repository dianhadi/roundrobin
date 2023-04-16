package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dianhadi/roundrobin/internal/config"
	"github.com/dianhadi/roundrobin/internal/handler/rr"
	"github.com/dianhadi/roundrobin/pkg/roundrobin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Init Config
	log.Println("Init Config")
	cfg, err := config.Init("main-config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Init Module
	log.Println("Init Module")
	rrModule := &roundrobin.RoundRobin{}
	rrModule.Init(cfg.RoundRobin.Instances)

	// Init Handler
	log.Println("Init Handler")
	handlerRR, err := rr.New(cfg.Handler, rrModule)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handlerRR.Request)

	log.Printf("Starting server on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
