package primes

import "testing"

func TestFactorizationGenerator(t *testing.T) {
	n := 13195
	cases := []int{5, 7, 13, 29}

	done := make(chan struct{})
	defer close(done)

	g := FactorizationGenerator(n, done)
	for i, c := range cases {
		got := <-g
		if got != c {
			t.Errorf("FactorizationGenerator(%d)#%d == %d, want %d", n, i, got, c)
		}
	}
}
