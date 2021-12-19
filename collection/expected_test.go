package collection

import (
	"errors"
	"testing"
)

func TestExpected(t *testing.T) {
	type D struct {
		name  string
		value int
	}
	t.Run("Expected", func(t *testing.T) {
		var want = &D{
			name:  "name",
			value: 101,
		}
		var expected Expected[*D, error] = NewExpected[*D, error](want)

		if expected.HasError() {
			t.Fatal("Expected must not have an error")
		}

		if !expected.HasValue() {
			t.Fatal("Expeceted must have a value")
		}

		if got := expected.Value(); got != want {
			t.Fatalf("value is not correct; want %v, got %v", want, got)
		}
	})
	t.Run("Unexpected", func(t *testing.T) {
		var want = errors.New("unexpected error")
		var unexpected Expected[*D, error] = NewUnexpected[*D, error](want)

		if !unexpected.HasError() {
			t.Fatal("Unexpected must have an error")
		}

		if unexpected.HasValue() {
			t.Fatal("Unexpeceted must not have a value")
		}

		if got := unexpected.Error(); got != want {
			t.Fatalf("value is not correct; want %v, got %v", want, got)
		}

	})
}
