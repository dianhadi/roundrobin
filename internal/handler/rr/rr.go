package rr

import (
	"net/http"
	"time"

	"github.com/dianhadi/roundrobin/internal/config"
)

type roundrobinModule interface {
	GetNextResource() string
}

type Handler struct {
	rrModule   roundrobinModule
	httpModule *http.Client
	maxTry     int
}

func New(cfg config.Handler, rr roundrobinModule) (*Handler, error) {
	httpClient := &http.Client{
		Timeout: time.Duration(cfg.MaxTimeout) * time.Millisecond,
	}

	return &Handler{
		rrModule:   rr,
		httpModule: httpClient,
		maxTry:     cfg.MaxRetry,
	}, nil
}
