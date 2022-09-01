// https://codeforces.com/problemset/problem/629/A

// 629A - Far Relative’s Birthday Cake

// Door's family is going celebrate Famil Doors's birthday party. They love Famil Door so they are planning to make his birthday cake weird!
// The cake is a n × n square consisting of equal squares with side length 1. Each square is either empty or consists of a single chocolate. They bought the cake and randomly started to put the chocolates on the cake. The value of Famil Door's happiness will be equal to the number of pairs of cells with chocolates that are in the same row or in the same column of the cake. Famil Doors's family is wondering what is the amount of happiness of Famil going to be?
// Please, note that any pair can be counted no more than once, as two different cells can't share both the same row and the same column.

// Input
// In the first line of the input, you are given a single integer n (1 ≤ n ≤ 100) — the length of the side of the cake.
// Then follow n lines, each containing n characters. Empty cells are denoted with '.', while cells that contain chocolates are denoted by 'C'.

// Output
// Print the value of Famil Door's happiness, i.e. the number of pairs of chocolate pieces that share the same row or the same column.

// Examples

// input
// 3
// .CC
// C..
// C.C
// output
// 4

// input
// 4
// CC..
// C..C
// .CC.
// .CC.
// output
// 9

// Note

// If we number rows from top to bottom and columns from left to right, then, pieces that share the same row in the first sample are:
//  1. (1, 2) and (1, 3)
//  2. (3, 1) and (3, 3)
// Pieces that share the same column are:
//  1. (2, 1) and (3, 1)
//  2. (1, 3) and (3, 3)

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, n, i, j, col_count, row_count, ans int
	var cake [][]byte
	var s string

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		cake = make([][]byte, n)

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &s)
			cake[i] = []byte(s)
		}

		ans = 0

		for i = 0; i < n; i++ {
			col_count = 0
			row_count = 0

			for j = 0; j < n; j++ {
				if cake[i][j] == 'C' {
					row_count++
				}

				if cake[j][i] == 'C' {
					col_count++
				}
			}

			ans += (col_count * (col_count - 1)) / 2
			ans += (row_count * (row_count - 1)) / 2
		}

		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
