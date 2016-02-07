package main

import (
	"log"

	"github.com/danilobuerger/euler/primes"
)

// https://projecteuler.net/problem=3
//
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
func main() {
	n := 600851475143

	done := make(chan struct{})
	defer close(done)

	g := primes.FactorizationGenerator(n, done)

	p := 0
	for p = range g {
	}

	log.Println(p)
}
