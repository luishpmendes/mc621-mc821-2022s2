// https://codeforces.com/problemset/problem/1682/A

// 1682A - Palindromic Indices

// You are given a palindromic string s of length n.
// You have to count the number of indices i (1≤i≤n) such that the string after removing si from s still remains a palindrome.
// For example, consider s = "aba"
//  1. If we remove s1 from s, the string becomes "ba" which is not a palindrome.
//  2. If we remove s2 from s, the string becomes "aa" which is a palindrome.
//  3. If we remove s3 from s, the string becomes "ab" which is not a palindrome.
// A palindrome is a string that reads the same backward as forward. For example, "abba", "a", "fef" are palindromes whereas "codeforces", "acd", "xy" are not.

// Input
// The input consists of multiple test cases. The first line of the input contains a single integer t (1≤t≤10^3)  — the number of test cases. Description of the test cases follows.
// The first line of each testcase contains a single integer n (2≤n≤10^5)  — the length of string s.
// The second line of each test case contains a string s consisting of lowercase English letters. It is guaranteed that s is a palindrome.
// It is guaranteed that sum of n over all test cases does not exceed 2⋅10^5.

// Output
// For each test case, output a single integer  — the number of indices i (1≤i≤n) such that the string after removing si from s still remains a palindrome.

// Example
// input
// 3
// 3
// aba
// 8
// acaaaaca
// 2
// dd
// output
// 1
// 4
// 2

// Note

// The first test case is described in the statement.
// In the second test case, the indices i that result in palindrome after removing si are 3,4,5,6. Hence the answer is 4.
// In the third test case, removal of any of the indices results in "d" which is a palindrome. Hence the answer is 2.

// Hint

// Read the statement carefully!! The given string is a palindrome.

// Tutorial

// Let's remove some index i from the first half of s and check whether the resulting string is a palindrome or not, the other half has the same approach. The prefix of length i−1 already matches with the suffix of the same length because the initial string was a palindrome, so we just need to check if t=s[i+1…n−i+1] is a palindrome.
// For t to be a palindrome, sn−i+1 should be equal to si+1 which was initially equal to sn−i, again which should be equal to si+2 and this goes on. Here we can see that si=si+1…=sn−i+1. So the answer is simply equal to the number of contiguous same characters in the center of the string which can be calculated in O(n).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, n, i int
	var t, cnt uint
	var s string

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &s)

			cnt = 0

			for i = (n - 1) >> 1; i >= 0; i-- {
				if s[i] == s[(n-1)>>1] {
					cnt++
				} else {
					break
				}
			}

			cnt <<= 1

			if n&1 == 1 {
				cnt--
			}

			fmt.Fprintln(writer, cnt)
		}
	}

	writer.Flush()
}
