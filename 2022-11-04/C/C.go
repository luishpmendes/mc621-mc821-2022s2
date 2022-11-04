// https://codeforces.com/problemset/problem/1285/D

// 1285D - Dr. Evil Underscores

// Today, as a friendship gift, Bakry gave Badawy n integers a1,a2,…,an and challenged him to choose an integer X such that the value max1≤i≤n(ai⊕X) is minimum possible, where ⊕ denotes the bitwise XOR operation.
// As always, Badawy is too lazy, so you decided to help him and find the minimum possible value of max1≤i≤n(ai⊕X).

// Input
// The first line contains integer n (1≤n≤10^5).
// The second line contains n integers a1,a2,…,an (0≤ai≤2^30−1).

// Output
// Print one integer — the minimum possible value of max1≤i≤n(ai⊕X).

// Examples

// input
// 3
// 1 2 3
// output
// 2

// input
// 2
// 1 5
// output
// 4

// Note

// In the first sample, we can choose X=3.
// In the second sample, we can choose X=5.

// Tutorial

// We will solve this problem recursively starting from the most significant bit. Let's split the elements into two groups, one with the elements which have the current bit on and one with the elements which have the current bit off. If either group is empty, we can assign the current bit of X accordingly so that we have the current bit off in our answer, so we will just proceed to the next bit. Otherwise, both groups aren't empty, so whatever value we assign to the current bit of X, we will have this bit on in our answer. Now, to decide which value to assign to the current bit of X, we will solve the same problem recursively for each of the groups for the next bit; let anson and ansoff be the answers of the recursive calls for the on and the off groups respectively. Note that if we assign 1 to the current bit of X, the answer will be 2i+ansoff, and if we assign 0 to the current bit of X, the answer will be 2i+anson, where i is the current bit. So, simply we will choose the minimum of these two cases for our answer to be 2i+min(anson,ansoff).
// Time complexity: O(nlog(maxai))

package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(c []uint, bit int) uint {
	var l, r []uint
	var rec_l, rec_r uint

	if bit < 0 {
		return 0
	}

	for _, i := range c {
		if ((i >> bit) & 1) == 0 {
			l = append(l, i)
		} else {
			r = append(r, i)
		}
	}

	if len(l) == 0 {
		return solve(r, bit-1)
	}

	if len(r) == 0 {
		return solve(l, bit-1)
	}

	rec_l = solve(l, bit-1)
	rec_r = solve(r, bit-1)

	if rec_l < rec_r {
		return rec_l + (1 << bit)
	} else {
		return rec_r + (1 << bit)
	}
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var n, i uint
	var a []uint

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		a = make([]uint, n)

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}

		fmt.Fprintln(writer, solve(a, 30))
	}

	writer.Flush()
}
