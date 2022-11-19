// https://codeforces.com/contest/1257/problem/C

// 1257C - Dominated Subarray
// Let's call an array t dominated by value v in the next situation.
// At first, array t should have at least 2 elements. Now, let's calculate number of occurrences of each number num in t and define it as occ(num). Then t is dominated (by v) if (and only if) occ(v)>occ(v′) for any other number v′. For example, arrays [1,2,3,4,5,2], [11,11] and [3,2,3,2,3] are dominated (by 2, 11 and 3 respectevitely) but arrays [3], [1,2] and [3,3,2,2,1] are not.
// Small remark: since any array can be dominated only by one number, we can not specify this number and just say that array is either dominated or not.
// You are given array a1,a2,…,an. Calculate its shortest dominated subarray or say that there are no such subarrays.
// The subarray of a is a contiguous part of the array a, i. e. the array ai,ai+1,…,aj for some 1≤i≤j≤n.

// Input
// The first line contains single integer T (1≤T≤1000) — the number of test cases. Each test case consists of two lines.
// The first line contains single integer n (1≤n≤2⋅10^5) — the length of the array a.
// The second line contains n integers a1,a2,…,an (1≤ai≤n) — the corresponding values of the array a.
// It's guaranteed that the total length of all arrays in one test doesn't exceed 2⋅10^5.

// Output
// Print T integers — one per test case. For each test case print the only integer — the length of the shortest dominated subarray, or −1 if there are no such subarrays.

// Example
// input
// 4
// 1
// 1
// 6
// 1 2 3 4 5 1
// 9
// 4 1 2 4 5 4 3 2 1
// 4
// 3 3 3 3
// output
// -1
// 6
// 3
// 2

// Note

// In the first test case, there are no subarrays of length at least 2, so the answer is −1.
// In the second test case, the whole array is dominated (by 1) and it's the only dominated subarray.
// In the third test case, the subarray a4,a5,a6 is the shortest dominated subarray.
// In the fourth test case, all subarrays of length more than one are dominated.

// Tutorial

// At first, let's prove that the shortest dominated subarray has pattern like v,c1,c2,…,ck,v with k≥0 and dominated by value v. Otherwise, we can decrease its length by erasing an element from one of its ends which isn't equal to v and it'd still be dominated.
// Now we should go over all pairs of the same numbers and check its subarrays... Or not? Let's look closely at the pattern: if v and all ci are pairwise distinct then the pattern is dominated subarray itself. Otherwise, we can find in our pattern other shorter pattern and either the found pattern is dominated or it has the pattern inside it and so on.
// What does it mean? It means that the answer is just the shortest pattern we can find. And all we need to find is the shortest subarray with the same first and last elements or just distance between two consecutive occurrences of each number. We can do it by iterating over current position i and keeping track of the last occurrence of each number in some array lst[v]. Then the current distance is i−lst[a[i]]+1.
// The total complexity is O(n).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, t, n, i, ans int
	var a []int
	var lst []int

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			a = make([]int, n)
			lst = make([]int, n+1)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])
				lst[i] = -1
			}

			lst[n] = -1

			ans = n + 1

			for i = 0; i < n; i++ {
				if lst[a[i]] != -1 && ans > i-lst[a[i]]+1 {
					ans = i - lst[a[i]] + 1
				}

				lst[a[i]] = i
			}

			if ans == n+1 {
				ans = -1
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
