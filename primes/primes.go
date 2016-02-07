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

// Generator generates prime numbers until closed
func Generator(done chan struct{}) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		select {
		case out <- 2:
		case <-done:
			return
		}

		select {
		case out <- 3:
		case <-done:
			return
		}

		for i := 5; true; i += 6 {
			if IsPrime(i) {
				select {
				case out <- i:
				case <-done:
					return
				}
			}

			if IsPrime(i + 2) {
				select {
				case out <- i + 2:
				case <-done:
					return
				}
			}
		}
	}()

	return out
}

// IsPrime returns true if n is prime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
