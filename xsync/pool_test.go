package xsync

import (
	"bytes"
	"runtime/debug"
	"testing"
)

func TestPool(t *testing.T) {
	pool := NewPool(func() *bytes.Buffer {
		return bytes.NewBuffer(make([]byte, 0, 25))
	})

	want := "TEST"
	var buff *bytes.Buffer = pool.Get()
	if buff.Cap() != 25 || buff.Len() != 0 {
		t.Fatal("Capacity and Len must be 0")
	}
	func() {
		val := debug.SetGCPercent(-1)
		defer debug.SetGCPercent(val)

		buff.Reset()
		buff.WriteString(want)
		pool.Put(buff)

		// GC is disabled
		buff = pool.Get()
		if got := buff.String(); got != want {
			t.Fatalf("Buffer should contain data; want %v, got %v", want, got)
		}
	}()
}
