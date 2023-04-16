package app

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/dianhadi/roundrobin/internal/entity"
)

func (h Handler) Request(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// logging to get machine information, for demo purpose
	host, _ := os.Hostname()
	log.Printf("The request is coming to %s", host)

	var payload entity.Point
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
