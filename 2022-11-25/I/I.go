// https://codeforces.com/contest/1228/problem/A

// 1228A - Distinct Digits
// You have two integers l and r. Find an integer x which satisfies the conditions below:
//  • l≤x≤r.
//  • All digits of x are different.
// If there are multiple answers, print any of them.

// Input
// The first line contains two integers l and r (1≤l≤r≤105).

// Output
// If an answer exists, print any of them. Otherwise, print −1.

// Examples

// input
// 121 130
// output
// 123

// input
// 98766 100000
// output
// -1

// Note

// In the first example, 123 is one of the possible answers. However, 121 can't be the answer, because there are multiple 1s on different digits.
// In the second example, there is no valid answer.

// Tutorial

// Let's see how to check if all digits of x are different. Since there can be only 10 different numbers(0 to 9) in single digit, you can count the occurrences of 10 numbers by looking all digits of x. You can count all digits by using modulo 10 or changing whole number to string.
// For example, if x=1217, then occurrence of each number will be [0,2,1,0,0,0,0,1,0,0], because there are two 1s, single 2 and single 7 in x. So 1217 is invalid number.
// Now do the same thing for all x where l≤x≤r. If you find any valid number then print it. Otherwise print −1.
// Time complexity is O((r−l)logr).

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
	var l, r, i, j uint
	var found, is_distinct bool
	var occurrences [10]uint

	for t, _ = fmt.Fscan(reader, &l, &r); t == 2; t, _ = fmt.Fscan(reader, &l, &r) {
		found = false

		for i = l; i <= r && !found; i++ {
			is_distinct = true
			occurrences = [10]uint{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

			for j = i; j > 0 && is_distinct; j /= 10 {
				occurrences[j%10]++

				if occurrences[j%10] > 1 {
					is_distinct = false
				}
			}

			if is_distinct {
				fmt.Fprintln(writer, i)
				found = true
			}
		}

		if !found {
			fmt.Fprintln(writer, -1)
		}
	}

	writer.Flush()
}
