// https://codeforces.com/problemset/problem/617/A

// 617A - Elephant

// An elephant decided to visit his friend. It turned out that the elephant's house is located at point 0 and his friend's house is located at point x(x > 0) of the coordinate line. In one step the elephant can move 1, 2, 3, 4 or 5 positions forward. Determine, what is the minimum number of steps he need to make in order to get to his friend's house.

// Input
// The first line of the input contains an integer x (1 ≤ x ≤ 1 000 000) — The coordinate of the friend's house.

// Output
// Print the minimum number of steps that elephant needs to make to get from point 0 to point x.

// Examples

// input
// 5
// output
// 1

// input
// 12
// output
// 3

// Note
// In the first sample the elephant needs to make one step of length 5 to reach the point x.
// In the second sample the elephant can get to point x if he moves by 3, 5 and 4. There are other ways to get the optimal answer but the elephant cannot reach x in less than three moves.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, x int

	for t, _ = fmt.Fscan(reader, &x); t == 1; t, _ = fmt.Fscan(reader, &x) {
		fmt.Fprintln(writer, (x+4)/5)
	}

	writer.Flush()
}
