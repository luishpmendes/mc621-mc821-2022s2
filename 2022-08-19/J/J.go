// https://codeforces.com/problemset/problem/1567/A

// 1567A - Domino Disaster

// Alice has a grid with 2 rows and n columns. She fully covers the grid using n dominoes of size 1×2 — Alice may place them vertically or horizontally, and each cell should be covered by exactly one domino.
// Now, she decided to show one row of the grid to Bob. Help Bob and figure out what the other row of the grid looks like!

// Input
// The input consists of multiple test cases. The first line contains an integer t (1≤t≤5000) — the number of test cases. The description of the test cases follows.
// The first line of each test case contains an integer n (1≤n≤100) — the width of the grid.
// The second line of each test case contains a string s consisting of n characters, each of which is either L, R, U, or D, representing the left, right, top, or bottom half of a domino, respectively (see notes for better understanding). This string represents one of the rows of the grid.
// Additional constraint on the input: each input corresponds to at least one valid tiling.

// Output
// For each test case, output one string — the other row of the grid, using the same format as the input string. If there are multiple answers, print any.

// Example
// input
// 4
// 1
// U
// 2
// LR
// 5
// LRDLR
// 6
// UUUUUU
// output
// D
// LR
// LRULR
// DDDDDD

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, t, n, i int
	var s, ans string

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &s)
			ans = s

			for i = 0; i < n; i++ {
				if s[i] == 'U' {
					ans = ans[:i] + "D" + ans[i+1:]
				} else if s[i] == 'D' {
					ans = ans[:i] + "U" + ans[i+1:]
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
