package rr

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/dianhadi/roundrobin/internal/config"
	"github.com/dianhadi/roundrobin/internal/entity"
	"github.com/dianhadi/roundrobin/pkg/roundrobin"
)

type mockRoundTripper struct{}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.URL.String() == "http://example-1.com/" || req.URL.String() == "http://example-3.com/" {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "OK"}`)),
		}, nil
	} else {
		return &http.Response{
			StatusCode: 500,
			Body:       ioutil.NopCloser(strings.NewReader(`{"message": "Internal Server Error"}`)),
		}, errors.New("Internal Server Error")
	}
}

func TestHandler_Request(t *testing.T) {
	rrModule := &roundrobin.RoundRobin{}
	rrModule.Init([]string{
		"http://example-1.com",
		"http://example-2.com",
		"http://example-3.com",
	})

	h, err := New(config.Handler{}, rrModule)
	if err != nil {
		t.Errorf("failed to create handler: %v", err)
	}

	httpClient := &http.Client{
		Timeout:   1000 * time.Millisecond,
		Transport: &mockRoundTripper{},
	}

	h.httpModule = httpClient

	payload := entity.Point{Game: "A", GamerID: "1", Points: 1}
	payloadBytes, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(payloadBytes))
	rrw := httptest.NewRecorder()

	// #1 scenario success hit example-1.com
	h.Request(rrw, req)

	if rrw.Code != http.StatusOK {
		t.Errorf("expected status code %d, but got %d", http.StatusOK, rrw.Code)
	}

	// #2 scenario failed hit example-2.com but success hit example-3.com
	h.Request(rrw, req)

	if rrw.Code != http.StatusOK {
		t.Errorf("expected status code %d, but got %d", http.StatusOK, rrw.Code)
	}
}
