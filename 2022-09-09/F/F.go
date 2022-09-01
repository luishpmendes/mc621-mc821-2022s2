// https://codeforces.com/problemset/problem/1674/C

// 1674C - Infinite Replacement

// You are given a string s, consisting only of Latin letters 'a', and a string t, consisting of lowercase Latin letters.
// In one move, you can replace any letter 'a' in the string s with a string t. Note that after the replacement string s might contain letters other than 'a'.
// You can perform an arbitrary number of moves (including zero). How many different strings can you obtain? Print the number, or report that it is infinitely large.
// Two strings are considered different if they have different length, or they differ at some index.

// Input
// The first line contains a single integer q (1≤q≤104) — the number of testcases.
// The first line of each testcase contains a non-empty string s, consisting only of Latin letters 'a'. The length of s doesn't exceed 50.
// The second line contains a non-empty string t, consisting of lowercase Latin letters. The length of t doesn't exceed 50.

// Output
// For each testcase, print the number of different strings s that can be obtained after an arbitrary amount of moves (including zero). If the number is infinitely large, print -1. Otherwise, print the number.

// Example
// input
// 3
// aaaa
// a
// aa
// abc
// a
// b
// output
// 1
// -1
// 2

// Note

// In the first example, you can replace any letter 'a' with the string "a", but that won't change the string. So no matter how many moves you make, you can't obtain a string other than the initial one.
// In the second example, you can replace the second letter 'a' with "abc". String s becomes equal to "aabc". Then the second letter 'a' again. String s becomes equal to "aabcbc". And so on, generating infinitely many different strings.
// In the third example, you can either leave string s as is, performing zero moves, or replace the only 'a' with "b". String s becomes equal to "b", so you can't perform more moves on it.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, q int
	var s, t string
	var has_a bool

	for c, _ = fmt.Fscan(reader, &q); c == 1; c, _ = fmt.Fscan(reader, &q) {
		for ; q > 0; q-- {
			fmt.Fscan(reader, &s, &t)

			has_a = strings.Contains(t, "a")

			if len(t) == 1 && has_a {
				fmt.Fprintln(writer, 1)
			} else if has_a {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, int64(1)<<len(s))
			}
		}
	}

	writer.Flush()
}
