// https://codeforces.com/problemset/problem/485/A

// 485A - Factory

// One industrial factory is reforming working plan. The director suggested to set a mythical detail production norm. If at the beginning of the day there were x details in the factory storage, then by the end of the day the factory has to produce  (remainder after dividing x by m) more details. Unfortunately, no customer has ever bought any mythical detail, so all the details produced stay on the factory.
// The board of directors are worried that the production by the given plan may eventually stop (that means that there will be а moment when the current number of details on the factory is divisible by m).
// Given the number of details a on the first day and number m check if the production stops at some moment.

// Input
// The first line contains two integers a and m (1 ≤ a, m ≤ 105).

// Output
// Print "Yes" (without quotes) if the production will eventually stop, otherwise print "No".

// Examples

// input
// 1 5
// output
// No

// input
// 3 6
// output
// Yes

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, a, m int

	for t, _ = fmt.Fscan(reader, &a, &m); t == 2; t, _ = fmt.Fscan(reader, &a, &m) {
		for m&1 == 0 {
			m >>= 1
		}

		if a%m == 0 {
			fmt.Fprintln(writer, "Yes")
		} else {
			fmt.Fprintln(writer, "No")
		}
	}

	writer.Flush()
}
