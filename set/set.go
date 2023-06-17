// Package stack implements a generic set.
// T must be comparable.
package set

import (
	"sync"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}
func (s Set[T]) Insert(key T) {
	s[key] = struct{}{}
}

func (s Set[T]) Exist(key T) bool {
	_, ok := s[key]
	return ok
}

func (s Set[T]) Delete(key T) {
	delete(s, key)
}

type ConcurrencySet[T comparable] struct {
	m sync.RWMutex
	Set[T]
}

func NewConcurrencySet[T comparable]() *ConcurrencySet[T] {
	return &ConcurrencySet[T]{
		Set: Set[T]{},
	}
}

func (s *ConcurrencySet[T]) Insert(key T) {
	s.m.Lock()
	defer s.m.Unlock()
	s.Set[key] = struct{}{}
}

func (s *ConcurrencySet[T]) Exist(key T) bool {
	s.m.RLock()
	defer s.m.RUnlock()
	_, ok := s.Set[key]
	return ok
}

func (s *ConcurrencySet[T]) Delete(key T) {
	s.m.Lock()
	defer s.m.Unlock()
	delete(s.Set, key)
}
