// https://codeforces.com/problemset/problem/1656/C

// 1656C - Make Equal With Mod

// You are given an array of n non-negative integers a1,a2,…,an. You can make the following operation: choose an integer x≥2 and replace each number of the array by the remainder when dividing that number by x, that is, for all 1≤i≤n set ai to aimodx.
// Determine if it is possible to make all the elements of the array equal by applying the operation zero or more times.

// Input
// The input consists of multiple test cases. The first line contains a single integer t (1≤t≤104) — the number of test cases. Description of the test cases follows.
// The first line of each test case contains an integer n (1≤n≤105) — the length of the array.
// The second line of each test case contains n integers a1,a2,…,an (0≤ai≤109) where ai is the i-th element of the array.
// The sum of n for all test cases is at most 2⋅105.

// Output
// For each test case, print a line with YES if you can make all elements of the list equal by applying the operation. Otherwise, print NO.
// You may print each letter in any case (for example, "YES", "Yes", "yes", "yEs" will all be recognized as a positive answer).

// Example
// input
// 4
// 4
// 2 5 6 8
// 3
// 1 1 1
// 5
// 4 1 7 0 8
// 4
// 5 9 17 5
// output
// YES
// YES
// NO
// YES

// Note

// In the first test case, one can apply the operation with x=3 to obtain the array [2,2,0,2], and then apply the operation with x=2 to obtain [0,0,0,0].
// In the second test case, all numbers are already equal.
// In the fourth test case, applying the operation with x=4 results in the array [1,1,1,1].

// Tutorial

// Note that, if 1 is not present in the array, we can always make all elements equal to 0 by repeatedly applying the operation with x=max(ai) until all elements become 0, as this operation will set the elements equal to the maximum to 0, while maintaining the others intact. So the answer is YES.
// If 1 is present and there are no two consecutive numbers in the array, we can similarly apply repeatedly the operation with x=max(ai)−1 until all elements become 1, as this operation will set the elements equal to the maximum to 1, while maintaining the others intact. So the answer is again YES.
// If 1 is present in the array, and there are two consecutive numbers m,m+1 in the array, the answer is NO. Note that if we have a 0 and a 1 in the array, we won't be able to make them equal after any number of operations, and so we cannot have any operation with an x that divides one of the ai's. The rest of operations will cause that m,m+1 remain consecutive (and thus different), meaning that it is impossible to make all the array equal.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, i uint
	var a []uint
	var one, consec bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			a = make([]uint, n)
			one = false
			consec = false

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])

				if a[i] == 1 {
					one = true
				}
			}

			sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

			for i = 1; i < n && !consec; i++ {
				if a[i] == a[i-1]+1 {
					consec = true
				}
			}

			if one && consec {
				fmt.Fprintln(writer, "NO")
			} else {
				fmt.Fprintln(writer, "YES")
			}
		}
	}

	writer.Flush()
}
