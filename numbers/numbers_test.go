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

func TestIsPalindrome(t *testing.T) {
	n := 9009
	want := true

	if got := IsPalindrome(n); got != want {
		t.Errorf("IsPalindrome(%d) == %t, want %t", n, got, want)
	}
}

func TestReverse(t *testing.T) {
	n := 12345
	want := 54321

	if got := Reverse(n); got != want {
		t.Errorf("Reverse(%d) == %d, want %d", n, got, want)
	}
}

func TestGreatestPerfectPower(t *testing.T) {
	limit := 100
	cases := []struct {
		in, want int
	}{
		{2, 6},
		{3, 4},
		{4, 3},
		{5, 2},
	}

	for _, c := range cases {
		got := GreatestPerfectPower(c.in, limit)
		if got != c.want {
			t.Errorf("GreatestPerfectPower(%d, %d) == %d, want %d", c.in, limit, got, c.want)
		}
	}
}
