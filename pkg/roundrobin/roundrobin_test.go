package roundrobin

import (
	"sync"
	"testing"
)

func TestRoundRobin_GetNextResource(t *testing.T) {
	rr := RoundRobin{}
	rr.Init([]string{"host1", "host2", "host3"})

	// Call GetNextResource three times and check that it returns hosts in the expected order
	if rr.GetNextResource() != "host1" {
		t.Errorf("GetNextResource() = %s, want %s", rr.GetNextResource(), "host1")
	}
	if rr.GetNextResource() != "host2" {
		t.Errorf("GetNextResource() = %s, want %s", rr.GetNextResource(), "host2")
	}
	if rr.GetNextResource() != "host3" {
		t.Errorf("GetNextResource() = %s, want %s", rr.GetNextResource(), "host3")
	}

	// Call GetNextResource one more time and check that it returns the first host again
	if rr.GetNextResource() != "host1" {
		t.Errorf("GetNextResource() = %s, want %s", rr.GetNextResource(), "host1")
	}
}

func TestRoundRobin_GetNextResource_Concurrency(t *testing.T) {
	rr := RoundRobin{}
	rr.Init([]string{"host1", "host2", "host3"})

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Call GetNextResource from multiple goroutines and store the results in a channel
	ch := make(chan string, 3)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			ch <- rr.GetNextResource()
			wg.Done()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Check that the results in the channel match the expected hosts ("host1", "host2", "host3")
	expectedHosts := []string{"host1", "host2", "host3"}
	for i := 0; i < 3; i++ {
		result := <-ch
		if result != expectedHosts[i] {
			t.Errorf("GetNextResource() = %s, want %s", result, expectedHosts[i])
		}
	}

	// Call GetNextResource one more time and check that it returns the first host again ("host1")
	if rr.GetNextResource() != "host1" {
		t.Errorf("GetNextResource() = %s, want %s", rr.GetNextResource(), "host1")
	}
}
