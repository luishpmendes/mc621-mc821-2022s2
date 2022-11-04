// https://codeforces.com/problemset/problem/762/C

// 762C - Two strings

// You are given two strings a and b. You have to remove the minimum possible number of consecutive (standing one after another) characters from string b in such a way that it becomes a subsequence of string a. It can happen that you will not need to remove any characters at all, or maybe you will have to remove all of the characters from b and make it empty.
// Subsequence of string s is any such string that can be obtained by erasing zero or more characters (not necessarily consecutive) from string s.

// Input
// The first line contains string a, and the second line — string b. Both of these strings are nonempty and consist of lowercase letters of English alphabet. The length of each string is no bigger than 10^5 characters.

// Output
// On the first line output a subsequence of string a, obtained from b by erasing the minimum number of consecutive characters.
// If the answer consists of zero characters, output «-» (a minus sign).

// Examples

// input
// hi
// bob
// output
// -

// input
// abca
// accepted
// output
// ac

// input
// abacaba
// abcdcba
// output
// abcba

// Note

// In the first example strings a and b don't share any symbols, so the longest string that you can get is empty.
// In the second example ac is a subsequence of a, and at the same time you can obtain it by erasing consecutive symbols cepted from string b.

// Tutorial

// Try thinking not about erasing a substring from B, but rather picking some number of characters (possibly zero) from the left, and some from the right.
// Two pointers
// For every prefix of B, count how big of a prefix of A you will require. Call these values p[i]. Put infinity in the cells where even whole A is not enough.
// Same for every suffix of B count the length of the required suffix of A. Call these values s[i].
// Now try increasing the length of prefix of B, while decreasing the length of the suffix until p[pref_len] + s[suf_len] is less or equal to the length of A.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, i, j, n, m, length, ans1, ans2 int
	var a, b string
	var l, r []int

	for t, _ = fmt.Fscan(reader, &a, &b); t == 2; t, _ = fmt.Fscan(reader, &a, &b) {
		n = len(a)
		m = len(b)
		l = make([]int, n+2)
		r = make([]int, n+2)
		length = m

		for i, j = 1, 1; i <= m && j <= n; i++ {
			for j <= n && a[j-1] != b[i-1] {
				l[j] = i - 1
				j++
			}

			if j <= n {
				l[j] = i
				j++
			}
		}

		for i, j = m, n; i > 0 && j > 0; i-- {
			for j > 0 && a[j-1] != b[i-1] {
				r[j] = i + 1
				j--
			}

			if j > 0 {
				r[j] = i
				j--
			}
		}

		l[0] = 0
		r[n+1] = m + 1

		for i = 0; i <= n; i++ {
			if r[i+1] < length+l[i]+1 {
				length = r[i+1] - l[i] - 1
				ans1 = l[i]
				ans2 = r[i+1]
			}
		}

		if length == m {
			fmt.Fprintln(writer, "-")
		} else if length <= 0 {
			fmt.Fprintln(writer, b)
		} else {
			fmt.Fprintln(writer, b[:ans1]+b[ans2-1:])
		}
	}

	writer.Flush()
}
