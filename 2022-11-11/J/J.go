// https://codeforces.com/contest/230/problem/A

// 230A - Dragons
// Kirito is stuck on a level of the MMORPG he is playing now. To move on in the game, he's got to defeat all n dragons that live on this level. Kirito and the dragons have strength, which is represented by an integer. In the duel between two opponents the duel's outcome is determined by their strength. Initially, Kirito's strength equals s.
// If Kirito starts duelling with the i-th (1 ≤ i ≤ n) dragon and Kirito's strength is not greater than the dragon's strength xi, then Kirito loses the duel and dies. But if Kirito's strength is greater than the dragon's strength, then he defeats the dragon and gets a bonus strength increase by yi.
// Kirito can fight the dragons in any order. Determine whether he can move on to the next level of the game, that is, defeat all dragons without a single loss.

// Input
// The first line contains two space-separated integers s and n (1 ≤ s ≤ 10^4, 1 ≤ n ≤ 10^3). Then n lines follow: the i-th line contains space-separated integers xi and yi (1 ≤ xi ≤ 10^4, 0 ≤ yi ≤ 10^4) — the i-th dragon's strength and the bonus for defeating it.

// Output
// On a single line print "YES" (without the quotes), if Kirito can move on to the next level and print "NO" (without the quotes), if he can't.

// Examples

// input
// 2 2
// 1 99
// 100 0
// output
// YES

// input
// 10 1
// 100 100
// output
// NO

// Note

// In the first sample Kirito's strength initially equals 2. As the first dragon's strength is less than 2, Kirito can fight it and defeat it. After that he gets the bonus and his strength increases to 2 + 99 = 101. Now he can defeat the second dragon and move on to the next level.
// In the second sample Kirito's strength is too small to defeat the only dragon and win.

// Tutorial

// Observe that if Kirito fights a dragon whose strength is less than Kirito's strength, then Kirito does not lose anything — in fact, he even gains a nonnegative strength increase. Taking note of this, let's for each step choose some dragon whose strength is less than Kirito's current strength, and fight it. After performing some amount of these steps we'll eventually end up in one of these two situations: either all dragons are slain (then the answer is "YES"), or only dragons whose strength is not less than Kirito's strength remain (then the answer is "NO"). On each step we can choose a suitable dragon to fight either by searching through all dragons or by sorting the dragons by strength in non-descending order in advance.
// The complexity of the solution is O(n^2) or O(n log n).

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type pair struct {
	first, second uint
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t int
	var s, n, i uint
	var dragons []pair

	for t, _ = fmt.Fscan(reader, &s, &n); t == 2; t, _ = fmt.Fscan(reader, &s, &n) {
		dragons = make([]pair, n)

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &dragons[i].first, &dragons[i].second)
		}

		sort.Slice(dragons, func(i, j int) bool {
			return dragons[i].first < dragons[j].first
		})

		for i = 0; i < n; i++ {
			if s <= dragons[i].first {
				break
			}

			s += dragons[i].second
		}

		if i == n {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}

	writer.Flush()
}
