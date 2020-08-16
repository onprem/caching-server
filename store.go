package main

// Basic in-memory store for caching.
type store struct {
	data map[string]([]byte)
}

func newStore() *store {
	return &store{
		data: make(map[string]([]byte)),
	}
}

func (s *store) set(key string, value []byte) {
	s.data[key] = value
}

func (s *store) get(key string) ([]byte, bool) {
	value, ok := s.data[key]
	return value, ok
}
