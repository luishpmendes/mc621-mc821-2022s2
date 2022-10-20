// https://codeforces.com/problemset/problem/1703/C

// 1703C - Cypher

// Luca has a cypher made up of a sequence of n wheels, each with a digit ai written on it. On the i-th wheel, he made bi moves. Each move is one of two types:
//  • up move (denoted by U): it increases the i-th digit by 1. After applying the up move on 9, it becomes 0.
//  • down move (denoted by D): it decreases the i-th digit by 1. After applying the down move on 0, it becomes 9.
// Luca knows the final sequence of wheels and the moves for each wheel. Help him find the original sequence and crack the cypher.

// Input
// The first line contains a single integer t (1≤t≤100) — the number of test cases.
// The first line of each test case contains a single integer n (1≤n≤100) — the number of wheels.
// The second line contains n integers ai (0≤ai≤9) — the digit shown on the i-th wheel after all moves have been performed.
// Then n lines follow, the i-th of which contains the integer bi (1≤bi≤10) and bi characters that are either U or D — the number of moves performed on the i-th wheel, and the moves performed. U and D represent an up move and a down move respectively.

// Output
// For each test case, output n space-separated digits  — the initial sequence of the cypher.

// Example
// input
// 3
// 3
// 9 3 1
// 3 DDD
// 4 UDUU
// 2 DU
// 2
// 0 9
// 9 DDDDDDDDD
// 9 UUUUUUUUU
// 5
// 0 5 9 8 3
// 10 UUUUUUUUUU
// 3 UUD
// 8 UUDUUDDD
// 10 UUDUUDUDDU
// 4 UUUU
// output
// 2 1 1
// 9 0
// 0 4 9 6 9

// Note

// In the first test case, we can prove that initial sequence was [2,1,1]. In that case, the following moves were performed:
//  • On the first wheel: 2 →D 1 →D 0 →D 9.
//  • On the second wheel: 1 →U 2 →D 1 →U 2 →U 3.
//  • On the third wheel: 1 →D 0 →U 1.
// The final sequence was [9,3,1], which matches the input.

// Tutorial

// We will perform each move in reverse from the final sequence of the cypher.
//  • down move: it increases the i-th digit by 1. After applying the up move on 9, it becomes 0.
//  • up move (denoted by D): it decreases the i-th digit by 1. After applying the down move on 0, it becomes 9.
// Now we just need to implement the two types of moves. The time complexity is O(n+∑ai) per test case.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, b, i, j uint
	var a []int
	var s string

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			a = make([]int, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])
			}

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &b)

				if b > 0 {
					fmt.Fscan(reader, &s)

					for j = 0; j < b; j++ {
						if s[j] == 'U' {
							a[i]--
						} else if s[j] == 'D' {
							a[i]++
						}

						if a[i] < 0 {
							a[i] += 10
						}
						if a[i] > 9 {
							a[i] -= 10
						}
					}
				}
			}

			for i = 0; i+1 < n; i++ {
				fmt.Fprint(writer, a[i], " ")
			}

			fmt.Fprintln(writer, a[n-1])
		}
	}

	writer.Flush()
}
