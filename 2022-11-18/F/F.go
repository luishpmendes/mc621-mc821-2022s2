// https://codeforces.com/problemset/problem/1713/A

// 1713A - Traveling Salesman Problem

// You are living on an infinite plane with the Cartesian coordinate system on it. In one move you can go to any of the four adjacent points (left, right, up, down).
// More formally, if you are standing at the point (x,y), you can:
//  • go left, and move to (x−1,y), or
//  • go right, and move to (x+1,y), or
//  • go up, and move to (x,y+1), or
//  • go down, and move to (x,y−1).
// There are n boxes on this plane. The i-th box has coordinates (xi,yi). It is guaranteed that the boxes are either on the x-axis or the y-axis. That is, either xi=0 or yi=0.
// You can collect a box if you and the box are at the same point. Find the minimum number of moves you have to perform to collect all of these boxes if you have to start and finish at the point (0,0).

// Input
// The first line contains a single integer t (1≤t≤100) — the number of test cases.
// The first line of each test case contains a single integer n (1≤n≤100) — the number of boxes.
// The i-th line of the following n lines contains two integers xi and yi (−100≤xi,yi≤100) — the coordinate of the i-th box. It is guaranteed that either xi=0 or yi=0.
// Do note that the sum of n over all test cases is not bounded.

// Output
// For each test case output a single integer — the minimum number of moves required.

// Example
// input
// 3
// 4
// 0 -2
// 1 0
// -1 0
// 0 2
// 3
// 0 2
// -3 0
// 0 -1
// 1
// 0 0
// output
// 12
// 12
// 0

// Note

// In the first test case, a possible sequence of moves that uses the minimum number of moves required is shown below.
// (0,0)→(1,0)→(1,1)→(1,2)→(0,2)→(−1,2)→(−1,1)→(−1,0)→(−1,−1)→(−1,−2)→(0,−2)→(0,−1)→(0,0)
// In the second test case, a possible sequence of moves that uses the minimum number of moves required is shown below.
// (0,0)→(0,1)→(0,2)→(−1,2)→(−2,2)→(−3,2)→(−3,1)→(−3,0)→(−3,−1)→(−2,−1)→(−1,−1)→(0,−1)→(0,0)
// In the third test case, we can collect all boxes without making any moves.

// Editorial

// Hint 1
// Do we actually need to go off the axis?

// Hint 2
// How to avoid visiting an axis more than once?

// Tutorial
// Suppose we only have boxes on the Ox+ axis, then the optimal strategy is going in the following way: (0,0),(xmax,0),(0,0). There is no way to do in less than 2⋅|xmax| moves.
// What if we have boxes on two axis? Let's assume it is Oy+, suppose we have a strategy to go in the following way: (0,0),(xmax,0),...,(0,ymax),(0,0). In this case it is optimal to fill the three dots with (0,0), which is just solving each axis independently.
// Therefore, the number of axis does not matters. For each axis that has at least one box, go from (0,0) to the farthest one, then come back to (0,0).
// Time complexity: O(n)

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
	var t, n, i uint
	var x, y, min_x, max_x, min_y, max_y int

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)

			min_x, max_x, min_y, max_y = 0, 0, 0, 0
			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &x, &y)

				if min_x > x {
					min_x = x
				}
				if max_x < x {
					max_x = x
				}
				if min_y > y {
					min_y = y
				}
				if max_y < y {
					max_y = y
				}
			}

			fmt.Fprintln(writer, 2*(max_x-min_x+max_y-min_y))
		}
	}

	writer.Flush()
}
