// https://codeforces.com/problemset/gymProblem/102267/B

// 102267B - Primes

// A prime number is a natural number greater than 1 and has exactly 2 divisors which are 1 and the number itself.
// You are given a prime number n, find any 2 prime numbers a and b such that a+b=n or state that no such pair of primes exists.

// Input
// The input contains a single prime number n(2≤n≤10^7).

// Output
// If there doesn't exist any 2 primes such that their summation is equal to n then print -1, otherwise print the 2 primes on a single line separated by a space.

// Examples

// input
// 5
// output
// 2 3

// input
// 11
// output
// -1

// Tutorial

// If n=2 then there's no answer, otherwise n is odd, that means one of the 2 numbers a and b should be even and the other should be odd.
// Since there's only one even prime which is 2, the other number should be equal to n−2, so all we have to check is if n−2 is a prime or not.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func sieve(n uint) []bool {
	var is_prime []bool
	var i, j uint

	is_prime = make([]bool, n+1)

	for i = 2; i <= n; i++ {
		is_prime[i] = true
	}

	for i = 2; i <= n; i++ {
		if is_prime[i] {
			for j = i * i; j <= n; j += i {
				is_prime[j] = false
			}
		}
	}

	return is_prime
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var n uint
	var is_prime []bool

	is_prime = sieve(10001000)

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		if is_prime[n-2] {
			fmt.Fprintln(writer, 2, n-2)
		} else {
			fmt.Fprintln(writer, -1)
		}
	}

	writer.Flush()
}
