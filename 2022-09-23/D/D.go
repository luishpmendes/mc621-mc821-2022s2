// https://codeforces.com/problemset/problem/630/A

// 630A - Again Twenty Five!

// The HR manager was disappointed again. The last applicant failed the interview the same way as 24 previous ones. "Do I give such a hard task?" — the HR manager thought. "Just raise number 5 to the power of n and get last two digits of the number. Yes, of course, n can be rather big, and one cannot find the power using a calculator, but we need people who are able to think, not just follow the instructions."
// Could you pass the interview in the machine vision company in IT City?

// Input
// The only line of the input contains a single integer n (2 ≤ n ≤ 2·1018) — the power in which you need to raise number 5.

// Output
// Output the last two digits of 5n without spaces between them.

// Examples
// input
// 2
// output
// 25

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
	var n uint

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		fmt.Fprintln(writer, 25)
	}

	writer.Flush()
}
