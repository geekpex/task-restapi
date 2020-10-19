package data

import "sync"

// Serial is used to get new serialized int64 value
type Serial struct {
	id int64
	m  sync.Mutex
}

// Get next value for s Serial
func (s *Serial) Get() int64 {
	s.m.Lock()
	defer s.m.Unlock()

	s.id++
	return s.id
}

// Set value id for s Serial
func (s *Serial) Set(id int64) {
	s.m.Lock()
	defer s.m.Unlock()

	s.id = id
}
