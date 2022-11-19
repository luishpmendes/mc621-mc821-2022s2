// https://codeforces.com/contest/535/problem/B

// 535B - Tavas and SaDDas
// Once again Tavas started eating coffee mix without water! Keione told him that it smells awful, but he didn't stop doing that. That's why Keione told his smart friend, SaDDas to punish him! SaDDas took Tavas' headphones and told him: "If you solve the following problem, I'll return it to you."
// The problem is:
// You are given a lucky number n. Lucky numbers are the positive integers whose decimal representations contain only the lucky digits 4 and 7. For example, numbers 47, 744, 4 are lucky and 5, 17, 467 are not.
// If we sort all lucky numbers in increasing order, what's the 1-based index of n?
// Tavas is not as smart as SaDDas, so he asked you to do him a favor and solve this problem so he can have his headphones back.

// Input
// The first and only line of input contains a lucky number n (1 ≤ n ≤ 10^9).

// Output
// Print the index of n among all lucky numbers.

// Examples

// input
// 4
// output
// 1

// input
// 7
// output
// 2

// input
// 77
// output
// 6

// Tutorial
// Sol1: Consider n has x digits, f(i) =  decimal representation of binary string i, m is a binary string of size x and its i - th digit is 0 if and only if the i - th digit of n is 4. Finally, answer equals to 21 + 22 + … + 2x - 1 + f(m) + 1.
// Time complexity: O(log(n))
// Sol2: Count the number of lucky numbers less than or equal to n using bitmask (assign a binary string to each lucky number by replacing 4s with 0 and 7s with 1).
// Time complexity: O(2^log(n))

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
	var n, unit, ans uint

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		unit = 1
		ans = 0

		for n != 0 {
			ans += unit * (n % 10 / 3)
			unit *= 2
			n /= 10
		}

		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
