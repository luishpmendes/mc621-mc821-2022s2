// https://codeforces.com/problemset/problem/1146/B

// 1146B - Hate "A"

// Bob has a string s consisting of lowercase English letters. He defines s′ to be the string after removing all "a" characters from s (keeping all other characters in the same order). He then generates a new string t by concatenating s and s′. In other words, t=s+s′ (look at notes for an example).
// You are given a string t. Your task is to find some s that Bob could have used to generate t. It can be shown that if an answer exists, it will be unique.

// Input
// The first line of input contains a string t (1≤|t|≤10^5) consisting of lowercase English letters.

// Output
// Print a string s that could have generated t. It can be shown if an answer exists, it is unique. If no string exists, print ":(" (without double quotes, there is no space between the characters).

// Examples

// input
// aaaaa
// output
// aaaaa

// input
// aacaababc
// output
// :(

// input
// ababacacbbcc
// output
// ababacac

// input
// baba
// output
// :(

// Note

// In the first example, we have s= "aaaaa", and s′= "".
// In the second example, no such s can work that will generate the given t.
// In the third example, we have s= "ababacac", and s′= "bbcc", and t=s+s′= "ababacacbbcc".

// Tutorial

// There are a few different ways to approach this.
// In one way, we can approach it by looking at the frequency of all characters. We want to find a split point so that all "a"s lie on the left side of the split, and all other characters are split evenly between the left and right side. This split can be uniquely determined and found in linear time by keeping track of prefix sums (it also might be possible there is no split, in which case the answer is impossible). After finding a split, we still need to make sure the characters appear in the same order, which is another linear time pass.
// In another way, let's count the number of "a" and non-"a" characters. Let these numbers be c0 and c1. If c1 is not divisible by 2, then the answer is impossible. Otherwise, we know the suffix of length c1/2 of t must be s′, and we can check there are no "a"s there. We also check that after removing all "a"s from t, the first and second half of the strings are equal.

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
	var c, i int
	var t string
	var s strings.Builder
	var res string
	var flag bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		s = strings.Builder{}

		for _, c := range t {
			if c != 'a' {
				s.WriteRune(c)
			}
		}

		res = s.String()
		flag = true

		if len(res)&1 == 0 {
			for i = 0; i < len(res)>>1 && flag; i++ {
				if res[i] != res[i+len(res)>>1] {
					flag = false
				} else if res[i] != t[len(t)+i-len(res)>>1] {
					flag = false
				}
			}
		} else {
			flag = false
		}

		if flag {
			for i = 0; i < len(t)-len(res)>>1; i++ {
				fmt.Fprint(writer, string(t[i]))
			}

			fmt.Fprintln(writer)
		} else {
			fmt.Fprintln(writer, ":(")
		}
	}

	writer.Flush()
}
