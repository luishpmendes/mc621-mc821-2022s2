// https://codeforces.com/problemset/problem/1335/D

// 1335D - Anti-Sudoku

// You are given a correct solution of the sudoku puzzle. If you don't know what is the sudoku, you can read about it here.
// Your task is to change at most 9 elements of this field (i.e. choose some 1≤i,j≤9 and change the number at the position (i,j) to any other number in range [1;9]) to make it anti-sudoku. The anti-sudoku is the 9×9 field, in which:
// • Any number in this field is in range [1;9];
// • each row contains at least two equal elements;
// • each column contains at least two equal elements;
// • each 3×3 block (you can read what is the block in the link above) contains at least two equal elements.
// It is guaranteed that the answer exists.
// You have to answer t independent test cases.

// Input
// The first line of the input contains one integer t (1≤t≤104) — the number of test cases. Then t test cases follow.
// Each test case consists of 9 lines, each line consists of 9 characters from 1 to 9 without any whitespaces — the correct solution of the sudoku puzzle.

// Output
// For each test case, print the answer — the initial field with at most 9 changed elements so that the obtained field is anti-sudoku. If there are several solutions, you can print any. It is guaranteed that the answer exists.

// Example
// input
// 1
// 154873296
// 386592714
// 729641835
// 863725149
// 975314628
// 412968357
// 631457982
// 598236471
// 247189563
// output
// 154873396
// 336592714
// 729645835
// 863725145
// 979314628
// 412958357
// 631457992
// 998236471
// 247789563

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, t, i, j int
	var s string

	// Reads t until EOF
	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			for i = 0; i < 9; i++ {
				fmt.Fscan(reader, &s)

				for j = 0; j < 9; j++ {
					if s[j] == '2' {
						fmt.Fprint(writer, "1")
					} else {
						fmt.Fprintf(writer, "%c", s[j])
					}
				}

				fmt.Fprintln(writer)
			}
		}
	}

	writer.Flush()
}
