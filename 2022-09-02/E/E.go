// https://codeforces.com/problemset/problem/967/B

// 967B - Watering System

// Arkady wants to water his only flower. Unfortunately, he has a very poor watering system that was designed for n flowers and so it looks like a pipe with n holes. Arkady can only use the water that flows from the first hole.
// Arkady can block some of the holes, and then pour A liters of water into the pipe. After that, the water will flow out from the non-blocked holes proportionally to their sizes s1,s2,…,sn. In other words, if the sum of sizes of non-blocked holes is S, and the i-th hole is not blocked, si⋅A/S liters of water will flow out of it.
// What is the minimum number of holes Arkady should block to make at least B liters of water flow out of the first hole?

// Input
// The first line contains three integers n, A, B (1≤n≤100000, 1≤B≤A≤104) — the number of holes, the volume of water Arkady will pour into the system, and the volume he wants to get out of the first hole.
// The second line contains n integers s1,s2,…,sn (1≤si≤104) — the sizes of the holes.

// Output
// Print a single integer — the number of holes Arkady should block.

// Examples

// input
// 4 10 3
// 2 2 2 2
// output
// 1

// input
// 4 80 20
// 3 2 1 4
// output
// 0

// input
// 5 10 10
// 1000 1 1 1 1
// output
// 4

// Note

// In the first example Arkady should block at least one hole. After that, 10⋅2/6≈3.333 liters of water will flow out of the first hole, and that suits Arkady.
// In the second example even without blocking any hole, 80⋅3/10=24 liters will flow out of the first hole, that is not less than 20.
// In the third example Arkady has to block all holes except the first to make all water flow out of the first hole.

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
	var t int
	var n, a, b, s1, i, sum int64
	var s []int64

	for t, _ = fmt.Fscan(reader, &n, &a, &b); t == 3; t, _ = fmt.Fscan(reader, &n, &a, &b) {
		s = make([]int64, n)
		fmt.Fscan(reader, &s1)
		sum = s1

		for i = 1; i < n; i++ {
			fmt.Fscan(reader, &s[i])
			sum += s[i]
		}

		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})

		for i = 0; s1*a < b*sum; i++ {
			sum -= s[n-1-i]
		}

		fmt.Fprintln(writer, i)
	}

	writer.Flush()
}
