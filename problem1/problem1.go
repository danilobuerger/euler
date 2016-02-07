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
	n := 1000 - 1
	sum := numbers.SumOfMultiples(3, n) + numbers.SumOfMultiples(5, n) - numbers.SumOfMultiples(15, n)

	log.Println(sum)
}
