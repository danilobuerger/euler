package fibonacci

import "testing"

func TestGenerator(t *testing.T) {
	cases := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}

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

func TestEvenGenerator(t *testing.T) {
	cases := []int{2, 8, 34, 144, 610, 2584}

	done := make(chan struct{})
	defer close(done)

	g := EvenGenerator(done)
	for i, c := range cases {
		got := <-g
		if got != c {
			t.Errorf("EvenGenerator()#%d == %d, want %d", i, got, c)
		}
	}
}
