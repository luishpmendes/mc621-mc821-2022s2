// https://codeforces.com/problemset/problem/1217/D

// 1217D - Coloring Edges
// You are given a directed graph with n vertices and m directed edges without self-loops or multiple edges.
// Let's denote the k-coloring of a digraph as following: you color each edge in one of k colors. The k-coloring is good if and only if there no cycle formed by edges of same color.
// Find a good k-coloring of given digraph with minimum possible k.

// Input
// The first line contains two integers n and m (2≤n≤5000, 1≤m≤5000) — the number of vertices and edges in the digraph, respectively.
// Next m lines contain description of edges — one per line. Each edge is a pair of integers u and v (1≤u,v≤n, u≠v) — there is directed edge from u to v in the graph.
// It is guaranteed that each ordered pair (u,v) appears in the list of edges at most once.

// Output
// In the first line print single integer k — the number of used colors in a good k-coloring of given graph.
// In the second line print m integers c1,c2,…,cm (1≤ci≤k), where ci is a color of the i-th edge (in order as they are given in the input).
// If there are multiple answers print any of them (you still have to minimize k).

// Examples

// input
// 4 5
// 1 2
// 1 3
// 3 4
// 2 4
// 1 4
// output
// 1
// 1 1 1 1 1

// input
// 3 3
// 1 2
// 2 3
// 3 1
// output
// 2
// 1 1 2

// Tutorial

// Let's run dfs on the graph and color all "back edges" ((u,v) is back edge if there is a path from v to u by edges from dfs tree) in black and all other edges in white.
// It can be proven that any cycle will have at least one white edge and at least black edge. Moreover each back edge connected with at least one cycle (path from v to u and (u,v) back edge). So the coloring we got is exactly the answer.
// How to prove that any cycle have at least one edge of both colors? Let's look only at edges from dfs trees. We can always renumerate vertices in such way that index of parent id(p) is bigger than the index of any its child id(c). We can process and assign id(p) with minimal free number after we processed all its children.
// Now we can note that for any white edge (u,v) (not only tree edge) condition id(u)>id(v) holds (because of properties of dfs: forward edges are obvious; cross edge (u,v) becomes cross because dfs at first processed vertex v and u after that, so id(v)<id(u)). And for each back edge (u,v) it's true that id(u)<id(v).
// Since any cycle have both id(u)>id(v) and id(u)<id(v) situations, profit!

package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	first, second uint
}

func dfs(u uint, adj [][]pair, color *[]uint, ans *[]uint, is_cyclic *bool) {
	var i int
	var v, id uint

	(*color)[u] = 1

	for i = 0; i < len(adj[u]); i++ {
		v = adj[u][i].first
		id = adj[u][i].second

		if (*color)[v] == 0 {
			dfs(v, adj, color, ans, is_cyclic)
			(*ans)[id] = 1
		} else if (*color)[v] == 1 {
			(*ans)[id] = 2
			*is_cyclic = true
		} else { // (*color)[v] == 2
			(*ans)[id] = 1
		}
	}

	(*color)[u] = 2
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var n, m, u, v, i uint
	var adj [][]pair
	var color, ans []uint
	var is_cyclic bool

	for t, _ = fmt.Fscan(reader, &n, &m); t == 2; t, _ = fmt.Fscan(reader, &n, &m) {
		adj = make([][]pair, n)
		color = make([]uint, n)
		ans = make([]uint, m)
		is_cyclic = false

		for i = 0; i < n; i++ {
			adj[i] = make([]pair, 0)
			color[i] = 0
		}

		for i = 0; i < m; i++ {
			ans[i] = 0
			fmt.Fscan(reader, &u, &v)
			u--
			v--
			adj[u] = append(adj[u], pair{v, i})
		}

		for i = 0; i < n; i++ {
			if color[i] == 0 {
				dfs(i, adj, &color, &ans, &is_cyclic)
			}
		}

		if is_cyclic {
			fmt.Fprintln(writer, 2)
		} else {
			fmt.Fprintln(writer, 1)
		}

		for i = 0; i+1 < m; i++ {
			fmt.Fprint(writer, ans[i], " ")
		}

		fmt.Fprintln(writer, ans[i])
	}

	writer.Flush()
}
