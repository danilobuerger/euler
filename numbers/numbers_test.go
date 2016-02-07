package numbers

import "testing"

func TestSumMultiples(t *testing.T) {
	n := 5
	limit := 10
	want := 15

	if got := SumMultiples(n, limit); got != want {
		t.Errorf("SumMultiples(%d, %d) == %d, want %d", n, limit, got, want)
	}
}

func TestSum1ToN(t *testing.T) {
	n := 100
	want := 5050

	if got := Sum1ToN(n); got != want {
		t.Errorf("Sum1ToN(%d) == %d, want %d", n, got, want)
	}
}
