// https://codeforces.com/problemset/problem/1741/A

// 1741A - Compare T-Shirt Sizes

// Two T-shirt sizes are given: a and b. The T-shirt size is either a string M or a string consisting of several (possibly zero) characters X and one of the characters S or L.
// For example, strings M, XXL, S, XXXXXXXS could be the size of some T-shirts. And the strings XM, LL, SX are not sizes.
// The letter M stands for medium, S for small, L for large. The letter X refers to the degree of size (from eXtra). For example, XXL is extra-extra-large (bigger than XL, and smaller than XXXL).
// You need to compare two given sizes of T-shirts a and b.
// The T-shirts are compared as follows:
//  • any small size (no matter how many letters X) is smaller than the medium size and any large size;
//  • any large size (regardless of the number of letters X) is larger than the medium size and any small size;
//  • the more letters X before S, the smaller the size;
//  • the more letters X in front of L, the larger the size.
// For example:
//  • XXXS < XS
//  • XXXL > XL
//  • XL > M
//  • XXL = XXL
//  • XXXXXS < M
//  • XL > XXXS

// Input
// The first line of the input contains a single integer t (1≤t≤10^4) — the number of test cases.
// Each test case consists of one line, in which a and b T-shirt sizes are written. The lengths of the strings corresponding to the T-shirt sizes do not exceed 50. It is guaranteed that all sizes are correct.

// Output
// For each test case, print on a separate line the result of comparing a and b T-shirt sizes (lines "<", ">" or "=" without quotes).

// Example
// input
// 6
// XXXS XS
// XXXL XL
// XL M
// XXL XXL
// XXXXXS M
// L M
// output
// <
// >
// >
// =
// <
// >

// Tutorial

// Let sa, sb are the last characters of lines a and b respectively. And |a|,|b| are the sizes of these strings.
//  1. sa≠sb: then the answer depends only on sa and sb and is uniquely defined as the inverse of sa to sb ("<" if sa>sb, ">" if sa<sb, since the characters S, M, L are in reverse order in the alphabet).
//  2. sa=sb:
//   • |a|=|b|. Then the answer is "=". This also covers the case sa=sb=M;
//   • sa=sb=S. Then the larger the size of the string, the smaller the size of the t-shirt. That is, the answer is "<" if |a|>|b| and ">" if |a|<|b|;
//   • sa=sb=L. Then the larger the size of the string, the smaller the size of the t-shirt. That is, the answer is "<" if |a|<|b| and ">" if |a|>|b|;

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
	var t uint
	var a, b string

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &a, &b)

			if a[len(a)-1] > b[len(b)-1] {
				fmt.Fprintln(writer, "<")
			} else if a[len(a)-1] < b[len(b)-1] {
				fmt.Fprintln(writer, ">")
			} else { // a[len(a)-1] == b[len(b)-1]
				if len(a) == len(b) {
					fmt.Fprintln(writer, "=")
				} else if a[len(a)-1] == 'S' {
					if len(a) > len(b) {
						fmt.Fprintln(writer, "<")
					} else {
						fmt.Fprintln(writer, ">")
					}
				} else if a[len(a)-1] == 'L' {
					if len(a) < len(b) {
						fmt.Fprintln(writer, "<")
					} else {
						fmt.Fprintln(writer, ">")
					}
				}
			}
		}
	}

	writer.Flush()
}
