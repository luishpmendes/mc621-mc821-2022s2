// https://codeforces.com/problemset/problem/844/B

// 844B - Rectangles

// You are given n × m table. Each cell of the table is colored white or black. Find the number of non-empty sets of cells such that:

//  1. All cells in a set have the same color.
//  2. Every two cells in a set share row or column.

// Input
// The first line of input contains integers n and m (1 ≤ n, m ≤ 50) — the number of rows and the number of columns correspondingly.
// The next n lines of input contain descriptions of rows. There are m integers, separated by spaces, in each line. The number equals 0 if the corresponding cell is colored white and equals 1 if the corresponding cell is colored black.

// Output
// Output single integer  — the number of non-empty sets from the problem description.

// Examples

// input
// 1 1
// 0
// output
// 1

// input
// 2 3
// 1 0 1
// 0 1 0
// output
// 8

// Note

// In the second example, there are six one-element sets. Additionally, there are two two-element sets, the first one consists of the first and the third cells of the first row, the second one consists of the first and the third cells of the second row. To sum up, there are 8 sets.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, n, m, i, j, cell, ans int
	var row, col []int

	for t, _ = fmt.Fscan(reader, &n, &m); t == 2; t, _ = fmt.Fscan(reader, &n, &m) {
		row = make([]int, n)
		col = make([]int, m)

		for i = 0; i < n; i++ {
			for j = 0; j < m; j++ {
				fmt.Fscan(reader, &cell)

				row[i] += cell
				col[j] += cell
			}
		}

		ans = 0

		for i = 0; i < n; i++ {
			ans += (1 << row[i]) - 1
			ans += (1 << (m - row[i])) - 1
		}

		for j = 0; j < m; j++ {
			ans += (1 << col[j]) - 1
			ans += (1 << (n - col[j])) - 1
		}

		ans -= n * m

		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
