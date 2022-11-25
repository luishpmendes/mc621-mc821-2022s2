// https://codeforces.com/contest/1257/problem/A

// 1257A - Two Rival Students
// You are the gym teacher in the school.
// There are n students in the row. And there are two rivalling students among them. The first one is in position a, the second in position b. Positions are numbered from 1 to n from left to right.
// Since they are rivals, you want to maximize the distance between them. If students are in positions p and s respectively, then distance between them is |p−s|.
// You can do the following operation at most x times: choose two adjacent (neighbouring) students and swap them.
// Calculate the maximum distance between two rivalling students after at most x swaps.

// Input
// The first line contains one integer t (1≤t≤100) — the number of test cases.
// The only line of each test case contains four integers n, x, a and b (2≤n≤100, 0≤x≤100, 1≤a,b≤n, a≠b) — the number of students in the row, the number of swaps which you can do, and positions of first and second rivaling students respectively.

// Output
// For each test case print one integer — the maximum distance between two rivaling students which you can obtain.

// Example
// input
// 3
// 5 1 3 2
// 100 33 100 1
// 6 0 2 3
// output
// 2
// 99
// 1

// Note

// In the first test case you can swap students in positions 3 and 4. And then the distance between the rivals is equal to |4−2|=2.
// In the second test case you don't have to swap students.
// In the third test case you can't swap students.

// Tutorial

// To solve the problem you need to understand two facts:
//  1. The answer can't be greater than n−1;
//  2. If current distance between rivaling student if less then n−1 we always can increment this distance by one swap;
// In means that answer is equal to min(n−1,|a−b|+x).

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
	var t, n, x, a, b, ans uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &x, &a, &b)

			if a > b {
				ans = a - b + x
			} else {
				ans = b - a + x
			}

			if ans > n-1 {
				ans = n - 1
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
