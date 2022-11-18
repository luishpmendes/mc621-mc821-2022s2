// https://www.spoj.com/problems/HS12MBR/

// HS12MBR - Minimum Bounding Rectangle
// Compute the Minimum Bounding Rectangle (MBR) that surrounds the given set of 2D objects, i.e., the axis-aligned rectangle, which contains all of the specified objects and is the one with minimum area among all rectangles with this property.

// Input
// First, you are given t (t<100) - the number of test cases.
// Each of the test cases starts with one integer n (n < 100) - the number of objects in the set. In the successive n lines, the descriptions of the objects follow.
// Each object is described by one character and some parameters:
//  • a point: p x y, where x and y are point coordinates.
//  • a circle: c x y r, where x and y are the center coordinates and r is the radius of the circle.
//  • a line segment: l x1 y1 x2 y2, where xi, yi are the coordinates of the endpoints of the line.
// Successive test cases are separated by an empty line.

// Output
// For each of the test cases output four numbers - the coordinates of the two points that correspond to the lower left and the upper right corner of the MBR, in the following order: first the x-coordinate of the lower left corner, then the y-coordinate of the lower left corner, the x-coordinate of the upper right corner and the y-coordinate of upper right corner.
// You can assume that all object parameters are integers and that -1000 -1000 1000 1000 is a bounding rectangle for all of them.

// Example
// input
// 3
// 1
// p 3 3
//
// 2
// c 10 10 20
// c 20 20 10
//
// 1
// l 0 0 100 20
// output
// 3 3 3 3
// -10 -10 30 30
// 0 0 100 20

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
	var x, y, r, x2, y2, min_x, min_y, max_x, max_y int
	var tp string

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {

		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &tp)

				if tp == "p" {
					fmt.Fscan(reader, &x, &y)

					if i == 0 || min_x > x {
						min_x = x
					}

					if i == 0 || min_y > y {
						min_y = y
					}

					if i == 0 || max_x < x {
						max_x = x
					}

					if i == 0 || max_y < y {
						max_y = y
					}
				} else if tp == "c" {
					fmt.Fscan(reader, &x, &y, &r)

					if i == 0 || min_x > x-r {
						min_x = x - r
					}

					if i == 0 || min_y > y-r {
						min_y = y - r
					}

					if i == 0 || max_x < x+r {
						max_x = x + r
					}

					if i == 0 || max_y < y+r {
						max_y = y + r
					}
				} else { // tp == "l"
					fmt.Fscan(reader, &x, &y, &x2, &y2)

					if x < x2 {
						if i == 0 || min_x > x {
							min_x = x
						}

						if i == 0 || max_x < x2 {
							max_x = x2
						}
					} else {
						if i == 0 || min_x > x2 {
							min_x = x2
						}

						if i == 0 || max_x < x {
							max_x = x
						}
					}

					if y < y2 {
						if i == 0 || min_y > y {
							min_y = y
						}

						if i == 0 || max_y < y2 {
							max_y = y2
						}
					} else {
						if i == 0 || min_y > y2 {
							min_y = y2
						}

						if i == 0 || max_y < y {
							max_y = y
						}
					}
				}
			}

			fmt.Fprintln(writer, min_x, min_y, max_x, max_y)
		}
	}

	writer.Flush()
}
