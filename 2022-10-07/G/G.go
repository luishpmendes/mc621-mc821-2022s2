// https://codeforces.com/problemset/problem/1343/C

// 1343C - Alternating Subsequence

// Recall that the sequence b is a a subsequence of the sequence a if b can be derived from a by removing zero or more elements without changing the order of the remaining elements. For example, if a=[1,2,1,3,1,2,1], then possible subsequences are: [1,1,1,1], [3] and [1,2,1,3,1,2,1], but not [3,2,3] and [1,1,1,1,2].
// You are given a sequence a consisting of n positive and negative elements (there is no zeros in the sequence).
// Your task is to choose maximum by size (length) alternating subsequence of the given sequence (i.e. the sign of each next element is the opposite from the sign of the current element, like positive-negative-positive and so on or negative-positive-negative and so on). Among all such subsequences, you have to choose one which has the maximum sum of elements.
// In other words, if the maximum length of alternating subsequence is k then your task is to find the maximum sum of elements of some alternating subsequence of length k.
// You have to answer t independent test cases.

// Input
// The first line of the input contains one integer t (1≤t≤10^4) — the number of test cases. Then t test cases follow.
// The first line of the test case contains one integer n (1≤n≤2⋅10^5) — the number of elements in a. The second line of the test case contains n integers a1,a2,…,an (−109≤ai≤109,ai≠0), where ai is the i-th element of a.
// It is guaranteed that the sum of n over all test cases does not exceed 2⋅105 (∑n≤2⋅10^5).

// Output
// For each test case, print the answer — the maximum sum of the maximum by size (length) alternating subsequence of a.

// Example
// input
// 4
// 5
// 1 2 3 -1 -2
// 4
// -1 -2 -1 -3
// 10
// -2 8 3 8 -4 -15 5 -2 -3 1
// 6
// 1 -1000000000 1 -1000000000 1 -1000000000
// output
// 2
// -1
// 6
// -2999999997

// Note

// In the first test case of the example, one of the possible answers is [1,2,3,-1,-2].
// In the second test case of the example, one of the possible answers is [-1,-2,-1,-3].
// In the third test case of the example, one of the possible answers is [-2,8,3,8,-4,-15,5,-2,-3,1].
// In the fourth test case of the example, one of the possible answers is [1,-1000000000,1,-1000000000,1,-1000000000].

// Tutorial

// Firstly, let's extract maximum by inclusion segments of the array that consists of the numbers with the same sign. For example, if the array is [1,1,2,−1,−5,2,1,−3] then these segments are [1,1,2], [−1,−5], [2,1] and [−3]. We can do it with any "two pointers"-like algorithm. The number of these segments is the maximum possible length of the alternating subsequence because we can take only one element from each block. And as we want to maximize the sum, we need to take the maximum element from each block.

// Time complexity: O(n).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func sgn(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, sum, cur int
	var t, n, i, j uint
	var a []int

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			a = make([]int, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])
			}

			sum = 0

			for i = 0; i < n; i++ {
				cur = 0

				for cur, j = a[i], i; j < n && sgn(a[i]) == sgn(a[j]); j++ {
					if cur < a[j] {
						cur = a[j]
					}
				}

				sum += cur
				i = j - 1
			}

			fmt.Fprintln(writer, sum)
		}
	}

	writer.Flush()
}
