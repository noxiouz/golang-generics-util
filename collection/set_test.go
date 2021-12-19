package collection

import (
	"sort"
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet[int]()
	set.Add(1)
	if !set.Contains(1) {
		t.Fatal("Set is expected to contain 1")
	}
	set.Add(2)
	set.Remove(1)
	if set.Contains(1) {
		t.Fatal("Set should not contain 1")
	}
	set.Add(3)
	set.Add(10)
	var sum = 0
	set.ForEach(func(v int) {
		sum += v
	})

	if sum != 2+3+10 {
		t.Fatal("ForEach was not called on all keys")
	}

	keys := set.Keys()
	sort.Ints(keys)
	if !Equal(keys, []int{2, 3, 10}) {
		t.Fatalf("Returned keys are incorrect; got %v", keys)
	}
}
