// https://www.spoj.com/problems/CSTREET/

// CSTREET - Cobbled streets

// The municipal chronicals of an unbelievable lordly major town in a land far, far away tell the following story:
// Once upon a time the new crowned king Günther decided to visit all towns in his kingdom. The people of the unbelievable lordly major town expected that king Günther would like to see some of the most famous buildings in their town. For the lordly citizens it seemed neccessary that all streets in the town that the king would have to use had to be cobbled with stone. Unfortunately the unbelievable lordly major town had not much money at that time as they used most of their savings to erect the highest cathedral the world had ever seen.
// Roumours were afloat that the real reason for their thriftiness was not that the town treasury was empty but that many people believed that king Günther came to the throne by deceiving his father king Erwin and that in his youth he made a pact with the devil. But anyway, the citizens of the unbelievable lordly major town decided to pave only as much streets as were absolutely necessary to reach every major building.
// Can you help the citizens of the unbelievable lordly major town to find out which streets should be paved?
// It might be usefull to know that all major buildings are either at the end of a street or at an intersection. In addition to that you can assume that all buildings are connected by the given streets.

// Input
// t [number of testcases (1 <= t <= 100)]
// p [price to pave one furlong of street (positive integer)]
// n [number of main buildings in the town (1 <= n <= 1000)]
// m [number of streets in the town (1 <= m <= 300000)]
// a b c [street from building a to building b with length c (lengths are given in furlong and the buildings are numbered from 1 to n)]

// Output
// For each testcase output the price of the cheapest possibility to reach all main buildings in the city on paved streets. You can assume that the result will be smaller than 2^32.

// Example
// input
// 1
// 2
// 5
// 7
// 1 2 1
// 2 3 2
// 2 4 6
// 5 2 1
// 5 1 3
// 4 5 2
// 3 4 3
// output
// 12

// mst

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
	var c int
	var t, p, n, m, i, ans uint
	var E, A []edge

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &p, &n, &m)

			E = make([]edge, m)

			for i = 0; i < m; i++ {
				fmt.Fscan(reader, &E[i].u, &E[i].v, &E[i].w)
				E[i].u--
				E[i].v--
			}

			A = kruskal(E, n)

			ans = 0

			for _, e := range A {
				ans += e.w * p
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
