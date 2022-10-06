// https://codeforces.com/problemset/problem/1352/A

// 1352A - Sum of Round Numbers

// A positive (strictly greater than zero) integer is called round if it is of the form d00...0. In other words, a positive integer is round if all its digits except the leftmost (most significant) are equal to zero. In particular, all numbers from 1 to 9 (inclusive) are round.
// For example, the following numbers are round: 4000, 1, 9, 800, 90. The following numbers are not round: 110, 707, 222, 1001.
// You are given a positive integer n (1≤n≤104). Represent the number n as a sum of round numbers using the minimum number of summands (addends). In other words, you need to represent the given number n as a sum of the least number of terms, each of which is a round number.

// Input
// The first line contains an integer t (1≤t≤10^4) — the number of test cases in the input. Then t test cases follow.
// Each test case is a line containing an integer n (1≤n≤10^4).

// Output
// Print t answers to the test cases. Each answer must begin with an integer k — the minimum number of summands. Next, k terms must follow, each of which is a round number, and their sum is n. The terms can be printed in any order. If there are several answers, print any of them.

// Example
// input
// 5
// 5009
// 7
// 9876
// 10000
// 10
// output
// 2
// 5000 9
// 1
// 7
// 4
// 800 70 6 9000
// 1
// 10000
// 1
// 10

// Tutorial

// Firstly, we need to understand the minimum amount of round numbers we need to represent n. It equals the number of non-zero digits in n. Why? Because we can "remove" exactly one non-zero digit in n using exactly one round number (so we need at most this amount of round numbers) and, on the other hand, the sum of two round numbers has at most two non-zero digits (the sum of three round numbers has at most three non-zero digits and so on) so this is useless to try to remove more than one digit using the sum of several round numbers.
// So we need to find all digits of n and print the required number for each of these digits. For example, if n=103 then n=1⋅10^2+0⋅10^1+3⋅10^0, so we need two round numbers: 1⋅10^2 and 3⋅10^0.

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
	var t, n, power uint
	var v []uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			v = make([]uint, 0)

			for power = 1; n > 0; power *= 10 {
				if n%10 > 0 {
					v = append(v, n%10*power)
				}

				n /= 10
			}

			fmt.Fprintln(writer, len(v))

			for _, x := range v {
				fmt.Fprint(writer, x, " ")
			}

			fmt.Fprintln(writer)
		}
	}

	writer.Flush()
}
