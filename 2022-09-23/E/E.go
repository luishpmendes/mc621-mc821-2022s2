// https://codeforces.com/problemset/problem/817/A

// 817A - Treasure Hunt

// Captain Bill the Hummingbird and his crew recieved an interesting challenge offer. Some stranger gave them a map, potion of teleportation and said that only this potion might help them to reach the treasure.
// Bottle with potion has two values x and y written on it. These values define four moves which can be performed using the potion:
//  • (a, b) → (a + x, b + y)
//  • (a, b) → (a + x, b - y)
//  • (a, b) → (a - x, b + y)
//  • (a, b) → (a - x, b - y)
// Map shows that the position of Captain Bill the Hummingbird is (x1, y1) and the position of the treasure is (x2, y2).
// You task is to tell Captain Bill the Hummingbird whether he should accept this challenge or decline. If it is possible for Captain to reach the treasure using the potion then output "YES", otherwise "NO" (without quotes).
// The potion can be used infinite amount of times.

// Input
// The first line contains four integer numbers x1, y1, x2, y2 ( - 105 ≤ x1, y1, x2, y2 ≤ 105) — positions of Captain Bill the Hummingbird and treasure respectively.
// The second line contains two integer numbers x, y (1 ≤ x, y ≤ 105) — values on the potion bottle.

// Output
// Print "YES" if it is possible for Captain to reach the treasure using the potion, otherwise print "NO" (without quotes).

// Examples

// input
// 0 0 0 6
// 2 3
// output
// YES

// input
// 1 1 3 6
// 1 5
// output
// NO

// Note
// In the first example there exists such sequence of moves:
//  1. (0, 0) → (2, 3) — the first type of move
//  2. (2, 3) → (0, 6) — the third type of move

// Tutorial

// Firstly, let's approach this problem as if the steps were (a, b) → (a ± x, 0) and (a, b) → (0, b ± y). Then the answer is "YES" if |x1 - x2| mod x = 0 and |y1 - y2| mod y = 0.
// It's easy to see that if the answer to this problem is "NO" then the answer to the original one is also "NO".
// Let's return to the original problem and take a look at some sequence of steps. It ends in some point (xe, ye). Define cntx as |xe - x1|/x and cnty as |ye - y1|/y. The parity of cntx is the same as the parity of cnty. It is like this because every type of move changes parity of both cntx and cnty.
// So the answer is "YES" if |x1 - x2| mod x = 0, |y1 - y2| mod y = 0 and |x1 - x2|/x mod 2 = |y1 - y2|/y mod 2.
// Overall complexity: O(1).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, x1, y1, x2, y2 int
	var x, y, delta_x, delta_y uint

	for t, _ = fmt.Fscan(reader, &x1, &y1, &x2, &y2, &x, &y); t == 6; t, _ = fmt.Fscan(reader, &x1, &y1, &x2, &y2, &x, &y) {
		if x1 > x2 {
			delta_x = uint(x1 - x2)
		} else {
			delta_x = uint(x2 - x1)
		}

		if y1 > y2 {
			delta_y = uint(y1 - y2)
		} else {
			delta_y = uint(y2 - y1)
		}

		if delta_x%x == 0 && delta_y%y == 0 && (delta_x/x)%2 == (delta_y/y)%2 {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}

	writer.Flush()
}
