package numbers

import "testing"

func TestSumOfMultiples(t *testing.T) {
	i := 5
	n := 10
	want := 15

	if got := SumOfMultiples(i, n); got != want {
		t.Errorf("SumOfMultiples(%d, %d) == %d, want %d", i, n, got, want)
	}
}

func TestSumOfNumbers(t *testing.T) {
	n := 100
	want := 5050

	if got := SumOfNumbers(n); got != want {
		t.Errorf("SumOfNumbers(%d) == %d, want %d", n, got, want)
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
