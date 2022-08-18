// https://codeforces.com/problemset/problem/1648/A

// 1648A - Weird Sum

// Egor has a table of size n×m, with lines numbered from 1 to n and columns numbered from 1 to m. Each cell has a color that can be presented as an integer from 1 to 105.
// Let us denote the cell that lies in the intersection of the r-th row and the c-th column as (r,c). We define the manhattan distance between two cells (r1,c1) and (r2,c2) as the length of a shortest path between them where each consecutive cells in the path must have a common side. The path can go through cells of any color. For example, in the table 3×4 the manhattan distance between (1,2) and (3,3) is 3, one of the shortest paths is the following: (1,2)→(2,2)→(2,3)→(3,3).
// Egor decided to calculate the sum of manhattan distances between each pair of cells of the same color. Help him to calculate this sum.

// Input
// The first line contains two integers n and m (1≤n≤m, n⋅m≤100000) — number of rows and columns in the table.
// Each of next n lines describes a row of the table. The i-th line contains m integers ci1,ci2,…,cim (1≤cij≤100000) — colors of cells in the i-th row.

// Output
// Print one integer — the the sum of manhattan distances between each pair of cells of the same color.

// Examples

// input
// 2 3
// 1 2 3
// 3 2 1
// output
// 7

// input
// 3 4
// 1 1 2 2
// 2 1 1 2
// 2 2 1 1
// output
// 76

// input
// 4 4
// 1 1 2 3
// 2 1 1 2
// 3 1 2 1
// 1 1 2 1
// output
// 129

// Note
// In the first sample there are three pairs of cells of same color: in cells (1,1) and (2,3), in cells (1,2) and (2,2), in cells (1,3) and (2,1). The manhattan distances between them are 3, 1 and 3, the sum is 7.

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
	var n, m, i, j, num_colors, ans int64
	var color, row, col [][]int64

	for t, _ = fmt.Fscan(reader, &n, &m); t == 2; t, _ = fmt.Fscan(reader, &n, &m) {
		color = make([][]int64, n)
		num_colors = 0
		ans = 0

		for i = 0; i < n; i++ {
			color[i] = make([]int64, m)

			for j = 0; j < m; j++ {
				fmt.Fscan(reader, &color[i][j])

				if num_colors < color[i][j] {
					num_colors = color[i][j]
				}
			}
		}

		row = make([][]int64, num_colors)
		col = make([][]int64, num_colors)

		for i = 0; i < n; i++ {
			for j = 0; j < m; j++ {
				row[color[i][j]-1] = append(row[color[i][j]-1], i)
				col[color[i][j]-1] = append(col[color[i][j]-1], j)
			}
		}

		for i = 0; i < num_colors; i++ {
			sort.Slice(row[i], func(a, b int) bool {
				return row[i][a] < row[i][b]
			})

			for j = 0; j < int64(len(row[i])); j++ {
				ans += (2*j + 1 - int64(len(row[i]))) * row[i][j]
			}

			sort.Slice(col[i], func(a, b int) bool {
				return col[i][a] < col[i][b]
			})

			for j = 0; j < int64(len(col[i])); j++ {
				ans += (2*j + 1 - int64(len(col[i]))) * col[i][j]
			}
		}

		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
