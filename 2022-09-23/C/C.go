// https://codeforces.com/problemset/problem/1025/B

// 1025B - Weakened Common Divisor

// During the research on properties of the greatest common divisor (GCD) of a set of numbers, Ildar, a famous mathematician, introduced a brand new concept of the weakened common divisor (WCD) of a list of pairs of integers.
// For a given list of pairs of integers (a1,b1), (a2,b2), ..., (an,bn) their WCD is arbitrary integer greater than 1, such that it divides at least one element in each pair. WCD may not exist for some lists.
// For example, if the list looks like [(12,15),(25,18),(10,24)], then their WCD can be equal to 2, 3, 5 or 6 (each of these numbers is strictly greater than 1 and divides at least one number in each pair).
// For example, if the list looks like [(12,15),(25,18),(10,24)], then their WCD can be equal to 2, 3, 5 or 6 (each of these numbers is strictly greater than 1 and divides at least one number in each pair).

// Input
// The first line contains a single integer n (1≤n≤150000) — the number of pairs.
// Each of the next n lines contains two integer values ai, bi (2≤ai,bi≤2⋅109).

// Output
// Print a single integer — the WCD of the set of pairs.
// If there are multiple possible answers, output any; if there is no answer, print −1.

// Examples

// input
// 3
// 17 18
// 15 24
// 12 15
// output
// 6

// input
// 2
// 10 16
// 7 17
// output
// -1

// input
// 5
// 90 108
// 45 105
// 75 40
// 165 175
// 33 30
// output
// 5

// Note
// In the first example the answer is 6 since it divides 18 from the first pair, 24 from the second and 12 from the third ones. Note that other valid answers will also be accepted.
// In the second example there are no integers greater than 1 satisfying the conditions.
// In the third example one of the possible answers is 5. Note that, for example, 15 is also allowed, but it's not necessary to maximize the output.

// Tutorial

// The obvious solution of finding all divisors of all numbers in O(n⋅amax−−−−√) definitely won't pass, so we have to come up with a slight optimization.
// One may notice that is's sufficient to take a prime WCD (because if some value is an answer, it can be factorized without losing the answer property). To decrease the number of primes to check, let's get rid of the redundant ones.
// Take an arbitrary pair of elements (ai, bi) and take the union of their prime factorizations. It's easy to see at this point that the cardinality of this union doesn't exceed 2⋅logamax. The rest is to check whether there exists such prime that divides at least one element in each pair in O(n⋅logamax) time.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func sieve(n uint) []uint {
	var primes []uint
	var is_prime []bool
	var i, j uint

	primes = make([]uint, 0)
	is_prime = make([]bool, n+1)

	for i = 2; i <= n; i++ {
		is_prime[i] = true
	}

	for i = 2; i <= n; i++ {
		if is_prime[i] {
			primes = append(primes, i)

			for j = i * i; j <= n; j += i {
				is_prime[j] = false
			}
		}
	}

	return primes
}

func prime_factors(n uint, primes []uint) map[uint]bool {
	var factors map[uint]bool
	var pf_index, pf uint

	factors = make(map[uint]bool)
	pf_index = 0
	pf = primes[pf_index]

	for pf*pf <= n {
		for n%pf == 0 {
			n /= pf
			factors[pf] = true
		}

		pf_index++
		pf = primes[pf_index]
	}

	if n != 1 {
		factors[n] = true
	}

	return factors
}

func union(a, b map[uint]bool) map[uint]bool {
	var c map[uint]bool

	c = make(map[uint]bool)

	for k := range a {
		c[k] = true
	}

	for k := range b {
		c[k] = true
	}

	return c
}

func intersection(a, b map[uint]bool) map[uint]bool {
	var c map[uint]bool

	c = make(map[uint]bool)

	for k := range a {
		if _, ok := b[k]; ok {
			c[k] = true
		}
	}

	return c
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var n, a, b, ans, i, j uint
	var primes, to_delete []uint
	var factors, union_factors map[uint]bool

	primes = sieve(45000)

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		fmt.Fscan(reader, &a, &b)

		union_factors = make(map[uint]bool)
		factors = prime_factors(a, primes)
		union_factors = union(union_factors, factors)
		factors = prime_factors(b, primes)
		union_factors = union(union_factors, factors)

		for i = 1; i < n; i++ {
			fmt.Fscan(reader, &a, &b)

			to_delete = make([]uint, 0)

			for j = range union_factors {
				if a%j != 0 && b%j != 0 {
					to_delete = append(to_delete, j)
				}
			}

			for _, j := range to_delete {
				delete(union_factors, j)
			}
		}

		if len(union_factors) == 0 {
			fmt.Fprintln(writer, -1)
		} else {
			for j := range union_factors {
				ans = j
				break
			}

			for j := range union_factors {
				if ans > j {
					ans = j
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
