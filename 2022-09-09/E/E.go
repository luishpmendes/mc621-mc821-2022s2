// https://codeforces.com/problemset/problem/152/C

// 152C - Pocket Book

// One day little Vasya found mom's pocket book. The book had n names of her friends and unusually enough, each name was exactly m letters long. Let's number the names from 1 to n in the order in which they are written.
// As mom wasn't home, Vasya decided to play with names: he chose three integers i, j, k (1 ≤ i < j ≤ n, 1 ≤ k ≤ m), then he took names number i and j and swapped their prefixes of length k. For example, if we take names "CBDAD" and "AABRD" and swap their prefixes with the length of 3, the result will be names "AABAD" and "CBDRD".
// You wonder how many different names Vasya can write instead of name number 1, if Vasya is allowed to perform any number of the described actions. As Vasya performs each action, he chooses numbers i, j, k independently from the previous moves and his choice is based entirely on his will. The sought number can be very large, so you should only find it modulo 1000000007 (109 + 7).

// Input
// The first input line contains two integers n and m (1 ≤ n, m ≤ 100) — the number of names and the length of each name, correspondingly. Then n lines contain names, each name consists of exactly m uppercase Latin letters.

// Output
// Print the single number — the number of different names that could end up in position number 1 in the pocket book after the applying the procedures described above. Print the number modulo 1000000007 (109 + 7).

// Examples

// input
// 2 3
// AAB
// BAA
// output
// 4

// input
// 4 5
// ABABA
// BCGDG
// AAAAA
// YABSA
// output
// 216

// Note

// In the first sample Vasya can get the following names in the position number 1: "AAB", "AAA", "BAA" and "BAB".

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, n, m, i, j, ans, cnt int
	var s string
	var v [][]bool

	for t, _ = fmt.Fscan(reader, &n, &m); t == 2; t, _ = fmt.Fscan(reader, &n, &m) {
		v = make([][]bool, m)

		for j = 0; j < m; j++ {
			v[j] = make([]bool, 'Z'-'A'+1)
		}

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &s)

			for j = 0; j < m; j++ {
				v[j][s[j]-'A'] = true
			}
		}

		ans = 1

		for j = 0; j < m; j++ {
			cnt = 0

			for i = 0; i < 'Z'-'A'+1; i++ {
				if v[j][i] {
					cnt++
				}
			}

			ans = (ans * cnt) % 1000000007
		}

		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
