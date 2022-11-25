// https://codeforces.com/problemset/problem/770/C

// 770C - Online Courses In BSU
// Now you can take online courses in the Berland State University! Polycarp needs to pass k main online courses of his specialty to get a diploma. In total n courses are availiable for the passage.
// The situation is complicated by the dependence of online courses, for each course there is a list of those that must be passed before starting this online course (the list can be empty, it means that there is no limitation).
// Help Polycarp to pass the least number of courses in total to get the specialty (it means to pass all main and necessary courses). Write a program which prints the order of courses.
// Polycarp passes courses consistently, he starts the next course when he finishes the previous one. Each course can't be passed more than once.

// Input
// The first line contains n and k (1 ≤ k ≤ n ≤ 10^5) — the number of online-courses and the number of main courses of Polycarp's specialty.
// The second line contains k distinct integers from 1 to n — numbers of main online-courses of Polycarp's specialty.
// Then n lines follow, each of them describes the next course: the i-th of them corresponds to the course i. Each line starts from the integer ti (0 ≤ ti ≤ n - 1) — the number of courses on which the i-th depends. Then there follows the sequence of ti distinct integers from 1 to n — numbers of courses in random order, on which the i-th depends. It is guaranteed that no course can depend on itself.
// It is guaranteed that the sum of all values ti doesn't exceed 10^5.

// Output
// Print -1, if there is no the way to get a specialty.
// Otherwise, in the first line print the integer m — the minimum number of online-courses which it is necessary to pass to get a specialty. In the second line print m distinct integers — numbers of courses which it is necessary to pass in the chronological order of their passage. If there are several answers it is allowed to print any of them.

// Examples

// input
// 6 2
// 5 3
// 0
// 0
// 0
// 2 2 1
// 1 4
// 1 5
// output
// 5
// 1 2 3 4 5

// input
// 9 3
// 3 9 5
// 0
// 0
// 3 9 4 5
// 0
// 0
// 1 8
// 1 6
// 1 2
// 2 1 2
// output
// 6
// 1 2 9 4 5 3

// input
// 3 3
// 1 2 3
// 1 2
// 1 3
// 1 1
// output
// -1

// Note

// In the first test firstly you can take courses number 1 and 2, after that you can take the course number 4, then you can take the course number 5, which is the main. After that you have to take only the course number 3, which is the last not passed main course.

// Tutorial

// Initially, we construct an oriented graph. If it is necessary firstly to take the course j to pass the course i, add the edge from i to j to the graph. Then run the search in depth from all vertices, which correspond to main courses. In the depth search, we will simultaneously check the existence of oriented cycles (we need the array of colors marked with white-gray-black) and to build a topological sorting, simply adding vertices in response at the moment when they become black.
// Here is the code for such a search in depth:
// void dfs(int u) {
//     if (color[u] == 0) {
//         color[u] = 1;
//         for (int to: g[u])
//             dfs(to);
//         color[u] = 2;
//         ord.push_back(u);
//     } else if (color[u] == 1)
//         cycle = true;
// }
// The variable cycle will be equal to true only when there is the cycle on the main courses or those on which they depend. So if cycle=true, print -1. Otherwise ord (the array with the result of a topological sorting) will contain the answer of the problem. This follows directly from its definition.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(u uint, adj [][]uint, color *[]uint, ans *[]uint) bool {
	var result bool
	var i int
	result = true

	if (*color)[u] == 0 {
		(*color)[u] = 1

		for i = 0; i < len(adj[u]) && result; i++ {
			if !dfs(adj[u][i], adj, color, ans) {
				result = false
			}
		}

		(*color)[u] = 2
		*ans = append(*ans, u)
	} else if (*color)[u] == 1 {
		return false
	}

	return result
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var n, k, t, i, u uint
	var adj [][]uint
	var specialty, color, ans []uint
	var has_cycle bool

	for c, _ = fmt.Fscan(reader, &n, &k); c == 2; c, _ = fmt.Fscan(reader, &n, &k) {
		adj = make([][]uint, n)
		color = make([]uint, n)
		specialty = make([]uint, k)
		has_cycle = false
		ans = make([]uint, 0)

		for i = 0; i < k; i++ {
			fmt.Fscan(reader, &specialty[i])
			specialty[i]--
		}

		for i = 0; i < n; i++ {
			adj[i] = make([]uint, 0)
			color[i] = 0
			fmt.Fscan(reader, &t)

			for ; t > 0; t-- {
				fmt.Fscan(reader, &u)
				adj[i] = append(adj[i], u-1)
			}
		}

		for i = 0; i < k && !has_cycle; i++ {
			if color[specialty[i]] == 0 {
				has_cycle = !dfs(specialty[i], adj, &color, &ans)
			}
		}

		if !has_cycle {
			fmt.Fprintln(writer, len(ans))
			for i = 0; i+1 < uint(len(ans)); i++ {
				fmt.Fprint(writer, ans[i]+1, " ")
			}
			fmt.Fprintln(writer, ans[i]+1)
		} else {
			fmt.Fprintln(writer, -1)
		}
	}

	writer.Flush()
}
