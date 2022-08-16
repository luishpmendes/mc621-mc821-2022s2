// https://codeforces.com/problemset/problem/628/A

// 628A - Tennis Tournament

// A tennis tournament with n participants is running. The participants are playing by an olympic system, so the winners move on and the losers drop out.
// The tournament takes place in the following way (below, m is the number of the participants of the current round):
//  • let k be the maximal power of the number 2 such that k ≤ m,
//  • k participants compete in the current round and a half of them passes to the next round, the other m - k participants pass to the next round directly,
//  • when only one participant remains, the tournament finishes.
// Each match requires b bottles of water for each participant and one bottle for the judge. Besides p towels are given to each participant for the whole tournament.
// Find the number of bottles and towels needed for the tournament.
// Note that it's a tennis tournament so in each match two participants compete (one of them will win and the other will lose).

// Input
// The only line contains three integers n, b, p (1 ≤ n, b, p ≤ 500) — the number of participants and the parameters described in the problem statement.

// Output
// Print two integers x and y — the number of bottles and towels need for the tournament.

// Examples

// input
// 5 2 3
// output
// 20 15

// input
// 8 2 4
// output
// 35 32

// Note
// In the first example will be three rounds:
//  1. in the first round will be two matches and for each match 5 bottles of water are needed (two for each of the participants and one for the judge),
//  2. in the second round will be only one match, so we need another 5 bottles of water,
//  3. in the third round will also be only one match, so we need another 5 bottles of water.
// So in total we need 20 bottles of water.
// In the second example no participant will move on to some round directly.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, n, b, p int

	for t, _ = fmt.Fscan(reader, &n, &b, &p); t == 3; t, _ = fmt.Fscan(reader, &n, &b, &p) {
		fmt.Fprintln(writer, (n-1)*(2*b+1), n*p)
	}

	writer.Flush()
}
