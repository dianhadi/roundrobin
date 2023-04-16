package roundrobin

import (
	"sync"
)

// RoundRobin is a struct that contains a list of resources and a mutex to handle concurrent access
type RoundRobin struct {
	resources []string
	mu        sync.Mutex
	index     int
}

// Init initializes a new RoundRobin instance with a list of resources
func (rr *RoundRobin) Init(resources []string) {
	rr.resources = resources
}

// GetNextResource returns the next host in the list of resources, using a round-robin algorithm
func (rr *RoundRobin) GetNextResource() string {
	rr.mu.Lock()
	defer rr.mu.Unlock()
	peer := rr.resources[rr.index]
	rr.index = (rr.index + 1) % len(rr.resources)
	return peer
}
