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

func TestGenerator(t *testing.T) {
	cases := []int{2, 3, 5, 7, 11, 13, 17, 19}

	done := make(chan struct{})
	defer close(done)

	g := Generator(done)
	for i, c := range cases {
		got := <-g
		if got != c {
			t.Errorf("Generator()#%d == %d, want %d", i, got, c)
		}
	}
}

func TestIsPrime(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
		{11, true},
	}

	for _, c := range cases {
		got := IsPrime(c.in)
		if got != c.want {
			t.Errorf("IsPrime(%d) == %t, want %t", c.in, got, c.want)
		}
	}
}
