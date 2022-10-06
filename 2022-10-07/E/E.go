// https://codeforces.com/problemset/problem/1675/B

// 1675B - Make It Increasing

// Given n integers a1,a2,…,an. You can perform the following operation on them:
//  • select any element ai (1≤i≤n) and divide it by 2 (round down). In other words, you can replace any selected element ai with the value ⌊ai2⌋ (where ⌊x⌋ is – round down the real number x).
// Output the minimum number of operations that must be done for a sequence of integers to become strictly increasing (that is, for the condition a1<a2<⋯<an to be satisfied). Or determine that it is impossible to obtain such a sequence. Note that elements cannot be swapped. The only possible operation is described above.

// For example, let n=3 and a sequence of numbers [3,6,5] be given. Then it is enough to perform two operations on it:
//  • Write the number ⌊62⌋=3 instead of the number a2=6 and get the sequence [3,3,5];
//  • Then replace a1=3 with ⌊32⌋=1 and get the sequence [1,3,5].
// The resulting sequence is strictly increasing because 1<3<5.

// Input
// The first line of the input contains an integer t (1≤t≤10^4) — the number of test cases in the input.
// The descriptions of the test cases follow.
// The first line of each test case contains a single integer n (1≤n≤30).
// The second line of each test case contains exactly n integers a1,a2,…,an (0≤ai≤2⋅10^9).

// Output
// For each test case, print a single number on a separate line — the minimum number of operations to perform on the sequence to make it strictly increasing. If a strictly increasing sequence cannot be obtained, print "-1".

// Example
// input
// 7
// 3
// 3 6 5
// 4
// 5 3 2 1
// 5
// 1 2 3 4 5
// 1
// 1000000000
// 4
// 2 8 7 5
// 5
// 8 26 5 21 10
// 2
// 5 14
// output
// 2
// -1
// 0
// 0
// 4
// 11
// 0

// Note

// The first test case is analyzed in the statement.
// In the second test case, it is impossible to obtain a strictly increasing sequence.
// In the third test case, the sequence is already strictly increasing.

// Tutorial

// We will process the elements of the sequence starting from the end of the sequence. Each element ai (1≤i≤n−1) will be divided by 2 until it is less than ai+1. If at some point it turns out that ai+1=0, it is impossible to obtain the desired sequence.

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
	var t, n, i, ans uint
	var a []uint
	var flag bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			a = make([]uint, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])
			}

			for i, flag, ans = n-1, true, 0; i > 0 && flag; i-- {
				for a[i] <= a[i-1] && flag {
					a[i-1] >>= 1
					ans++

					if a[i] == 0 {
						flag = false
					}
				}
			}

			if flag {
				fmt.Fprintln(writer, ans)
			} else {
				fmt.Fprintln(writer, -1)
			}
		}
	}

	writer.Flush()
}
