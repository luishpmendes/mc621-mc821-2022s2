// https://www.spoj.com/problems/BLINNET/

// SPOJ BLINNET - Bytelandian Blingors Network
// We have discovered the fastest communication medium Bytelandian scientists announced, and they called it blingors. The blingors are incomparably better than other media known before. Many companies in Byteland started to build blingors networks, so the information society in the kingdom of Bytes is fact! The priority is to build the core of the blingors network, joinig main cities in the country. Assume there is some number of cities that will be connected at the beginning. The cost of building blingors connection between two cities depends on many elements, but it has been successfully estimated. Your task is to design the blingors network connections between some cities in this way that between any pair of cities is a communication route. The cost of this network should be as small as possible.

// Remarks
//  • The name of the city is a string of at most 10 letters from a,...,z.
//  • The cost of the connection between two cities is a positive integer.
//  • The sum of all connections is not greater than 232-1.
//  • The number of cities is not greater than 10 000.

// Input

// Output
// Print a single integer — the number of ways you can choose the c prisoners.

// Example
// input
// 2
//
// 4
// gdansk
// 2
// 2 1
// 3 3
// bydgoszcz
// 3
// 1 1
// 3 1
// 4 4
// torun
// 3
// 1 3
// 2 1
// 4 1
// warszawa
// 2
// 2 4
// 3 1
//
// 3
// ixowo
// 2
// 2 1
// 3 3
// iyekowo
// 2
// 1 1
// 3 7
// zetowo
// 2
// 1 3
// 2 7
//
// output
// 3
// 4

// Tutorial

// The severity of crimes form an integer sequence. Find all the contiguous sequences without any integer greater than t. If the length of any sequence is L, then we can choose c prisoners from them in L - c + 1 ways.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type edge struct {
	u, v, w uint
}

func make_set(p *[]uint, rank *[]uint, x uint) {
	(*p)[x] = x
	(*rank)[x] = 0
}

func link_set(p *[]uint, rank *[]uint, x, y uint) {
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
	link_set(p, rank, find_set(p, x), find_set(p, y))
}

func kruskal(E []edge, n uint) []edge {
	var p = make([]uint, n)
	var rank = make([]uint, n)
	var A []edge

	sort.Slice(E, func(i, j int) bool {
		return E[i].w < E[j].w
	})

	for i := uint(0); i < n; i++ {
		make_set(&p, &rank, i)
	}

	for _, e := range E {
		if find_set(&p, e.u) != find_set(&p, e.v) {
			A = append(A, e)
			union_set(&p, &rank, e.u, e.v)
		}
	}

	return A
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var s, n, p, neigh, cost, i, ans uint
	var name []string
	var id_to_name map[uint]string
	var E, A []edge

	for t, _ = fmt.Fscan(reader, &s); t == 1; t, _ = fmt.Fscan(reader, &s) {
		for ; s > 0; s-- {
			id_to_name = make(map[uint]string)
			E = make([]edge, 0)
			fmt.Fscan(reader, &n)
			name = make([]string, n)

			for i = 1; i <= n; i++ {
				fmt.Fscan(reader, &name[i-1], &p)
				id_to_name[p] = name[i-1]

				for ; p > 0; p-- {
					fmt.Fscan(reader, &neigh, &cost)
					E = append(E, edge{i - 1, neigh - 1, cost})
				}
			}

			A = kruskal(E, n)
			ans = 0

			for _, e := range A {
				ans += e.w
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
