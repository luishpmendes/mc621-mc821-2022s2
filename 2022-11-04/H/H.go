// https://codeforces.com/problemset/problem/1654/B

// 1654B - Prefix Removals

// You are given a string s consisting of lowercase letters of the English alphabet. You must perform the following algorithm on s:
//  • Let x be the length of the longest prefix of s which occurs somewhere else in s as a contiguous substring (the other occurrence may also intersect the prefix). If x=0, break. Otherwise, remove the first x characters of s, and repeat.
// A prefix is a string consisting of several first letters of a given string, without any reorders. An empty prefix is also a valid prefix. For example, the string "abcd" has 5 prefixes: empty string, "a", "ab", "abc" and "abcd".
// For instance, if we perform the algorithm on s= "abcabdc",
//  • Initially, "ab" is the longest prefix that also appears somewhere else as a substring in s, so s= "cabdc" after 1 operation.
//  • Then, "c" is the longest prefix that also appears somewhere else as a substring in s, so s= "abdc" after 2 operations.
//  • Now x=0 (because there are no non-empty prefixes of "abdc" that also appear somewhere else as a substring in s), so the algorithm terminates.
// Find the final state of the string after performing the algorithm.

// Input
// The first line contains a single integer t (1≤t≤10^4) — the number of test cases.
// This is followed by t lines, each containing a description of one test case. Each line contains a string s. The given strings consist only of lowercase letters of the English alphabet and have lengths between 1 and 2⋅10^5 inclusive.
// It is guaranteed that the sum of the lengths of s over all test cases does not exceed 2⋅10^5.

// Output
// For each test case, print a single line containing the string s after executing the algorithm. It can be shown that such string is non-empty.

// Example
// input
// 6
// abcabdc
// a
// bbbbbbbbbb
// codeforces
// cffcfccffccfcffcfccfcffccffcfccf
// zyzyzwxxyyxxyyzzyzzxxwzxwywxwzxxyzzw
// output
// abdc
// a
// b
// deforces
// cf
// xyzzw

// Note

// The first test case is explained in the statement.
// In the second test case, no operations can be performed on s.
// In the third test case,
//  • Initially, s= "bbbbbbbbbb".
//  • After 1 operation, s= "b".
// In the fourth test case,
//  • Initially, s= "codeforces".
//  • After 1 operation, s= "odeforces".
//  • After 2 operations, s= "deforces".

// Hints

// Are there any characters that you can never remove?
// You can't remove the rightmost occurrence of each letter. Can you remove the other characters?
// If you can remove s1,s2,…,si−1 and si is not the rightmost occurrence of a letter, you can also remove si.

// Tutorial

// Let a be the initial string. For a string z, let's define z(l,r)=zlzl+1…zr, i.e., the substring of z from l to r. The final string is a(k,n) for some k.
// In the final string, x=0, so the first character doesn't appear anywhere else in s. This means that ak≠ak+1,ak+2,…,an. In other words, ak is the rightmost occurrence of a letter in s.
// Can you ever remove ai, if ai≠ai+1,ai+2,…,an? Notice that you would need to remove a(l,r) (l≤i≤r): this means that there must exist a(l′,r′)=a(l,r) for some l′>l. So, ai+l′−l=ai, and this is a contradiction.
// Therefore, k is the smallest index such that ak≠ak+1,ak+2,…,an.
// You can find k by iterating over the string from right to left and updating the frequency of each letter. Indeed ai≠ai+1,ai+2,…,an if and only if the frequency of the letter ai is 0 up to now (in the iteration from right to left we are performing). The value of k is the minimum such index i.
// Complexity: O(n)

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, i int
	var t uint
	var s string
	var frequency map[rune]int
	var x rune

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &s)
			frequency = make(map[rune]int)

			for _, x = range s {
				frequency[x]++
			}

			for i, x = range s {
				frequency[x]--

				if frequency[x] == 0 {
					break
				}
			}

			fmt.Fprintln(writer, s[i:])
		}
	}

	writer.Flush()
}
