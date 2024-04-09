package math

import "testing"

func TestMaxNumeric(t *testing.T) {
	tests := []struct {
		a, b, want int
	}{
		{1, 2, 2},
		{2, 1, 2},
		{1, 1, 1},
		{0, 0, 0},
		{-1, -2, -1},
		{-2, -1, -1},
		{-1, -1, -1},
		{0, -1, 0},
		{-1, 0, 0},
	}

	tests2 := []struct {
		a, b, want float32
	}{
		{1.0, 2, 2},
		{2.3, 1, 2.3},
	}

	for _, test := range tests {
		if got := MaxNumeric(test.a, test.b); got != test.want {
			t.Errorf("MaxNumeric(%d, %d) = %d; want %d", test.a, test.b, got, test.want)
		}
	}

	for _, test := range tests2 {
		if got := MaxNumeric(test.a, test.b); got != test.want {
			t.Errorf("MaxNumeric(%f, %f) = %f; want %f", test.a, test.b, got, test.want)
		}
	}
}
