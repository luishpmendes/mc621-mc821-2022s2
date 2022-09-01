// https://codeforces.com/problemset/problem/1582/B

// 1582B - Luntik and Subsequences

// Luntik came out for a morning stroll and found an array a of length n. He calculated the sum s of the elements of the array (s=∑ni=1ai). Luntik calls a subsequence of the array a nearly full if the sum of the numbers in that subsequence is equal to s−1.
// Luntik really wants to know the number of nearly full subsequences of the array a. But he needs to come home so he asks you to solve that problem!
// A sequence x is a subsequence of a sequence y if x can be obtained from y by deletion of several (possibly, zero or all) elements.

// Input
// The first line contains a single integer t (1≤t≤1000) — the number of test cases. The next 2⋅t lines contain descriptions of test cases. The description of each test case consists of two lines.
// The first line of each test case contains a single integer n (1≤n≤60) — the length of the array.
// The second line contains n integers a1,a2,…,an (0≤ai≤109) — the elements of the array a.

// Output
// For each test case print the number of nearly full subsequences of the array.

// Example
// input
// 5
// 5
// 1 2 3 4 5
// 2
// 1000 1000
// 2
// 1 0
// 5
// 3 0 2 1 1
// 5
// 2 1 0 3 0
// output
// 1
// 0
// 2
// 4
// 4

// Note

// In the first test case, s=1+2+3+4+5=15, only (2,3,4,5) is a nearly full subsequence among all subsequences, the sum in it is equal to 2+3+4+5=14=15−1.
// In the second test case, there are no nearly full subsequences.
// In the third test case, s=1+0=1, the nearly full subsequences are (0) and () (the sum of an empty subsequence is 0).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, a uint
	var cnt_0, cnt_1 uint64

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			cnt_0, cnt_1 = 0, 0

			for fmt.Fscan(reader, &n); n > 0; n-- {
				fmt.Fscan(reader, &a)

				if a == 0 {
					cnt_0++
				} else if a == 1 {
					cnt_1++
				}
			}

			fmt.Fprintln(writer, (1<<cnt_0)*cnt_1)
		}
	}

	writer.Flush()
}
