// https://www.spoj.com/problems/MAGSUB1/

// MAGSUB1 - Lucifer and Magical Substrings

// Lucifer MorningStar is interested in deepest desires of Zing. Being a programmer Zing said he has a desire of knowing number of magical substrings in a string. A substring of string S is said to be magical if it contains atleast one magical character (A character is magical if its value is prime, and we assign values to characters as:  A is assigned 1 , B is assigned 2......Z is assigned 26). So you have to calculate total number of magical substrings for S in order to help Lucifer who is absolutely newbie in programming, so that he does not disappoint Zing.

// Input
// First line contains number of test cases. ( 1<=T<=10).
// For each case input will contain two lines.
// First line contains length of string N ( 1<=N<=10^5).
// Second line will contain a string S of length N. (String will only contain uppercase letters.)

// Output
// For each test case output a single integer denoting number of magical substrings of S in new line.

// Example
// input
// 1
// 3
// ABC
// output
// 5

// Tutorial

// DP: f(i) = i, if A[i] is prime
//                max_{j < i && A[j] is prime} j, if i == 0 and A[i] is prme
// return sum(f(i) for i in 0:n)
// https://ideone.com/Cg7P4T

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
	var c, t, n, i, result int
	var s string
	var is_prime []bool
	var v []int

	is_prime = sieve(26)

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &s)
			v = make([]int, n)

			if is_prime[s[0]-'A'+1] {
				v[0] = 0
			} else {
				v[0] = -1
			}

			for i = 1; i < n; i++ {
				if is_prime[s[i]-'A'+1] {
					v[i] = i
				} else {
					v[i] = v[i-1]
				}
			}

			result = 0

			for i = 0; i < n; i++ {
				result += v[i] + 1
			}

			fmt.Fprintln(writer, result)
		}
	}

	writer.Flush()
}
