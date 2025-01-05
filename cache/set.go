package cache

import "sync"

type set[K comparable, V any] struct {
	container map[K]V
	mutex     sync.RWMutex
}

func newSet[K comparable, V any]() *set[K, V] {
	return &set[K, V]{
		container: map[K]V{},
		mutex:     sync.RWMutex{},
	}
}

func (s *set[K, V]) get(key K) (V, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	val, ok := s.container[key]
	return val, ok
}

func (s *set[K, V]) set(key K, val V) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.container[key] = val
}

func (s *set[K, V]) delete(key K) (V, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	val, ok := s.container[key]
	if !ok {
		return val, false
	}

	delete(s.container, key)
	return val, ok
}
