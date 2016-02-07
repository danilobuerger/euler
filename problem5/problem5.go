package main

import (
	"log"
	"math"

	"github.com/danilobuerger/euler/numbers"
	"github.com/danilobuerger/euler/primes"
)

// https://projecteuler.net/problem=5
//
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
func main() {
	k := 20

	done := make(chan struct{})
	defer close(done)

	g := primes.Generator(done)

	n := 1
	limit := int(math.Sqrt(float64(k)))
	for p := range g {
		if p > k {
			break
		}

		a := 1
		if p <= limit {
			a = numbers.GreatestPerfectPower(p, k)
		}

		n *= int(math.Pow(float64(p), float64(a)))
	}

	log.Println(n)
}
