// https://codeforces.com/contest/427/problem/B

// 427B - Prison Transfer
// The prison of your city has n prisoners. As the prison can't accommodate all of them, the city mayor has decided to transfer c of the prisoners to a prison located in another city.
// For this reason, he made the n prisoners to stand in a line, with a number written on their chests. The number is the severity of the crime he/she has committed. The greater the number, the more severe his/her crime was.
// Then, the mayor told you to choose the c prisoners, who will be transferred to the other prison. He also imposed two conditions. They are,
//  • The chosen c prisoners has to form a contiguous segment of prisoners.
//  • Any of the chosen prisoner's crime level should not be greater then t. Because, that will make the prisoner a severe criminal and the mayor doesn't want to take the risk of his running away during the transfer.
// Find the number of ways you can choose the c prisoners.

// Input
// The first line of input will contain three space separated integers n (1 ≤ n ≤ 2·10^5), t (0 ≤ t ≤ 10^9) and c (1 ≤ c ≤ n). The next line will contain n space separated integers, the ith integer is the severity ith prisoner's crime. The value of crime severities will be non-negative and will not exceed 10^9.

// Output
// Print a single integer — the number of ways you can choose the c prisoners.

// Examples

// input
// 4 3 3
// 2 3 1 1
// output
// 2

// input
// 1 1 1
// 2
// output
// 0

// input
// 11 4 2
// 2 2 0 7 3 2 2 4 9 1 4
// output
// 6

// Tutorial

// The severity of crimes form an integer sequence. Find all the contiguous sequences without any integer greater than t. If the length of any sequence is L, then we can choose c prisoners from them in L - c + 1 ways.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var q, n, t, c, i, severity, ans int
	var v []int

	for q, _ = fmt.Fscan(reader, &n, &t, &c); q == 3; q, _ = fmt.Fscan(reader, &n, &t, &c) {
		v = make([]int, 0)
		v = append(v, -1)

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &severity)

			if severity > t {
				v = append(v, i)
			}
		}

		v = append(v, n)
		ans = 0

		for i = 1; i < len(v); i++ {
			if v[i] > v[i-1]+c {
				ans += v[i] - v[i-1] - c
			}
		}

		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
