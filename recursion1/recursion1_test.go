package recursion1

import "testing"

func TestFactorial(t *testing.T) {
	var tests = []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
		{8, 40320},
		{12, 479001600},
	}
	for _, test := range tests {
		if got := factorial(test.n); got != test.want {
			t.Errorf("factorial(%v) = %v [wanted: %v]", test.n, got, test.want)
		}
	}
}

func TestBunnyEars2(t *testing.T) {
	var tests = []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 2},
		{2, 5},
		{3, 7},
		{4, 10},
		{5, 12},
		{6, 15},
		{10, 25},
	}
	for _, test := range tests {
		if got := bunnyEars2(test.n); got != test.want {
			t.Errorf("bunnyEars2(%v) = %v [wanted: %v]", test.n, got, test.want)
		}
	}
}
