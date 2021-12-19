package collection

// Set is a map[T]struct{} wrapper with by-value access
type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() Set[T] {
	return Set[T]{
		m: make(map[T]struct{}),
	}
}

// Add adds an element to the set
func (s *Set[T]) Add(key T) {
	s.m[key] = struct{}{}
}

// Contains returns an element is in the set
func (s *Set[T]) Contains(key T) bool {
	_, ok := s.m[key]
	return ok
}

// Remove removes an element from the set
func (s *Set[T]) Remove(key T) {
	delete(s.m, key)
}

// ForEach applies a function to each key
func (s *Set[T]) ForEach(fn func(key T)) {
	for key := range s.m {
		fn(key)
	}
}

// Keys returns all keys
func (s *Set[T]) Keys() []T {
	keys := make([]T, 0, len(s.m))
	for key := range s.m {
		keys = append(keys, key)
	}
	return keys
}
