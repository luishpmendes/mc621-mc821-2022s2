// https://codeforces.com/problemset/problem/851/B

// 851B - Arpa and an exam about geometry
// Arpa is taking a geometry exam. Here is the last problem of the exam.
// You are given three points a, b, c.
// Find a point and an angle such that if we rotate the page around the point by the angle, the new position of a is the same as the old position of b, and the new position of b is the same as the old position of c.
// Arpa is doubting if the problem has a solution or not (i.e. if there exists a point and an angle satisfying the condition). Help Arpa determine if the question has a solution or not.

// Input
// The only line contains six integers ax, ay, bx, by, cx, cy (|ax|, |ay|, |bx|, |by|, |cx|, |cy| ≤ 10^9). It's guaranteed that the points are distinct.

// Output
// Print "Yes" if the problem has a solution, "No" otherwise.
// You can print each letter in any case (upper or lower).

// Examples

// input
// 0 1 1 1 1 0
// output
// Yes

// input
// 1 1 0 0 1000 1000
// output
// No

// Note

// In the first sample test, rotate the page around (0.5, 0.5) by .
// In the second sample test, you can't find any solution.

// Tutorial

// If a, b, c are on the same line or dis(a, b) ≠ dis(b, c), the problem has not any solution.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var ax, ay, bx, by, cx, cy int64

	for t, _ = fmt.Fscan(reader, &ax, &ay, &bx, &by, &cx, &cy); t == 6; t, _ = fmt.Fscan(reader, &ax, &ay, &bx, &by, &cx, &cy) {
		if (ax-bx)*(ax-bx)+(ay-by)*(ay-by) != (bx-cx)*(bx-cx)+(by-cy)*(by-cy) || (by-ay)*(cx-ax) == (cy-ay)*(bx-ax) {
			fmt.Fprintln(writer, "No")
		} else {
			fmt.Fprintln(writer, "Yes")

		}
	}

	writer.Flush()
}
