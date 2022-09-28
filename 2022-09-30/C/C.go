// https://codeforces.com/problemset/problem/1485/A

// 1485A - Add and Divide

// You have two positive integers a and b.
// You can perform two kinds of operations:
//  • a=⌊a/b⌋ (replace a with the integer part of the division between a and b)
//  • b=b+1  (increase b by 1)
// Find the minimum number of operations required to make a=0.

// Input
// The first line contains a single integer t (1≤t≤100) — the number of test cases.
// The only line of the description of each test case contains two integers a, b (1≤a,b≤109).

// Output
// For each test case, print a single integer: the minimum number of operations required to make a=0.

// Example
// input
// 6
// 9 2
// 1337 1
// 1 1
// 50000000 4
// 991026972 997
// 1234 5678
// output
// 4
// 9
// 2
// 12
// 3
// 1

// Note

// In the first test case, one of the optimal solutions is:
//  1. Divide a by b. After this operation a=4 and b=2.
//  2. Divide a by b. After this operation a=2 and b=2.
//  3. Increase b. After this operation a=2 and b=3.
//  4. Divide a by b. After this operation a=0 and b=3.

// Tutorial

// Hint 1
// Suppose that you can use x operations of type 1 and y operations of type 2. Try to reorder the operations in such a way that a becomes the minimum possible.

// Hint 2
// You should use operations of type 2 first, then moves of type 1. How many operations do you need in the worst case? (a=109, b=1)

// Hint 3
// You need at most 30 operations. Iterate over the number of operations of type 2.

// Solution
// Notice how it is never better to increase b after dividing (⌊a/(b+1)⌋≤⌊a/b⌋).
// So we can try to increase b to a certain value and then divide a by b until it is 0. Being careful as not to do this with b<2, the number of times we divide is going to be O(loga). In particular, if you reach b≥2 (this requires at most 1 move), you need at most ⌊log2(10^9)⌋=29 moves to finish.
// Let y be the number of moves of type 2; we can try all values of y (0≤y≤30) and, for each y, check how many moves of type 1 are necessary.
// Complexity: O(log^2(a)).
// If we notice that it is never convenient to increase b over 6, we can also achieve a solution with better complexity.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func num_second_operation(a, b uint) uint {
	var result uint

	for result = 0; a > 0; result++ {
		a /= b
	}

	return result
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c int
	var t, a, b, ans, i, new_ans uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &a, &b)
			ans = 1e9

			i = b
			if i < 2 {
				i = 2
			}

			for ; i <= b+ans; i++ {
				new_ans = i - b + num_second_operation(a, i)

				if ans > new_ans {
					ans = new_ans
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
