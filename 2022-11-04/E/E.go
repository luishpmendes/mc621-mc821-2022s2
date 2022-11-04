// https://codeforces.com/problemset/problem/471/D

// 471D - MUH and Cube Walls

// Polar bears Menshykov and Uslada from the zoo of St. Petersburg and elephant Horace from the zoo of Kiev got hold of lots of wooden cubes somewhere. They started making cube towers by placing the cubes one on top of the other. They defined multiple towers standing in a line as a wall. A wall can consist of towers of different heights.
// Horace was the first to finish making his wall. He called his wall an elephant. The wall consists of w towers. The bears also finished making their wall but they didn't give it a name. Their wall consists of n towers. Horace looked at the bears' tower and wondered: in how many parts of the wall can he "see an elephant"? He can "see an elephant" on a segment of w contiguous towers if the heights of the towers on the segment match as a sequence the heights of the towers in Horace's wall. In order to see as many elephants as possible, Horace can raise and lower his wall. He even can lower the wall below the ground level (see the pictures to the samples for clarification).
// Your task is to count the number of segments where Horace can "see an elephant".

// Input
// The first line contains two integers n and w (1 ≤ n, w ≤ 2·10^5) — the number of towers in the bears' and the elephant's walls correspondingly. The second line contains n integers ai (1 ≤ ai ≤ 10^9) — the heights of the towers in the bears' wall. The third line contains w integers bi (1 ≤ bi ≤ 10^9) — the heights of the towers in the elephant's wall.

// Output
// Print the number of segments in the bears' wall where Horace can "see an elephant".

// Example
// input
// 13 5
// 2 4 5 5 4 3 2 2 2 3 3 2 1
// 3 4 4 3 2
// output
// 2

// Tutorial

// In this problem we are given two arrays of integers and we need to find how many times we can see second array as a subarray in the first array if we can add some arbitrary constant value to every element of the second array. Let's call these arrays a and b. As many people noticed or knew in advance this problem can be solved easily if we introduce difference arrays like that:
// aDiffi = ai - ai + 1 (for every i =  = 0..n - 1)
// If we do that with both input arrays we will receive two arrays both of which have one element less than original arrays. Then with these arrays the problem simply becomes the word search problem (though with possibly huge alphabet). This can be solved using your favourite string data structure or algorithm. Originally it was intended to look for linear solution but then we made time limit higher in case if somebody will decide to send O(NlogN) solution. I haven't seen such solutions (that is understandable) but some people tried to squeeze in a quadratic solution.
// Linear solution can be made using Z-function or KMP algorithm. In order to add a logarithmic factor you can exercise with suffix array for example. I had suffix array solution as well, but it's a lot messier than linear solution.
// There is one corner case you need to consider — when Horace's wall contains only one tower, then it matches bears' wall in every tower so the answer is n. Though for some algorithms it might not even be a corner case if you assume that empty string matches everywhere. Another error which several people did was to use actual string data structures to solve this problem, so they converted the differences to chars. This doesn't work since char can't hold the entire range an integer type can hold.
// I didn't think that switching to difference arrays will be that obvious or well-known, so I didn't expect that this problem will be solved by that many people.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func z_function(s []int) []int {
	var z []int
	var i, l, r int

	z = make([]int, len(s))

	for i, l, r = 1, 0, 0; i < len(s); i++ {
		if i <= r {
			if r+1 < z[i-l]+i {
				z[i] = r - i + 1
			} else {
				z[i] = z[i-l]
			}
		}

		for i+z[i] < len(s) && s[z[i]] == s[i+z[i]] {
			z[i]++
		}

		if i+z[i] > r+1 {
			l = i
			r = i + z[i] - 1
		}
	}

	return z
}

func find_number_of_occurences(v []int, pattern []int) int {
	if len(pattern) == 0 {
		return len(v) + 1
	}

	var str, z []int
	var result, i int

	str = make([]int, 0)
	str = append(str, pattern...)
	str = append(str, v...)

	z = z_function(str)
	result = 0

	for i = len(pattern); i < len(z); i++ {
		if z[i] >= len(pattern) {
			result++
		}
	}

	return result
}

func to_diff_array(v []int) []int {
	var result []int
	var i int
	result = make([]int, len(v)-1)

	for i = 0; i+1 < len(v); i++ {
		result[i] = v[i+1] - v[i]
	}

	return result
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var n, w, i uint
	var a, b, a_diff, b_diff []int

	for t, _ = fmt.Fscan(reader, &n, &w); t == 2; t, _ = fmt.Fscan(reader, &n, &w) {
		a = make([]int, n)
		b = make([]int, w)

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &a[i])
		}

		for i = 0; i < w; i++ {
			fmt.Fscan(reader, &b[i])
		}

		a_diff = to_diff_array(a)
		b_diff = to_diff_array(b)

		fmt.Fprintln(writer, find_number_of_occurences(a_diff, b_diff))
	}

	writer.Flush()
}
