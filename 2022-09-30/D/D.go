// https://codeforces.com/problemset/problem/1411/C

// 1411C - Peaceful Rooks

// You are given a n×n chessboard. Rows and columns of the board are numbered from 1 to n. Cell (x,y) lies on the intersection of column number x and row number y.
// Rook is a chess piece, that can in one turn move any number of cells vertically or horizontally. There are m rooks (m<n) placed on the chessboard in such a way that no pair of rooks attack each other. I.e. there are no pair of rooks that share a row or a column.
// In one turn you can move one of the rooks any number of cells vertically or horizontally. Additionally, it shouldn't be attacked by any other rook after movement. What is the minimum number of moves required to place all the rooks on the main diagonal?
// The main diagonal of the chessboard is all the cells (i,i), where 1≤i≤n.

// Input
// The first line contains the number of test cases t (1≤t≤103). Description of the t test cases follows.
// The first line of each test case contains two integers n and m — size of the chessboard and the number of rooks (2≤n≤105, 1≤m<n). Each of the next m lines contains two integers xi and yi — positions of rooks, i-th rook is placed in the cell (xi,yi) (1≤xi,yi≤n). It's guaranteed that no two rooks attack each other in the initial placement.
// The sum of n over all test cases does not exceed 105.

// Output
// For each of t test cases print a single integer — the minimum number of moves required to place all the rooks on the main diagonal.
// It can be proved that this is always possible.

// Example
// input
// 4
// 3 1
// 2 3
// 3 2
// 2 1
// 1 2
// 5 3
// 2 3
// 3 1
// 1 2
// 5 4
// 4 5
// 5 1
// 2 2
// 3 3
// output
// 1
// 3
// 4
// 2

// Note

// Possible moves for the first three test cases:
//  1. (2,3)→(2,2)
//  2. (2,1)→(2,3), (1,2)→(1,1), (2,3)→(2,2)
//  3. (2,3)→(2,4), (2,4)→(4,4), (3,1)→(3,3), (1,2)→(1,1)

// Tutorial

// Consider rooks as edges in a graph. The position (x,y) will correspond to an edge (x→y). From the condition that there're at most one edge exits leading from each vertex and at most one edge leading to each vertex, it follows that the graph decomposes into cycles, paths, and loops (edges of type v→v).
// What happens to the graph when we move the rook? The edge changes exactly one of its endpoints. By such operations, we must turn all edges into loops, and the constraint on the number of edges going in and out of a vertex must be satisfied. A path is quite easy to turn into loops, just start from one end. A cycle must first be turned into a path, which is always possible. We've only spent one extra move, it's not hard to see that this is optimal.
// The answer is the number of rooks minus the number of loops plus the number of cycles.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func make_set(p *[]uint, rank *[]uint, x uint) {
	(*p)[x] = x
	(*rank)[x] = 0
}

func link(p *[]uint, rank *[]uint, x, y uint) {
	if (*rank)[x] > (*rank)[y] {
		(*p)[y] = x
	} else {
		(*p)[x] = y

		if (*rank)[x] == (*rank)[y] {
			(*rank)[y]++
		}
	}
}

func find_set(p *[]uint, x uint) uint {
	if (*p)[x] != x {
		(*p)[x] = find_set(p, (*p)[x])
	}

	return (*p)[x]
}

func union_set(p *[]uint, rank *[]uint, x, y uint) {
	link(p, rank, find_set(p, x), find_set(p, y))
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, n, m, x, y, ans, i uint
	var p, rank []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &m)
			p = make([]uint, n)
			rank = make([]uint, n)
			ans = 0

			for i = 0; i < n; i++ {
				make_set(&p, &rank, i)
			}

			for i = 0; i < m; i++ {
				fmt.Fscan(reader, &x, &y)
				x--
				y--

				if x != y {
					if find_set(&p, x) == find_set(&p, y) {
						ans += 2
					} else {
						ans++
						union_set(&p, &rank, x, y)
					}
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
