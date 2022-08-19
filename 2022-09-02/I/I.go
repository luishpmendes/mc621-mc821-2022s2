// https://codeforces.com/problemset/problem/1716/A

// 1716A - 2-3 Moves

// You are standing at the point 0 on a coordinate line. Your goal is to reach the point n. In one minute, you can move by 2 or by 3 to the left or to the right (i. e., if your current coordinate is x, it can become x−3, x−2, x+2 or x+3). Note that the new coordinate can become negative.
// Your task is to find the minimum number of minutes required to get from the point 0 to the point n.
// You have to answer t independent test cases.

// Input
// The first line of the input contains one integer t (1≤t≤104) — the number of test cases. Then t lines describing the test cases follow.
// The i-th of these lines contains one integer n (1≤n≤109) — the goal of the i-th test case.

// Output
// For each test case, print one integer — the minimum number of minutes required to get from the point 0 to the point n for the corresponding test case.

// Example
// input
// 4
// 1
// 3
// 4
// 12
// output
// 2
// 1
// 2
// 4

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, t, n int

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)

			if n == 1 {
				fmt.Fprintln(writer, 2)
			} else if n <= 3 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, (n+2)/3)
			}
		}
	}

	writer.Flush()
}
