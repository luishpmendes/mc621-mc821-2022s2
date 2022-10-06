// https://codeforces.com/problemset/problem/1674/B

// 1674B - Dictionary

// The Berland language consists of words having exactly two letters. Moreover, the first letter of a word is different from the second letter. Any combination of two different Berland letters (which, by the way, are the same as the lowercase letters of Latin alphabet) is a correct word in Berland language.
// The Berland dictionary contains all words of this language. The words are listed in a way they are usually ordered in dictionaries. Formally, word a comes earlier than word b in the dictionary if one of the following conditions hold:
//  • the first letter of a is less than the first letter of b;
//  • the first letters of a and b are the same, and the second letter of a is less than the second letter of b.
// So, the dictionary looks like that:
//  • Word 1: ab
//  • Word 2: ac
//  • ...
//  • Word 25: az
//  • Word 26: ba
//  • Word 27: bc
//  • ...
//  • Word 649: zx
//  • Word 650: zy
// You are given a word s from the Berland language. Your task is to find its index in the dictionary.

// Input
// The first line contains one integer t (1≤t≤650) — the number of test cases.
// Each test case consists of one line containing s — a string consisting of exactly two different lowercase Latin letters (i. e. a correct word of the Berland language).

// Output
// For each test case, print one integer — the index of the word s in the dictionary.

// Example
// input
// 7
// ab
// ac
// az
// ba
// bc
// zx
// zy
// output
// 1
// 2
// 25
// 26
// 27
// 649
// 650

// Tutorial

// There are many different ways to solve this problem:
//  • generate all Berland words with two for-loops and store them in an array, then for each test case, go through the array of words to find the exact word you need;
//  • generate all Berland words with two for-loops and store them in a dictionary-like data structure (map in C++, dict in Python, etc), using words as keys and their numbers as values. This allows to search for the index of the given word quickly;
//  • for each test case, run two for-loops to iterate over the words, count the number of words we skipped, and stop at the word from the test case;
//  • try to invent some formulas that allow counting the number of words before the given one.

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
	var t, i uint
	var s string
	var idx map[string]uint
	var chr1, chr2 byte

	i = 1
	idx = make(map[string]uint)

	for chr1 = 'a'; chr1 <= 'z'; chr1++ {
		for chr2 = 'a'; chr2 <= 'z'; chr2++ {
			if chr1 != chr2 {
				s = string(chr1) + string(chr2)
				idx[s] = i
				i++
			}
		}
	}

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &s)
			fmt.Fprintln(writer, idx[s])
		}
	}

	writer.Flush()
}
