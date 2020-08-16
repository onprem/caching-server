package main

import "net/http"

// Basic in-memory store for caching.
type store struct {
	data map[string]response
}

type response struct {
	body    []byte
	headers http.Header
}

func newStore() *store {
	return &store{
		data: make(map[string]response),
	}
}

func (s *store) set(key string, body []byte, headers http.Header) {
	s.data[key] = response{body: body, headers: headers}
}

func (s *store) get(key string) (response, bool) {
	value, ok := s.data[key]
	return value, ok
}
