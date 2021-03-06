package collection

// Option represents an optional value
type Option[T any] interface {
	// HasValue is true if Option contains an actual value
	HasValue() bool
	// Value returns stored value
	// If HasValue is false the result is not defined.
	// It may return zero-inialized T
	Value() T
}

// None is an Option without value
func None[T any]() Option[T] {
	return noneImpl[T]{}
}

type noneImpl[T any] struct{}

func (noneImpl[T]) HasValue() bool {
	return false
}

func (noneImpl[T]) Value() T {
	v := new(T)
	return *v
}

// Some is an Option with Value
func Some[T any](value T) Option[T] {
	return &someImpl[T]{
		value: value,
	}
}

type someImpl[T any] struct {
	value T
}

func (s *someImpl[T]) HasValue() bool {
	return true
}

func (s *someImpl[T]) Value() T {
	return s.value
}

// Map applies a function to an optional value
func Map[T any, U any](opt Option[T], fn func(T) U) Option[U] {
	if !opt.HasValue() {
		return None[U]()
	}

	return Some(fn(opt.Value()))
}

// FlatMap applies a function to an optional value. Similar to Map, but the function
// returns an Option
func FlatMap[T any, U any](opt Option[T], fn func(T) Option[U]) Option[U] {
	if !opt.HasValue() {
		return None[U]()
	}

	return fn(opt.Value())
}
