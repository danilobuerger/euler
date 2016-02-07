package main

import (
	"log"

	"github.com/danilobuerger/euler/primes"
)

// https://projecteuler.net/problem=10
//
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.
func main() {
	limit := 2000000

	done := make(chan struct{})
	defer close(done)

	g := primes.Generator(done)

	sum := 0
	for p := range g {
		if p >= limit {
			break
		}
		sum += p
	}

	log.Println(sum)
}
