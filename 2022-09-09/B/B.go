// https://codeforces.com/problemset/problem/630/F

// 630F - Selection of Personnel

// One company of IT City decided to create a group of innovative developments consisting from 5 to 7 people and hire new employees for it. After placing an advertisment the company received n resumes. Now the HR department has to evaluate each possible group composition and select one of them. Your task is to count the number of variants of group composition to evaluate.

// Input
// The only line of the input contains one integer n (7 ≤ n ≤ 777) — the number of potential employees that sent resumes.

// Output
// Output one integer — the number of different variants of group composition.

// Example
// input
// 7
// output
// 29

package main

import (
	"bufio"
	"fmt"
	"os"
)

func c(n, k int) int64 {
	var res int64
	var i int

	res = 1

	for i = 1; i <= k; i++ {
		res = res * int64(n-k+i) / int64(i)
	}

	return res
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, n int

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		fmt.Fprintln(writer, c(n, 5)+c(n, 6)+c(n, 7))
	}

	writer.Flush()
}
