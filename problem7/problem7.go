package main

import (
	"log"

	"github.com/danilobuerger/euler/primes"
)

// https://projecteuler.net/problem=7
//
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?
func main() {
	n := 10001

	done := make(chan struct{})
	defer close(done)

	g := primes.Generator(done)

	p := 0
	for i := 0; i < n; i++ {
		p = <-g
	}

	log.Println(p)
}
