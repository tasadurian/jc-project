package store

import (
	"errors"
	"sync"
)

// Storer ...
type Storer struct {
	data          map[int]string
	totalRequests int64
	latency       float64
	mux           sync.Mutex
}

// NewStore returns a new store.
func NewStore() *Storer {
	db := make(map[int]string)
	return &Storer{
		data:          db,
		totalRequests: 0,
		latency:       0,
	}
}

// Get retrieves the value stored at the given key.
func (s *Storer) Get(key int) (val string, err error) {
	s.mux.Lock()
	val, ok := s.data[key]
	if !ok {
		val = ""
		err = errors.New("key not found")
	}
	s.mux.Unlock()
	return
}

// Put adds a new k:v to the map
func (s *Storer) Put(key int, value string) {
	s.mux.Lock()
	s.data[key] = value
	s.mux.Unlock()
}

// GetAvgLatency ...
func (s *Storer) GetAvgLatency() (float64, error) {
	if s.totalRequests > 0 {
		return s.latency / float64(s.totalRequests), nil
	}
	return 0, errors.New("There has been no requests!")
}

// GetReqCount ...
func (s *Storer) GetReqCount() int64 {
	return s.totalRequests
}

// PutLatency ...
func (s *Storer) PutLatency(latency float64) {
	s.latency += latency
	s.totalRequests++
}
