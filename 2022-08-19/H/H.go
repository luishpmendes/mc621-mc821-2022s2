// https://codeforces.com/problemset/problem/611/C

// 611C - New Year and Domino

// They say "years are like dominoes, tumbling one after the other". But would a year fit into a grid? I don't think so.
// Limak is a little polar bear who loves to play. He has recently got a rectangular grid with h rows and w columns. Each cell is a square, either empty (denoted by '.') or forbidden (denoted by '#'). Rows are numbered 1 through h from top to bottom. Columns are numbered 1 through w from left to right.
// Also, Limak has a single domino. He wants to put it somewhere in a grid. A domino will occupy exactly two adjacent cells, located either in one row or in one column. Both adjacent cells must be empty and must be inside a grid.
// Limak needs more fun and thus he is going to consider some queries. In each query he chooses some rectangle and wonders, how many way are there to put a single domino inside of the chosen rectangle?

// Input
// The first line of the input contains two integers h and w (1 ≤ h, w ≤ 500) – the number of rows and the number of columns, respectively.
// The next h lines describe a grid. Each line contains a string of the length w. Each character is either '.' or '#' — denoting an empty or forbidden cell, respectively.
// The next line contains a single integer q (1 ≤ q ≤ 100 000) — the number of queries.
// Each of the next q lines contains four integers r1i, c1i, r2i, c2i (1 ≤ r1i ≤ r2i ≤ h, 1 ≤ c1i ≤ c2i ≤ w) — the i-th query. Numbers r1i and c1i denote the row and the column (respectively) of the upper left cell of the rectangle. Numbers r2i and c2i denote the row and the column (respectively) of the bottom right cell of the rectangle.

// Output
// Print q integers, i-th should be equal to the number of ways to put a single domino inside the i-th rectangle.

// Examples

// input
// 5 8
// ....#..#
// .#......
// ##.#....
// ##..#.##
// ........
// 4
// 1 1 2 3
// 4 1 4 1
// 1 2 4 5
// 2 5 5 8

// output
// 4
// 0
// 10
// 15

// input
// 7 39
// .......................................
// .###..###..#..###.....###..###..#..###.
// ...#..#.#..#..#.........#..#.#..#..#...
// .###..#.#..#..###.....###..#.#..#..###.
// .#....#.#..#....#.....#....#.#..#..#.#.
// .###..###..#..###.....###..###..#..###.
// .......................................
// 6
// 1 1 3 20
// 2 10 6 30
// 2 10 7 30
// 2 2 7 7
// 1 7 7 7
// 1 8 7 8

// output
// 53
// 89
// 120
// 23
// 0
// 2

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, h, w, i, j, q, r1, c1, r2, c2, ans int
	var grid []string
	var dp_hor, dp_ver [][]int

	for t, _ = fmt.Fscan(reader, &h, &w); t == 2; t, _ = fmt.Fscan(reader, &h, &w) {
		grid = make([]string, h)
		dp_hor = make([][]int, h+1)
		dp_ver = make([][]int, h+1)

		for i = 0; i <= h; i++ {
			dp_hor[i] = make([]int, w+1)
			dp_ver[i] = make([]int, w+1)
		}

		for i = 0; i < h; i++ {
			fmt.Fscan(reader, &grid[i])
		}

		for i = 0; i < h; i++ {
			for j = 0; j < w; j++ {
				dp_hor[i+1][j+1] = dp_hor[i+1][j] + dp_hor[i][j+1] - dp_hor[i][j]
				dp_ver[i+1][j+1] = dp_ver[i+1][j] + dp_ver[i][j+1] - dp_ver[i][j]

				if grid[i][j] == '.' {
					if j+1 < w && grid[i][j+1] == '.' {
						dp_hor[i+1][j+1]++
					}

					if i+1 < h && grid[i+1][j] == '.' {
						dp_ver[i+1][j+1]++
					}
				}
			}
		}

		for fmt.Fscan(reader, &q); q > 0; q-- {
			fmt.Fscan(reader, &r1, &c1, &r2, &c2)
			r1--
			c1--
			ans = 0
			ans += dp_hor[r2][c2-1] - dp_hor[r2][c1] - dp_hor[r1][c2-1] + dp_hor[r1][c1]
			ans += dp_ver[r2-1][c2] - dp_ver[r2-1][c1] - dp_ver[r1][c2] + dp_ver[r1][c1]
			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
