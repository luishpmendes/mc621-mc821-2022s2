// https://codeforces.com/problemset/problem/1699/B

// 1699B - Almost Ternary Matrix

// You are given two even integers n and m. Your task is to find any binary matrix a with n rows and m columns where every cell (i,j) has exactly two neighbours with a different value than ai,j.
// Two cells in the matrix are considered neighbours if and only if they share a side. More formally, the neighbours of cell (x,y) are: (x−1,y), (x,y+1), (x+1,y) and (x,y−1).
// It can be proven that under the given constraints, an answer always exists.

// Input
// Each test contains multiple test cases. The first line of input contains a single integer t (1≤t≤100) — the number of test cases. The following lines contain the descriptions of the test cases.
// The only line of each test case contains two even integers n and m (2≤n,m≤50) — the height and width of the binary matrix, respectively.

// Output
// For each test case, print n lines, each of which contains m numbers, equal to 0 or 1 — any binary matrix which satisfies the constraints described in the statement.
// It can be proven that under the given constraints, an answer always exists.

// Example
// input
// 3
// 2 4
// 2 2
// 4 4
// output
// 1 0 0 1
// 0 1 1 0
// 1 0
// 0 1
// 1 0 1 0
// 0 0 1 1
// 1 1 0 0
// 0 1 0 1

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, t, n, m, i, j int

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &m)

			for i = 1; i <= n; i++ {
				for j = 1; j <= m; j++ {
					if (i%4 <= 1 && j%4 > 1) || (i%4 > 1 && j%4 <= 1) {
						fmt.Fprint(writer, 1)
					} else {
						fmt.Fprint(writer, 0)
					}

					fmt.Fprint(writer, " ")
				}

				fmt.Fprintln(writer)
			}
		}
	}

	writer.Flush()
}
