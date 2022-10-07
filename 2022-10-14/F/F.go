// https://codeforces.com/problemset/problem/1671/B

// 1671B - Consecutive Points Segment

// You are given n points with integer coordinates on a coordinate axis OX. The coordinate of the i-th point is xi. All points' coordinates are distinct and given in strictly increasing order.
// For each point i, you can do the following operation no more than once: take this point and move it by 1 to the left or to the right (i..e., you can change its coordinate xi to xi−1 or to xi+1). In other words, for each point, you choose (separately) its new coordinate. For the i-th point, it can be either xi−1, xi or xi+1.
// Your task is to determine if you can move some points as described above in such a way that the new set of points forms a consecutive segment of integers, i. e. for some integer l the coordinates of points should be equal to l,l+1,…,l+n−1.
// Note that the resulting points should have distinct coordinates.
// You have to answer t independent test cases.

// Input
// The first line of the input contains one integer t (1≤t≤2⋅10^4) — the number of test cases. Then t test cases follow.
// The first line of the test case contains one integer n (1≤n≤2⋅10^5) — the number of points in the set x.
// The second line of the test case contains n integers x1<x2<…<xn (1≤xi≤10^6), where xi is the coordinate of the i-th point.
// It is guaranteed that the points are given in strictly increasing order (this also means that all coordinates are distinct). It is also guaranteed that the sum of n does not exceed 2⋅10^5 (∑n≤2⋅10^5).

// Output
// For each test case, print the answer — if the set of points from the test case can be moved to form a consecutive segment of integers, print YES, otherwise print NO.

// Example
// input
// 5
// 2
// 1 4
// 3
// 1 2 3
// 4
// 1 2 3 7
// 1
// 1000000
// 3
// 2 5 6
// output
// YES
// YES
// NO
// YES
// YES

// Tutorial

// We can see that the answer is YES if and only if there are no more than two gaps of length 1 between the given points. If there is no gap, the answer is obviously YES. If there is only one gap of length 1, we can just move the left (or the right) part of the set to this gap. When there are two gaps, we can move the part before the first gap to the right and the part after the second gap to the left. Of course, if there is a gap of length at least 3 (or multiple gaps with the total length 3), we can't move the points from the left and the right part to satisfy the middle gap.
// Time complexity: O(n).

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
	var t, n, i, first_x, last_x uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)

			fmt.Fscan(reader, &first_x)

			for i = 1; i < n; i++ {
				fmt.Fscan(reader, &last_x)
			}

			if last_x <= n+first_x+1 {
				fmt.Fprintln(writer, "YES")
			} else {
				fmt.Fprintln(writer, "NO")
			}
		}
	}

	writer.Flush()
}
