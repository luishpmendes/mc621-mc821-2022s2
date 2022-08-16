// https://codeforces.com/problemset/problem/1695/E

// 1695E - Ambiguous Dominoes

// Polycarp and Monocarp are both solving the same puzzle with dominoes. They are given the same set of n dominoes, the i-th of which contains two numbers xi and yi. They are also both given the same m by k grid of values aij such that m⋅k=2n.
// The puzzle asks them to place the n dominoes on the grid in such a way that none of them overlap, and the values on each domino match the aij values that domino covers. Dominoes can be rotated arbitrarily before being placed on the grid, so the domino (xi,yi) is equivalent to the domino (yi,xi).
// They have both solved the puzzle, and compared their answers, but noticed that not only did their solutions not match, but none of the n dominoes were in the same location in both solutions! Formally, if two squares were covered by the same domino in Polycarp's solution, they were covered by different dominoes in Monocarp's solution. The diagram below shows one potential a grid, along with the two players' solutions.
// Polycarp and Monocarp remember the set of dominoes they started with, but they have lost the grid a. Help them reconstruct one possible grid a, along with both of their solutions, or determine that no such grid exists.

// Input
// The first line contains a single integer n (1≤n≤3⋅105).

// Output
// If there is no solution, print a single integer −1.
// Otherwise, print m and k, the height and width of the puzzle grid, on the first line of output. These should satisfy m⋅k=2n.
// The i-th of the next m lines should contain k integers, the j-th of which is aij.
// The next m lines describe Polycarp's solution. Print m lines of k characters each. For each square, if it is covered by the upper half of a domino in Polycarp's solution, it should contain a "U". Similarly, if it is covered by the bottom, left, or right half of a domino, it should contain "D", "L", or "R", respectively.
// The next m lines should describe Monocarp's solution, in the same format as Polycarp's solution.
// If there are multiple answers, print any.

// Examples

// input
// 1
// 1 2
// output
// -1

// input
// 2
// 1 1
// 1 2
// output
// 2 2
// 2 1
// 1 1
// LR
// LR
// UU
// DD

// input
// 10
// 1 3
// 1 1
// 2 1
// 3 4
// 1 5
// 1 5
// 3 1
// 2 4
// 3 3
// 4 1
// output
// 4 5
// 1 2 5 1 5
// 3 4 1 3 1
// 1 2 4 4 1
// 1 3 3 3 1
// LRULR
// LRDLR
// ULRLR
// DLRLR
// UULRU
// DDUUD
// LRDDU
// LRLRD

package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(adj [][]struct{ v, i int }, i int, used *[]bool, usedEdges *[]bool, lst *[]int) {
	*lst = append(*lst, i)

	if !(*used)[i] {
		(*used)[i] = true

		for _, p := range adj[i] {
			if !(*usedEdges)[p.i] {
				(*usedEdges)[p.i] = true
				dfs(adj, p.v, used, usedEdges, lst)
				*lst = append(*lst, i)
			}
		}
	}
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, n, x, y, i, j, idx, k int
	var adj [][]struct{ v, i int }
	var used, usedEdges []bool
	var ptop, pbot, mtop, mbot []byte
	var lst []int
	var flag bool
	var ans [][]int

	// Reads n and m until EOF
	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		adj = make([][]struct{ v, i int }, 2*n)
		used = make([]bool, 2*n)
		usedEdges = make([]bool, 2*n)
		ptop = make([]byte, n)
		pbot = make([]byte, n)
		mtop = make([]byte, n)
		mbot = make([]byte, n)
		ans = make([][]int, 2)
		ans[0] = make([]int, n)
		ans[1] = make([]int, n)
		lst = make([]int, 0)

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &x, &y)
			x--
			y--
			adj[x] = append(adj[x], struct{ v, i int }{y, i})
			adj[y] = append(adj[y], struct{ v, i int }{x, i})
			used[i] = false
			used[n+i] = false
			usedEdges[i] = false
			usedEdges[n+i] = false
			ptop[i] = 'U'
			pbot[i] = 'D'
			mtop[i] = 'U'
			mbot[i] = 'D'
		}

		idx = 0
		flag = false

		for i = 0; i < 2*n && !flag; i++ {
			if !used[i] {
				dfs(adj, i, &used, &usedEdges, &lst)
				lst = lst[:len(lst)-1]
				k = len(lst) / 2

				if k != 1 {
					for j = 0; j < k; j++ {
						ans[0][idx+j] = lst[j]
						ans[1][idx+j] = lst[2*k-j-1]
					}

					for j = 0; j < k-1; j += 2 {
						ptop[idx+j] = 'L'
						pbot[idx+j] = 'L'
						ptop[idx+j+1] = 'R'
						pbot[idx+j+1] = 'R'
					}

					for j = 1; j < k-1; j += 2 {
						mtop[idx+j] = 'L'
						mbot[idx+j] = 'L'
						mtop[idx+j+1] = 'R'
						mbot[idx+j+1] = 'R'
					}

					lst = lst[:0]
					idx += k
				} else { // k == 1
					flag = true
				}
			}
		}

		if flag {
			fmt.Fprintln(writer, -1)
		} else {
			fmt.Fprintln(writer, 2, n)

			for i = 0; i < 2; i++ {
				for j = 0; j < n-1; j++ {
					fmt.Fprint(writer, ans[i][j]+1, " ")
				}

				fmt.Fprintln(writer, ans[i][n-1]+1)
			}

			fmt.Fprintln(writer, string(ptop))
			fmt.Fprintln(writer, string(pbot))
			fmt.Fprintln(writer, string(mtop))
			fmt.Fprintln(writer, string(mbot))
		}
	}

	writer.Flush()
}
