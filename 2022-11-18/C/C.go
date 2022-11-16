// https://codeforces.com/problemset/problem/6/A

// 6A - Triangle

// Johnny has a younger sister Anne, who is very clever and smart. As she came home from the kindergarten, she told his brother about the task that her kindergartener asked her to solve. The task was just to construct a triangle out of four sticks of different colours. Naturally, one of the sticks is extra. It is not allowed to break the sticks or use their partial length. Anne has perfectly solved this task, now she is asking Johnny to do the same.
// The boy answered that he would cope with it without any difficulty. However, after a while he found out that different tricky things can occur. It can happen that it is impossible to construct a triangle of a positive area, but it is possible to construct a degenerate triangle. It can be so, that it is impossible to construct a degenerate triangle even. As Johnny is very lazy, he does not want to consider such a big amount of cases, he asks you to help him.

// Input
// The first line of the input contains four space-separated positive integer numbers not exceeding 100 — lengthes of the sticks.

// Output
// Output TRIANGLE if it is possible to construct a non-degenerate triangle. Output SEGMENT if the first case cannot take place and it is possible to construct a degenerate triangle. Output IMPOSSIBLE if it is impossible to construct any triangle. Remember that you are to use three sticks. It is not allowed to break the sticks or use their partial length.

// Examples

// input
// 4 2 1 3
// output
// TRIANGLE

// input
// 7 2 2 4
// output
// SEGMENT

// input
// 3 5 9 1
// output
// IMPOSSIBLE

// Tutorial

// For each of the possible combinations of three sticks , we can make a triangle if sum of the lengths of the smaller two is greater than the length of the third and we can make a segment in case of equality.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func triangle(a, b, c uint) bool {
	return a+b > c && a+c > b && b+c > a
}

func segment(a, b, c uint) bool {
	return a+b == c || a+c == b || b+c == a
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var a, b, c, d uint

	for t, _ = fmt.Fscan(reader, &a, &b, &c, &d); t == 4; t, _ = fmt.Fscan(reader, &a, &b, &c, &d) {
		if triangle(a, b, c) || triangle(a, b, d) || triangle(a, c, d) || triangle(b, c, d) {
			fmt.Fprintln(writer, "TRIANGLE")
		} else if segment(a, b, c) || segment(a, b, d) || segment(a, c, d) || segment(b, c, d) {
			fmt.Fprintln(writer, "SEGMENT")
		} else {
			fmt.Fprintln(writer, "IMPOSSIBLE")
		}
	}

	writer.Flush()
}
