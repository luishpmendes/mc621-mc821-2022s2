// https://codeforces.com/problemset/problem/822/B

// 822B - Crossword solving

// Erelong Leha was bored by calculating of the greatest common divisor of two factorials. Therefore he decided to solve some crosswords. It's well known that it is a very interesting occupation though it can be very difficult from time to time. In the course of solving one of the crosswords, Leha had to solve a simple task. You are able to do it too, aren't you?
// Leha has two strings s and t. The hacker wants to change the string s at such way, that it can be found in t as a substring. All the changes should be the following: Leha chooses one position in the string s and replaces the symbol in this position with the question mark "?". The hacker is sure that the question mark in comparison can play the role of an arbitrary symbol. For example, if he gets string s="ab?b" as a result, it will appear in t="aabrbb" as a substring.
// Guaranteed that the length of the string s doesn't exceed the length of the string t. Help the hacker to replace in s as few symbols as possible so that the result of the replacements can be found in t as a substring. The symbol "?" should be considered equal to any other symbol

// Input
// The first line contains two integers n and m (1 ≤ n ≤ m ≤ 1000) — the length of the string s and the length of the string t correspondingly.
// The second line contains n lowercase English letters — string s.
// The third line contains m lowercase English letters — string t.

// Output
// In the first line print single integer k — the minimal number of symbols that need to be replaced.
// In the second line print k distinct integers denoting the positions of symbols in the string s which need to be replaced. Print the positions in any order. If there are several solutions print any of them. The numbering of the positions begins from one.

// Examples

// input
// 3 5
// abc
// xaybz
// output
// 2
// 2 3

// input
// 4 10
// abcd
// ebceabazcd
// output
// 1
// 2

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, n, m, i, j int
	var s, t string
	var ans, new_ans []int

	for c, _ = fmt.Fscan(reader, &n, &m); c == 2; c, _ = fmt.Fscan(reader, &n, &m) {
		fmt.Fscan(reader, &s, &t)
		ans = make([]int, 0, n)

		for i = 0; i < n; i++ {
			ans = append(ans, i)
		}

		for i = 0; i < m-n+1; i++ {
			new_ans = make([]int, 0, n)

			for j = 0; j < n; j++ {
				if s[j] != t[i+j] {
					new_ans = append(new_ans, j)
				}
			}

			if len(ans) > len(new_ans) {
				ans = new_ans
			}
		}

		fmt.Fprintln(writer, len(ans))

		if len(ans) > 0 {
			for i = 0; i < len(ans)-1; i++ {
				fmt.Fprint(writer, ans[i]+1, " ")
			}

			fmt.Fprintln(writer, ans[len(ans)-1]+1)
		}
	}

	writer.Flush()
}
