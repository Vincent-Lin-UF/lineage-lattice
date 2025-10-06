package llt

import "sync"

type Store struct {
	mu sync.RWMutex
	m  map[string][]byte
}

func NewStore() *Store {
	return &Store{
		m: make(map[string][]byte),
	}
}

func (s *Store) Put(key string, val []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = append([]byte(nil), val...)
}

func (s *Store) Get(key string) ([]byte, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.m[key]
	if !ok {
		return nil, false
	}

	out := make([]byte, len(v))
	copy(out, v)
	return out, true
}
