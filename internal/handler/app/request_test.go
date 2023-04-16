package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dianhadi/roundrobin/internal/entity"
)

func TestHandler_Request(t *testing.T) {
	h := Handler{}

	payload := entity.Point{Game: "A", GamerID: "1", Points: 1}
	payloadBytes, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(payloadBytes))

	rr := httptest.NewRecorder()

	h.Request(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			contentType, "application/json")
	}

	expectedResponse, _ := json.Marshal(payload)
	if rr.Body.String() != string(expectedResponse) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), string(expectedResponse))
	}
}
