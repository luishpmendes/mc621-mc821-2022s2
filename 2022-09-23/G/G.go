// https://codeforces.com/problemset/problem/1549/A

// 1549A - Gregor and Cryptography

// Gregor is learning about RSA cryptography, and although he doesn't understand how RSA works, he is now fascinated with prime numbers and factoring them.
// Gregor's favorite prime number is P. Gregor wants to find two bases of P. Formally, Gregor is looking for two integers a and b which satisfy both of the following properties.
//  • P mod a = P mod b , where xmody denotes the remainder when x is divided by y, and
//  • 2 ≤ a < b ≤ P.
// Help Gregor find two bases of his favorite prime number!

// Input
// Each test contains multiple test cases. The first line contains the number of test cases t (1≤t≤1000).
// Each subsequent line contains the integer P (5≤P≤109), with P guaranteed to be prime.

// Output
// Your output should consist of t lines. Each line should consist of two integers a and b (2≤a<b≤P). If there are multiple possible solutions, print any.

// Example
// input
// 2
// 17
// 5
// output
// 3 5
// 2 4

// Note
// The first query is P = 17. a = 3 and b = 5 are valid bases in this case, because 17mod3=17mod5=2. There are other pairs which work as well.
// In the second query, with P = 5, the only solution is a = 2 and b = 4.

// Hint 1
// Fix a into a convenient constant.

// Solution
// Since P ≥ 5 and is also a prime number, we know that P − 1 is an even composite number. A even composite number is guaranteed to have at least 2 unique divisors greater than 1. Let two of these divisors be a and b. It is guaranteed that P mod a = P mod b = 1, and thus this selection is valid.
// For example, we can simply pick a = 2 and b = P − 1, and we will get a correct solution.
// The time complexity is O(Q).

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
	var t, p uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &p)
			fmt.Fprintln(writer, 2, p-1)
		}
	}

	writer.Flush()
}
