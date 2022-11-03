// https://codeforces.com/problemset/problem/1684/A

// 1684A. Digit Minimization

// There is an integer n without zeros in its decimal representation. Alice and Bob are playing a game with this integer. Alice starts first. They play the game in turns.
// On her turn, Alice must swap any two digits of the integer that are on different positions. Bob on his turn always removes the last digit of the integer. The game ends when there is only one digit left.
// You have to find the smallest integer Alice can get in the end, if she plays optimally.

// Input
// The input consists of multiple test cases. The first line contains a single integer t (1≤t≤10^4) — the number of test cases. Description of the test cases follows.
// The first and the only line of each test case contains the integer n (10≤n≤10^9) — the integer for the game. n does not have zeros in its decimal representation.

// Output
// For each test case output a single integer — the smallest integer Alice can get in the end of the game.

// Example
// input
// 3
// 12
// 132
// 487456398
// output
// 2
// 1
// 3

// Note

// In the first test case Alice has to swap 1 and 2. After that Bob removes the last digit, 1, so the answer is 2.
// In the second test case Alice can swap 3 and 1: 312. After that Bob deletes the last digit: 31. Then Alice swaps 3 and 1: 13 and Bob deletes 3, so the answer is 1.

// Tutorial

// Let k be the length of n. Let ni be the i-th digit of n (1-indexation from the left).
//  • k=1
//    The game ends immediately so the answer is n itself.
//  • k=2
//    Alice should make the first move and she has to swap n1 and n2. After that Bob removes n1 and in the end there is only n2.
//  • k≥3
//    Alice can make swaps in such way that when there are only two digits left the second digit will be the maximal digit of n. Then she will make a swap and the maximal digit will be on the first position. The other one will be removed by Bob. This way she can always get the maximal digit of n in the end of the game.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, i int
	var t uint
	var n string
	var ans byte

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)

			if len(n) == 2 {
				ans = n[1]
			} else {
				ans = n[0]

				for i = 1; i < len(n); i++ {
					if ans > n[i] {
						ans = n[i]
					}
				}
			}

			fmt.Fprintln(writer, ans-'0')
		}
	}

	writer.Flush()
}
