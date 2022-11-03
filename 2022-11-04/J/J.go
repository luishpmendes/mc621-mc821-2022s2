// https://codeforces.com/problemset/problem/448/B

// 448B - Suffix Structures

// Bizon the Champion isn't just a bison. He also is a favorite of the "Bizons" team.
// At a competition the "Bizons" got the following problem: "You are given two distinct words (strings of English letters), s and t. You need to transform word s into word t". The task looked simple to the guys because they know the suffix data structures well. Bizon Senior loves suffix automaton. By applying it once to a string, he can remove from this string any single character. Bizon Middle knows suffix array well. By applying it once to a string, he can swap any two characters of this string. The guys do not know anything about the suffix tree, but it can help them do much more.
// Bizon the Champion wonders whether the "Bizons" can solve the problem. Perhaps, the solution do not require both data structures. Find out whether the guys can solve the problem and if they can, how do they do it? Can they solve it either only with use of suffix automaton or only with use of suffix array or they need both structures? Note that any structure may be used an unlimited number of times, the structures may be used in any order.

// Input
// The first line contains a non-empty word s. The second line contains a non-empty word t. Words s and t are different. Each word consists only of lowercase English letters. Each word contains at most 100 letters.

// Output
// In the single line print the answer to the problem. Print "need tree" (without the quotes) if word s cannot be transformed into word t even with use of both suffix array and suffix automaton. Print "automaton" (without the quotes) if you need only the suffix automaton to solve the problem. Print "array" (without the quotes) if you need only the suffix array to solve the problem. Print "both" (without the quotes), if you need both data structures to solve the problem.
// It's guaranteed that if you can solve the problem only with use of suffix array, then it is impossible to solve it only with use of suffix automaton. This is also true for suffix automaton.

// Examples

// input
// automaton
// tomat
// output
// automaton

// input
// array
// arary
// output
// array

// input
// both
// hot
// output
// both

// input
// need
// tree
// output
// need tree

// Note

// In the third sample you can act like that: first transform "both" into "oth" by removing the first character using the suffix automaton and then make two swaps of the string using the suffix array and get "hot".

// Tutorial

// Consider each case separately. If we use only suffix automaton then s transform to some of its subsequence. Checking that t is a subsequence of s can be performed in different ways. Easiest and fastest — well-known two pointers method. In case of using suffix array we can get every permutation of s. If it is not obvious for you, try to think. Thus, s and t must be anagrams. If we count number of each letter in each string, we can check this. If every letter appears in s the same times as in t then words are anagrams. In case of using both structures strategy is: remove some letters and shuffle the rest. It is possible if every letter appears in s not less times than in t. Otherwise it is impossible to make t from s. Total complexity O(|s| + |t| + 26).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, i, j int
	var s, t string
	var cnt [26]int
	var automaton, array, both bool

	for c, _ = fmt.Fscan(reader, &s, &t); c == 2; c, _ = fmt.Fscan(reader, &s, &t) {
		for i = 0; i < 26; i++ {
			cnt[i] = 0
		}

		automaton = false

		for i, j = 0, 0; i < len(s); i++ {
			if j < len(t) && s[i] == t[j] {
				j++
			}

			if j == len(t) {
				automaton = true
			}
		}

		for i = 0; i < len(s); i++ {
			cnt[s[i]-'a']++
		}

		for i = 0; i < len(t); i++ {
			cnt[t[i]-'a']--
		}

		array = true
		both = true

		for i = 0; i < 26; i++ {
			array = array && (cnt[i] == 0)
			both = both && (cnt[i] >= 0)
		}

		if automaton {
			fmt.Fprintln(writer, "automaton")
		} else if array {
			fmt.Fprintln(writer, "array")
		} else if both {
			fmt.Fprintln(writer, "both")
		} else {
			fmt.Fprintln(writer, "need tree")
		}
	}

	writer.Flush()
}
