// https://codeforces.com/problemset/problem/1703/B

// 1703B - ICPC Balloons

// In an ICPC contest, balloons are distributed as follows:
//  • Whenever a team solves a problem, that team gets a balloon.
//  • The first team to solve a problem gets an additional balloon.
// A contest has 26 problems, labelled A, B, C, ..., Z. You are given the order of solved problems in the contest, denoted as a string s, where the i-th character indicates that the problem si has been solved by some team. No team will solve the same problem twice.
// Determine the total number of balloons that the teams received. Note that some problems may be solved by none of the teams.

// Input
// The first line of the input contains an integer t (1≤t≤100) — the number of testcases.
// The first line of each test case contains an integer n (1≤n≤50) — the length of the string.
// The second line of each test case contains a string s of length n consisting of uppercase English letters, denoting the order of solved problems.

// Output
// For each test case, output a single integer — the total number of balloons that the teams received.

// Example
// input
// 6
// 3
// ABA
// 1
// A
// 3
// ORZ
// 5
// BAAAA
// 4
// BKPT
// 10
// CODEFORCES
// output
// 5
// 2
// 6
// 7
// 8
// 17

// Note

// In the first test case, 5 balloons are given out:
//  • Problem A is solved. That team receives 2 balloons: one because they solved the problem, an an additional one because they are the first team to solve problem A.
//  • Problem B is solved. That team receives 2 balloons: one because they solved the problem, an an additional one because they are the first team to solve problem B.
//  • Problem A is solved. That team receives only 1 balloon, because they solved the problem. Note that they don't get an additional balloon because they are not the first team to solve problem A.
// The total number of balloons given out is 2+2+1=5.
// In the second test case, there is only one problem solved. The team who solved it receives 2 balloons: one because they solved the problem, an an additional one because they are the first team to solve problem A.

// Tutorial

// Let's keep an array a of booleans, ai denoting whether or not some team has solved the ith problem already. Now we can iterate through the string from left to right and keep a running total tot. If ai is true (the i-th problem has already been solved), increase tot by 1; otherwise, increase tot by 2 and set ai to true.
// The time complexity is O(n).
// Bonus: the answer is always n+number of distinct characters in s. Can you see why?

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
	var t, n, tot, i uint
	var s string
	var a [26]bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &s)
			tot = 0

			for i = 0; i < 26; i++ {
				a[i] = false
			}

			for i = 0; i < n; i++ {
				if a[s[i]-'A'] {
					tot++
				} else {
					tot += 2
					a[s[i]-'A'] = true
				}
			}

			fmt.Fprintln(writer, tot)
		}
	}

	writer.Flush()
}
