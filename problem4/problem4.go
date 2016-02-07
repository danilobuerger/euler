package main

import (
	"log"

	"github.com/danilobuerger/euler/numbers"
)

// https://projecteuler.net/problem=4
//
// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
func main() {
	p := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= i; j-- {
			num := i * j
			if num <= p {
				break
			}
			if numbers.IsPalindrome(num) {
				p = num
			}
		}
	}

	log.Println(p)
}
