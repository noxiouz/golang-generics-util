package collection

// Expected is a type that contains an expected value T
// or an error (unexpected) E
type Expected[T, E any] interface {
	HasValue() bool
	HasError() bool
	Value() T
	Error() E
}

// NewExpected returns an Expected that contains a value
func NewExpected[T, E any](value T) Expected[T, E] {
	return &expectedImpl[T, E]{
		value: value,
	}
}

// NewUnexpected return Expected with "error"
func NewUnexpected[T, E any](err E) Expected[T, E] {
	return &unexpectedImpl[T, E]{
		err: err,
	}
}

type expectedImpl[T, E any] struct {
	value T
}

func (*expectedImpl[T, E]) HasValue() bool {
	return true
}

func (*expectedImpl[T, E]) HasError() bool {
	return false
}

func (e *expectedImpl[T, E]) Value() T {
	return e.value
}

func (*expectedImpl[T, E]) Error() E {
	var err = new(E)
	return *err
}

type unexpectedImpl[T, E any] struct {
	err E
}

func (*unexpectedImpl[T, E]) HasValue() bool {
	return false
}

func (*unexpectedImpl[T, E]) HasError() bool {
	return true
}

func (*unexpectedImpl[T, E]) Value() T {
	var t = new(T)
	return *t
}

func (u *unexpectedImpl[T, E]) Error() E {
	return u.err
}
