package numbers

// SumMultiples returns the sum of multiples of n up until limit
func SumMultiples(n, limit int) int {
	return n * Sum1ToN(limit/n)
}

// Sum1ToN returns the sum from 1 to n, e.g. 1 + 2 + ... + 100 = 5050
func Sum1ToN(n int) int {
	return n * (n + 1) / 2
}
