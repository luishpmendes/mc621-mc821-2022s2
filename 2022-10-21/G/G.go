// https://codeforces.com/problemset/problem/1660/C

// 1660C - Get an Even String

// A string a=a1a2…an is called even if it consists of a concatenation (joining) of strings of length 2 consisting of the same characters. In other words, a string a is even if two conditions are satisfied at the same time:
//  • its length n is even;
//  • for all odd i (1≤i≤n−1), ai=ai+1 is satisfied.
// For example, the following strings are even: "" (empty string), "tt", "aabb", "oooo", and "ttrrrroouuuuuuuukk". The following strings are not even: "aaa", "abab" and "abba".
// Given a string s consisting of lowercase Latin letters. Find the minimum number of characters to remove from the string s to make it even. The deleted characters do not have to be consecutive.

// Input
// The first line of input data contains an integer t (1≤t≤10^4) —the number of test cases in the test.
// The descriptions of the test cases follow.
// Each test case consists of one string s (1≤|s|≤2⋅10^5), where |s| — the length of the string s. The string consists of lowercase Latin letters.
// It is guaranteed that the sum of |s| on all test cases does not exceed 2⋅10^5.

// Output
// For each test case, print a single number — the minimum number of characters that must be removed to make s even.

// Examples
// input
// 6
// aabbdabdccc
// zyx
// aaababbb
// aabbcc
// oaoaaaoo
// bmefbmuyw
// output
// 3
// 3
// 2
// 0
// 2
// 7

// Note

// In the first test case you can remove the characters with indices 6, 7, and 9 to get an even string "aabbddcc".
// In the second test case, each character occurs exactly once, so in order to get an even string, you must remove all characters from the string.
// In the third test case, you can get an even string "aaaabb" by removing, for example, 4-th and 6-th characters, or a string "aabbbb" by removing the 5-th character and any of the first three.

// Tutorial

// We will act greedily: we will make an array prev consisting of 26 elements, in which we will mark prev[i]=true if the letter is already encountered in the string, and prev[i]=false otherwise. In the variable m we will store the length of the even string that can be obtained from s. We will go through the string by executing the following algorithm:
//  • if prev[i]=false, mark prev[i]=true.
//  • if prev[i]=true, then we already have a pair of repeating characters to add to an even string — add 2 to the number m and clear the array used.
// Clearing prev is necessary because both characters that will make up the next pair must be in the s string after the current character. In other words, if the last character in the current pair was st, then the first character in the new pair can be sk, where t<k<n.
// Then we calculate the answer as n−m.

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
	var t, m, n, i uint
	var s string
	var prev [26]bool
	var char rune

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &s)
			m = 0
			n = uint(len(s))

			for i = 0; i < 26; i++ {
				prev[i] = false
			}

			for _, char = range s {
				if prev[char-'a'] {
					m += 2
					for i = 0; i < 26; i++ {
						prev[i] = false
					}
				} else {
					prev[char-'a'] = true
				}
			}

			fmt.Fprintln(writer, n-m)
		}
	}

	writer.Flush()
}
