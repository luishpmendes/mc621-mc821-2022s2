// https://codeforces.com/problemset/problem/1520/D

// 1520D - Same Differences

// You are given an array a of n integers. Count the number of pairs of indices (i, j) such that i < j and aj − ai = j − i.

// Input
// The first line contains one integer t (1 ≤ t ≤ 10^4). Then t test cases follow.
// The first line of each test case contains one integer n (1 ≤ n ≤ 2⋅10^5).
// The second line of each test case contains n integers a1, a2, …, an (1 ≤ ai ≤ n) — array a.
// It is guaranteed that the sum of n over all test cases does not exceed 2⋅10^5.

// Output
// For each test case output the number of pairs of indices (i, j) such that i < j and aj − ai = j − i.

// Example
// input
// 4
// 6
// 3 5 1 4 6 6
// 3
// 1 2 3
// 4
// 1 3 3 4
// 6
// 1 6 3 4 5 6
// output
// 1
// 3
// 3
// 10

// Tutorial

// Let's rewrite the original equality a bit:
// aj − ai = j − i,
// aj − j = ai − i
// Let's replace each ai with bi = ai − i. Then the answer is the number of pairs (i, j) such that i < j and bi = bj. To calculate this value you can use map or sorting.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, ans, i, j uint
	var a []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			a = make([]uint, n)
			ans = 0

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])
				a[i] -= i
			}

			sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

			for i = 0; i < n; i = j {
				for j = i + 1; j < n && a[i] == a[j]; j++ {
				}
				ans += (j - i) * (j - i - 1) / 2
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
