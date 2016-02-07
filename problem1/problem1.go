package main

import (
	"log"

	"github.com/danilobuerger/euler/numbers"
)

// https://projecteuler.net/problem=1
//
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
// The sum of these multiples is 23. Find the sum of all the multiples of 3 or 5 below 1000.
func main() {
	limit := 1000 - 1
	sum := numbers.SumMultiples(3, limit) + numbers.SumMultiples(5, limit) - numbers.SumMultiples(15, limit)

	log.Println(sum)
}
