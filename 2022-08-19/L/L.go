// https://codeforces.com/problemset/problem/56/E

// 56E - Domino Principle

// Vasya is interested in arranging dominoes. He is fed up with common dominoes and he uses the dominoes of different heights. He put n dominoes on the table along one axis, going from left to right. Every domino stands perpendicular to that axis so that the axis passes through the center of its base. The i-th domino has the coordinate xi and the height hi. Now Vasya wants to learn for every domino, how many dominoes will fall if he pushes it to the right. Help him do that.
// Consider that a domino falls if it is touched strictly above the base. In other words, the fall of the domino with the initial coordinate x and height h leads to the fall of all dominoes on the segment [x + 1, x + h - 1].

// Input
// The first line contains integer n (1 ≤ n ≤ 105) which is the number of dominoes. Then follow n lines containing two integers xi and hi ( - 108 ≤ xi ≤ 108, 2 ≤ hi ≤ 108) each, which are the coordinate and height of every domino. No two dominoes stand on one point.

// Output
// Print n space-separated numbers zi — the number of dominoes that will fall if Vasya pushes the i-th domino to the right (including the domino itself).

// Examples

// input
// 4
// 16 5
// 20 5
// 10 10
// 18 2
// output
// 3 1 4 1

// input
// 4
// 0 10
// 1 5
// 9 10
// 15 10
// output
// 4 1 2 1

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
	var t, n, i int
	var v []struct{ x, h, i int }
	var ans, stack []int

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		v = make([]struct{ x, h, i int }, n)
		ans = make([]int, n)
		stack = make([]int, 0)

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &v[i].x, &v[i].h)
			v[i].i = i
		}

		sort.Slice(v,
			func(i, j int) bool {
				return v[i].x < v[j].x
			})

		for i = n; i > 0; i-- {
			ans[v[i-1].i] = 1

			for len(stack) > 0 && v[i-1].x+v[i-1].h > v[stack[len(stack)-1]].x {
				ans[v[i-1].i] += ans[v[stack[len(stack)-1]].i]
				stack = stack[:len(stack)-1]
			}

			stack = append(stack, i-1)
		}

		for i = 0; i < n-1; i++ {
			fmt.Fprint(writer, ans[i], " ")
		}

		fmt.Fprintln(writer, ans[n-1])
	}

	writer.Flush()
}
