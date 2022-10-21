// https://www.spoj.com/problems/EDIST/

// EDIST - Edit distance

// You are given two strings, A and B. Answer, what is the smallest number of operations you need to
// transform A to B?
// Operations are:
//  1. Delete one letter from one of strings
//  2. Insert one letter into one of strings
//  3. Replace one of letters from one of strings with another letter

// Input
// T - number of test cases
// For each test case:
//  • String A
//  • String B
// Both strings will contain only uppercase characters and they won't be longer than 2000 characters.
// There will be 10 test cases in data set.

// Output
// For each test case, one line, minimum number of operations.

// Example
// input
// 1
// FOOD
// MONEY
// output
// 4

package main

import (
	"bufio"
	"fmt"
	"os"
)

func distance(a, b string) uint {
	var dp [][]uint
	var i, j uint

	dp = make([][]uint, len(a)+1)

	for i = 0; i <= uint(len(a)); i++ {
		dp[i] = make([]uint, len(b)+1)

	}

	for i = 1; i <= uint(len(a)); i++ {
		dp[i][0] = i
	}

	for j = 1; j <= uint(len(b)); j++ {
		dp[0][j] = j
	}

	for j = 1; j <= uint(len(b)); j++ {
		for i = 1; i <= uint(len(a)); i++ {
			dp[i][j] = dp[i-1][j] + 1

			if dp[i][j] > dp[i][j-1]+1 {
				dp[i][j] = dp[i][j-1] + 1
			}

			if a[i-1] == b[j-1] {
				if dp[i][j] > dp[i-1][j-1] {
					dp[i][j] = dp[i-1][j-1]
				}
			} else { // a[i-1] != b[j-1]
				if dp[i][j] > dp[i-1][j-1]+1 {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}
		}
	}

	return dp[len(a)][len(b)]
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t uint
	var a, b string

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &a, &b)
			fmt.Fprintln(writer, distance(a, b))
		}
	}

	writer.Flush()
}
