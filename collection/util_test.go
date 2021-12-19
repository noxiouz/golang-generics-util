package collection

import "testing"

func TestEqual(t *testing.T) {
	var cases = []struct {
		A    []int
		B    []int
		want bool
	}{
		{
			A:    []int{1, 4, 7},
			B:    []int{1, 7, 4},
			want: false,
		},
		{
			A:    []int{1, 4, 7},
			B:    []int{1, 4},
			want: false,
		},
		{
			A:    []int{1, 4, 7},
			B:    []int{1, 4, 7},
			want: true,
		},
	}

	for _, cs := range cases {
		if got := Equal(cs.A, cs.B); got != cs.want {
			t.Fatalf("Equal failed: %v %v; got %v, want %v", cs.A, cs.B, got, cs.want)
		}
	}
}
