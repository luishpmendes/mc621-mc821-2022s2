// https://codeforces.com/problemset/problem/1706/A

// 1706A - Another String Minimization Problem

// You have a sequence a1,a2,…,an of length n, consisting of integers between 1 and m. You also have a string s, consisting of m characters B.
// You are going to perform the following n operations.
//  • At the i-th (1≤i≤n) operation, you replace either the ai-th or the (m+1−ai)-th character of s with A. You can replace the character at any position multiple times through the operations.
// Find the lexicographically smallest string you can get after these operations.
// A string x is lexicographically smaller than a string y of the same length if and only if in the first position where x and y differ, the string x has a letter that appears earlier in the alphabet than the corresponding letter in y.

// Input
// The first line contains the number of test cases t (1≤t≤2000).
// The first line of each test case contains two integers n and m (1≤n,m≤50) — the length of the sequence a and the length of the string s respectively.
// The second line contains n integers a1,a2,…,an (1≤ai≤m) — the sequence a.

// Output
// For each test case, print a string of length m — the lexicographically smallest string you can get. Each character of the string should be either capital English letter A or capital English letter B.

// Example
// input
// 6
// 4 5
// 1 1 3 1
// 1 5
// 2
// 4 1
// 1 1 1 1
// 2 4
// 1 3
// 2 7
// 7 5
// 4 5
// 5 5 3 5
// output
// ABABA
// BABBB
// A
// AABB
// ABABBBB
// ABABA

// Note

// In the first test case, the sequence a=[1,1,3,1]. One of the possible solutions is the following.
//  • At the 1-st operation, you can replace the 1-st character of s with A. After it, s becomes ABBBB.
//  • At the 2-nd operation, you can replace the 5-th character of s with A (since m+1−a2=5). After it, s becomes ABBBA.
//  • At the 3-rd operation, you can replace the 3-rd character of s with A. After it, s becomes ABABA.
//  • At the 4-th operation, you can replace the 1-st character of s with A. After it, s remains equal to ABABA.
// The resulting string is ABABA. It is impossible to produce a lexicographically smaller string.
// In the second test case, you are going to perform only one operation. You can replace either the 2-nd character or 4-th character of s with A. You can get strings BABBB and BBBAB after the operation. The string BABBB is the lexicographically smallest among these strings.
// In the third test case, the only string you can get is A.
// In the fourth test case, you can replace the 1-st and 2-nd characters of s with A to get AABB.
// In the fifth test case, you can replace the 1-st and 3-rd characters of s with A to get ABABBBB.

// Tutorial

// Let's iterate through the elements of a. For convenience, we'll make ai=min(ai,m+1−ai). If the ai-th character of s is not currently A, then we should replace it. Otherwise, we replace the (m+1−ai)-th character. This is because if we have the choice between replacing two characters, replacing the one with the smaller index will result in a lexicographically smaller string.
// Alternatively, we can keep track of how many times either x or m+1−x appears in a for each 1≤x≤⌈m/2⌉.
//  • If they appear 0 times, neither of these indices in s can become A.
//  • If they appear 1 time, it is optimal to set the x-th character to A, since this will produce a lexicographically smaller string.
//  • Otherwise, they appear at least 2 times, and it is possible to set both the x-th and (m+1−x)-th character to A.

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
	var t, n, m, a, i uint
	var s []byte
	var cnt []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &m)
			cnt = make([]uint, m)
			s = make([]byte, m)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a)
				a--

				if a > m-1-a {
					a = m - 1 - a
				}

				cnt[a]++
			}

			for i = 0; i < m; i++ {
				s[i] = 'B'
			}

			for i = 0; i < m; i++ {
				if cnt[i] > 0 {
					s[i] = 'A'

					if cnt[i] > 1 {
						s[m-i-1] = 'A'
					}
				}
			}

			fmt.Fprintln(writer, string(s))
		}
	}

	writer.Flush()
}
