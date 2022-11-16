// https://codeforces.com/problemset/problem/1620/B

// 1620B - Triangles on a Rectangle

// A rectangle with its opposite corners in (0,0) and (w,h) and sides parallel to the axes is drawn on a plane.
// You are given a list of lattice points such that each point lies on a side of a rectangle but not in its corner. Also, there are at least two points on every side of a rectangle.
// Your task is to choose three points in such a way that:
//  • exactly two of them belong to the same side of a rectangle;
//  • the area of a triangle formed by them is maximum possible.
// Print the doubled area of this triangle. It can be shown that the doubled area of any triangle formed by lattice points is always an integer.

// Input
// The first line contains a single integer t (1≤t≤10^4) — the number of testcases.
// The first line of each testcase contains two integers w and h (3≤w,h≤10^6) — the coordinates of the corner of a rectangle.
// The next two lines contain the description of the points on two horizontal sides. First, an integer k (2≤k≤2⋅10^5) — the number of points. Then, k integers x1<x2<⋯<xk (0<xi<w) — the x coordinates of the points in the ascending order. The y coordinate for the first line is 0 and for the second line is h.
// The next two lines contain the description of the points on two vertical sides. First, an integer k (2≤k≤2⋅10^5) — the number of points. Then, k integers y1<y2<⋯<yk (0<yi<h) — the y coordinates of the points in the ascending order. The x coordinate for the first line is 0 and for the second line is w.
// The total number of points on all sides in all testcases doesn't exceed 2⋅10^5.

// Output
// For each testcase print a single integer — the doubled maximum area of a triangle formed by such three points that exactly two of them belong to the same side.

// Example
// input
// 3
// 5 8
// 2 1 2
// 3 2 3 4
// 3 1 4 6
// 2 4 5
// 10 7
// 2 3 9
// 2 1 7
// 3 1 3 4
// 3 4 5 6
// 11 5
// 3 1 6 8
// 3 3 6 8
// 3 1 3 4
// 2 2 4
// output
// 25
// 42
// 35

// Note

// The points in the first testcase of the example:
//  • (1,0), (2,0);
//  • (2,8), (3,8), (4,8);
//  • (0,1), (0,4), (0,6);
//  • (5,4), (5,5).
// The largest triangle is formed by points (0,1), (0,6) and (5,4) — its area is 252. Thus, the doubled area is 25. Two points that are on the same side are: (0,1) and (0,6).

// Tutorial

// The area of a triangle is equal to its base multiplied by its height divided by 2. Let the two points that have to be on the same side of a rectangle form its base. To maximize it, let's choose such two points that are the most apart from each other — the first and the last in the list.
// Then the height will be determined by the distance from that side to the remaining point. Since there are points on all sides, the points on the opposite side are the furthest. Thus, the height is always one of h or w, depending on whether we picked the horizontal or the vertical side.
// So we have to check four options to pick the side and choose the best answer among them.

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
	var t, w, h, k, x, ans, i, j, first, last uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &w, &h)
			ans = 0

			for i = 0; i < 4; i++ {
				fmt.Fscan(reader, &k)

				for j = 0; j < k; j++ {
					fmt.Fscan(reader, &x)

					if j == 0 {
						first = x
					}

					if j+1 == k {
						last = x
					}
				}

				if i < 2 {
					if ans < (last-first)*h {
						ans = (last - first) * h
					}
				} else {
					if ans < (last-first)*w {
						ans = (last - first) * w
					}
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
