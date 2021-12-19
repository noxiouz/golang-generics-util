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

func TestMap(t *testing.T) {
	t.Run("MapSome", func(t *testing.T) {
		var data = "somedata"
		var opt = Some(data)

		optInt := Map(opt, func(v string) int {
			return len(v)
		})

		if !optInt.HasValue() {
			t.Fatal("Optional with Int should have a value")
		}

		if got := optInt.Value(); got != len(data) {
			t.Fatalf("OptInt has an unexpected value; want %v, got %v", len(data), got)
		}
	})

	t.Run("MapNone", func(t *testing.T) {
		var opt = None[string]()

		var called = 0
		optInt := Map(opt, func(v string) int {
			called += 1
			return len(v)
		})

		if optInt.HasValue() {
			t.Fatal("Optional with Int should not have a value")
		}
		if called != 0 {
			t.Fatalf("Callback should have not been called")
		}
	})
}

func TestFlatMap(t *testing.T) {
	t.Run("FlatMapSome", func(t *testing.T) {
		var data = "somedata"
		var opt = Some(data)

		optInt := FlatMap(opt, func(v string) Option[int] {
			return Some(len(v))
		})
		if !optInt.HasValue() {
			t.Fatal("Optional with Int should have a value")
		}
		if got := optInt.Value(); got != len(data) {
			t.Fatalf("OptInt has an unexpected value; want %v, got %v", len(data), got)
		}

		noneOptInt := FlatMap(opt, func(string) Option[int] {
			return None[int]()
		})

		if noneOptInt.HasValue() {
			t.Fatal("NoneOptional should not have a value")
		}
	})

	t.Run("FlatMapNone", func(t *testing.T) {
		var opt = None[string]()

		var called = 0
		optInt := FlatMap(opt, func(v string) Option[int] {
			called += 1
			return Some(len(v))
		})

		if optInt.HasValue() {
			t.Fatal("Optional with Int should not have a value")
		}
		if called != 0 {
			t.Fatalf("Callback should have not been called")
		}
	})
}
