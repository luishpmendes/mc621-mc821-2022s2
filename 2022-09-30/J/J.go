// https://codeforces.com/problemset/problem/1249/B1

// 1249B1 - Books Exchange (easy version)

// The only difference between easy and hard versions is constraints.

// There are n kids, each of them is reading a unique book. At the end of any day, the i-th kid will give his book to the pi-th kid (in case of i=pi the kid will give his book to himself). It is guaranteed that all values of pi are distinct integers from 1 to n (i.e. p is a permutation). The sequence p doesn't change from day to day, it is fixed.
// For example, if n=6 and p=[4,6,1,3,5,2] then at the end of the first day the book of the 1-st kid will belong to the 4-th kid, the 2-nd kid will belong to the 6-th kid and so on. At the end of the second day the book of the 1-st kid will belong to the 3-th kid, the 2-nd kid will belong to the 2-th kid and so on.
// Your task is to determine the number of the day the book of the i-th child is returned back to him for the first time for every i from 1 to n.
// Consider the following example: p=[5,1,2,4,3]. The book of the 1-st kid will be passed to the following kids:
//  • after the 1-st day it will belong to the 5-th kid,
//  • after the 2-nd day it will belong to the 3-rd kid,
//  • after the 3-rd day it will belong to the 2-nd kid,
//  • after the 4-th day it will belong to the 1-st kid,
// So after the fourth day, the book of the first kid will return to its owner. The book of the fourth kid will return to him for the first time after exactly one day.
// You have to answer q independent queries.

// Input
// The first line of the input contains one integer q (1≤q≤200) — the number of queries. Then q queries follow.
// The first line of the query contains one integer n (1≤n≤200) — the number of kids in the query. The second line of the query contains n integers p1,p2,…,pn (1≤pi≤n, all pi are distinct, i.e. p is a permutation), where pi is the kid which will get the book of the i-th kid.

// Output
// The first line of the query contains one integer n (1≤n≤200) — the number of kids in the query. The second line of the query contains n integers p1,p2,…,pn (1≤pi≤n, all pi are distinct, i.e. p is a permutation), where pi is the kid which will get the book of the i-th kid.

// Example
// input
// 6
// 5
// 1 2 3 4 5
// 3
// 2 3 1
// 6
// 4 6 2 1 5 3
// 1
// 1
// 4
// 3 4 1 2
// 5
// 5 1 2 4 3
// output
// 1 1 1 1 1
// 3 3 3
// 2 3 3 2 1 3
// 1
// 2 2 2 2
// 4 4 4 1 4

// Tutorial

// In this problem you just need to implement what is written in the problem statement. For the kid i the following pseudocode will calculate the answer (indices of the array p and its values are 0-indexed):

// pos = p[i]
// ans = 1
// while pos != i:
//     ans += 1
//     pos = p[pos]

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var q, n, i, pos, ans uint
	var p []uint

	for t, _ = fmt.Fscan(reader, &q); t == 1; t, _ = fmt.Fscan(reader, &q) {
		for ; q > 0; q-- {
			fmt.Fscan(reader, &n)
			p = make([]uint, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &p[i])
				p[i]--
			}

			for i = 0; i < n; i++ {
				pos = p[i]
				ans = 1

				for pos != i {
					ans++
					pos = p[pos]
				}

				fmt.Fprint(writer, ans)

				if i < n-1 {
					fmt.Fprint(writer, " ")
				} else {
					fmt.Fprintln(writer)
				}
			}
		}
	}

	writer.Flush()
}
