package primes

import "math"

// FactorizationGenerator generates prime factors of n until closed
func FactorizationGenerator(n int, done chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n%2 == 0 {
			select {
			case out <- 2:
				n = n / 2
			case <-done:
				return
			}
		}

		max := n
		for i := 3; n > 1 && i <= max; i += 2 {
			for n%i == 0 {
				select {
				case out <- i:
					n = n / i
				case <-done:
					return
				}
			}
			max = int(math.Sqrt(float64(n)))
		}

		if n > 1 {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()

	return out
}
