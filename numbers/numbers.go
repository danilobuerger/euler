package numbers

import "math"

// SumMultiples returns the sum of multiples of n up until limit
func SumMultiples(n, limit int) int {
	return n * Sum1ToN(limit/n)
}

// Sum1ToN returns the sum from 1 to n, e.g. 1 + 2 + ... + 100 = 5050
func Sum1ToN(n int) int {
	return n * (n + 1) / 2
}

// IsPalindrome returns true if n is a palindrome, e.g. 9009
func IsPalindrome(n int) bool {
	return n == Reverse(n)
}

// Reverse reverses n, e.g. 12345 becomes 54321
func Reverse(n int) int {
	rev := 0
	for n > 0 {
		rev = rev*10 + n%10
		n = n / 10
	}
	return rev
}

// GreatestPerfectPower returns the greatest exponent of base^e < limit
func GreatestPerfectPower(base, limit int) int {
	return int(math.Log(float64(limit)) / math.Log(float64(base)))
}
