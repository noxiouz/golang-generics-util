package collection

import "testing"

func TestNone(t *testing.T) {
	t.Run("NoneOfInt", func(t *testing.T) {
		var v Option[int] = None[int]()
		var zeroInt int

		if v.HasValue() {
			t.Fatal("None is expected to have no value")
		}

		if val := v.Value(); val != zeroInt {
			t.Fatalf("Invalid zero value; wanted %v; got %v", zeroInt, val)
		}
	})

	t.Run("NoneOfPointer", func(t *testing.T) {
		type D struct {
			name  string
			value int
		}
		var v Option[*D] = None[*D]()
		var zeroD *D

		if v.HasValue() {
			t.Fatal("None is expected to have no value")
		}

		if val := v.Value(); val != zeroD {
			t.Fatalf("Invalid zero value; wanted %v; got %v", zeroD, val)
		}
	})
}

func TestSome(t *testing.T) {
	t.Run("SomeOfInt", func(t *testing.T) {
		var want = 100
		var v Option[int] = Some(want)

		if !v.HasValue() {
			t.Fatal("Some is expected to have no value")
		}

		if got := v.Value(); got != want {
			t.Fatalf("Invalid zero value; want %v; got %v", want, got)
		}
	})

	t.Run("NoneOfPointer", func(t *testing.T) {
		type D struct {
			name  string
			value int
		}
		var want = &D{
			name:  "name",
			value: 100,
		}
		var v Option[*D] = Some(want)

		if !v.HasValue() {
			t.Fatal("Some is expected to have no value")
		}

		if val := v.Value(); val != want {
			t.Fatalf("Invalid zero value; wanted %v; got %v", want, val)
		}
	})

}
