package xsync

import (
	"sync"
)

// Synchronized serializes access to an underlined value
type Synchronized[T any] struct {
	mu    sync.Mutex
	value *T
}

func NewSynchronized[T any](value *T) *Synchronized[T] {
	return &Synchronized[T]{
		value: value,
	}
}

// WithLock provides an access by pointer
func (s *Synchronized[T]) WithLock(fn func(value *T)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fn(s.value)
}
