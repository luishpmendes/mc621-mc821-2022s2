// https://codeforces.com/problemset/problem/1324/A

// 1324A - Yet Another Tetris Problem

// You are given some Tetris field consisting of n columns. The initial height of the i-th column of the field is ai blocks. On top of these columns you can place only figures of size 2×1 (i.e. the height of this figure is 2 blocks and the width of this figure is 1 block). Note that you cannot rotate these figures.
// Your task is to say if you can clear the whole field by placing such figures.
// More formally, the problem can be described like this:
// The following process occurs while at least one ai is greater than 0:
// 1. You place one figure 2×1 (choose some i from 1 to n and replace ai with ai+2);
// 2. then, while all ai are greater than zero, replace each ai with ai−1.
// And your task is to determine if it is possible to clear the whole field (i.e. finish the described process), choosing the places for new figures properly.
// You have to answer t independent test cases.

// Input
// The first line of the input contains one integer t (1≤t≤100) — the number of test cases.
// The next 2t lines describe test cases. The first line of the test case contains one integer n (1≤n≤100) — the number of columns in the Tetris field. The second line of the test case contains n integers a1,a2,…,an (1≤ai≤100), where ai is the initial height of the i-th column of the Tetris field.

// Output
// For each test case, print the answer — "YES" (without quotes) if you can clear the whole Tetris field and "NO" otherwise.

// Example
// input
// 4
// 3
// 1 1 3
// 4
// 1 1 2 1
// 2
// 11 11
// 1
// 100
// output
// YES
// NO
// YES
// YES

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, t, n, a, prev_a int
	var ans bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)

			ans = true

			for i := 0; i < n; i++ {
				fmt.Fscan(reader, &a)
				a %= 2

				if i > 0 && prev_a != a {
					ans = false
				}

				prev_a = a
			}

			if ans {
				fmt.Fprintln(writer, "YES")
			} else {
				fmt.Fprintln(writer, "NO")
			}
		}
	}

	writer.Flush()
}
