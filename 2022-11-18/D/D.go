// https://codeforces.com/problemset/problem/1271/C

// 1271C - Shawarma Tent

// The map of the capital of Berland can be viewed on the infinite coordinate plane. Each point with integer coordinates contains a building, and there are streets connecting every building to four neighbouring buildings. All streets are parallel to the coordinate axes.
// The main school of the capital is located in (sx,sy). There are n students attending this school, the i-th of them lives in the house located in (xi,yi). It is possible that some students live in the same house, but no student lives in (sx,sy).
// After classes end, each student walks from the school to his house along one of the shortest paths. So the distance the i-th student goes from the school to his house is |sx−xi|+|sy−yi|.
// The Provision Department of Berland has decided to open a shawarma tent somewhere in the capital (at some point with integer coordinates). It is considered that the i-th student will buy a shawarma if at least one of the shortest paths from the school to the i-th student's house goes through the point where the shawarma tent is located. It is forbidden to place the shawarma tent at the point where the school is located, but the coordinates of the shawarma tent may coincide with the coordinates of the house of some student (or even multiple students).
// You want to find the maximum possible number of students buying shawarma and the optimal location for the tent itself.

// Input
// The first line contains three integers n, sx, sy (1≤n≤200000, 0≤sx,sy≤10^9) — the number of students and the coordinates of the school, respectively.
// Then n lines follow. The i-th of them contains two integers xi, yi (0≤xi,yi≤10^9) — the location of the house where the i-th student lives. Some locations of houses may coincide, but no student lives in the same location where the school is situated.

// Output
// The output should consist of two lines. The first of them should contain one integer c — the maximum number of students that will buy shawarmas at the tent.
// The second line should contain two integers px and py — the coordinates where the tent should be located. If there are multiple answers, print any of them. Note that each of px and py should be not less than 0 and not greater than 10^9.

// Examples

// input
// 4 3 2
// 1 3
// 4 2
// 5 1
// 4 1
// output
// 3
// 4 2

// input
// 3 100 100
// 0 0
// 0 0
// 100 200
// output
// 2
// 99 100

// input
// 7 10 12
// 5 6
// 20 23
// 15 4
// 16 5
// 4 54
// 12 1
// 4 15
// output
// 4
// 10 11

// Note

// In the first example, If we build the shawarma tent in (4,2), then the students living in (4,2), (4,1) and (5,1) will visit it.
// In the second example, it is possible to build the shawarma tent in (1,1), then both students living in (0,0) will visit it.

// Tutorial

// Suppose that the point (tx,ty) is the answer. If the distance between this point t and the school is greater than 1, then there exists at least one point (Tx,Ty) such that:
//  • the distance between T and the school is exactly 1;
//  • T lies on the shortest path between the school and t.
// Now we claim that the T can also be the optimal answer. That's because, if there exists a shortest path from the school to some point i going through t, the shortest path from s to t going through T can be extended to become the shortest path to i.
// So we only need to check four points adjacent to the school as possible answers. To check whether a point (ax,ay) lies on the shortest path from (bx,by) to (cx,cy), we need to verify that min(bx,cx)≤ax≤max(bx,cx) and min(by,cy)≤ay≤max(by,cy).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var n, s_x, s_y, x, y, cnt_n, cnt_s, cnt_e, cnt_w, max_cnt uint

	for t, _ = fmt.Fscan(reader, &n, &s_x, &s_y); t == 3; t, _ = fmt.Fscan(reader, &n, &s_x, &s_y) {
		cnt_n, cnt_s, cnt_e, cnt_w = 0, 0, 0, 0

		for ; n > 0; n-- {
			fmt.Fscan(reader, &x, &y)

			if y > s_y {
				cnt_n++
			} else if y < s_y {
				cnt_s++
			}
			if x > s_x {
				cnt_e++
			} else if x < s_x {
				cnt_w++
			}
		}

		max_cnt = cnt_n

		if cnt_s > max_cnt {
			max_cnt = cnt_s
		}

		if cnt_e > max_cnt {
			max_cnt = cnt_e
		}

		if cnt_w > max_cnt {
			max_cnt = cnt_w
		}

		fmt.Fprintln(writer, max_cnt)

		if cnt_n == max_cnt {
			fmt.Fprintln(writer, s_x, s_y+1)
		} else if cnt_s == max_cnt {
			fmt.Fprintln(writer, s_x, s_y-1)
		} else if cnt_e == max_cnt {
			fmt.Fprintln(writer, s_x+1, s_y)
		} else if cnt_w == max_cnt {
			fmt.Fprintln(writer, s_x-1, s_y)
		}
	}

	writer.Flush()
}
