// https://codeforces.com/problemset/problem/1426/F

// 1426F - Number of Subsequences

// You are given a string s consisting of lowercase Latin letters "a", "b" and "c" and question marks "?".
// Let the number of question marks in the string s be k. Let's replace each question mark with one of the letters "a", "b" and "c". Here we can obtain all 3k possible strings consisting only of letters "a", "b" and "c". For example, if s="ac?b?c" then we can obtain the following strings: ["acabac", "acabbc", "acabcc", "acbbac", "acbbbc", "acbbcc", "accbac", "accbbc", "accbcc"].
// Your task is to count the total number of subsequences "abc" in all resulting strings. Since the answer can be very large, print it modulo 10^9+7.
// A subsequence of the string t is such a sequence that can be derived from the string t after removing some (possibly, zero) number of letters without changing the order of remaining letters. For example, the string "baacbc" contains two subsequences "abc" — a subsequence consisting of letters at positions (2,5,6) and a subsequence consisting of letters at positions (3,5,6).

// Input
// The first line of the input contains one integer n (3≤n≤200000) — the length of s.
// The second line of the input contains the string s of length n consisting of lowercase Latin letters "a", "b" and "c" and question marks"?".

// Output
// Print the total number of subsequences "abc" in all strings you can obtain if you replace all question marks with letters "a", "b" and "c", modulo 10^9+7.

// Examples

// input
// 6
// ac?b?c
// output
// 24

// input
// 7
// ???????
// output
// 2835

// input
// 9
// cccbbbaaa
// output
// 0

// input
// 5
// a???c
// output
// 46

// Note

// In the first example, we can obtain 9 strings:
//  • "acabac" — there are 2 subsequences "abc",
//  • "acabbc" — there are 4 subsequences "abc",
//  • "acabcc" — there are 4 subsequences "abc",
//  • "acbbac" — there are 2 subsequences "abc",
//  • "acbbbc" — there are 3 subsequences "abc",
//  • "acbbcc" — there are 4 subsequences "abc",
//  • "accbac" — there is 1 subsequence "abc",
//  • "accbbc" — there are 2 subsequences "abc",
//  • "accbcc" — there are 2 subsequences "abc".
// So, there are 2+4+4+2+3+4+1+2+2=24 subsequences "abc" in total.

// Tutorial

// There are several more or less complicated combinatorial solutions to this problem, but I will describe a dynamic programming one which, I think, is way easier to understand and to implement.
// Suppose we have fixed the positions of a, b and c that compose the subsequence (let these positions be pa, pb and pc). How many strings contain the required subsequence on these positions? Obviously, if some of these characters is already not a question mark and does not match the expected character on that position, the number of strings containing the subsequence on that position is 0. Otherwise, since we have fixed three characters, all question marks on other positions can be anything we want — so the number of such strings is 3x, where x is the number of question marks on positions other than pa, pb and pc. It allows us to write an O(n3) solution by iterating on pa, pb and pc, and for every such triple, calculating the number of strings containing the required subsequence on those positions.
// But that's too slow. Let's notice that, for every such subsequence, the number of strings containing it is 3k−qPos(pa,pb,pc), where qPos(pa,pb,pc) is the number of positions from pa,pb,pc that contain a question mark. So, for each integer i from 0 to 3, let's calculate the number of subsequences matching abc that contain exactly i question marks — and that will allow us to solve the problem faster.
// How can we calculate the required number of subsequences for every i? In my opinion, the simplest way is dynamic programming: let dpi,j,k be the number of subsequences of s that end up in position i, match j first characters of abc and contain k question marks. The transitions in this dynamic programming are quadratic (since we have to iterate on the next/previous position from the subsequence), but can be sped up to linear if we rewrite dpi,j,k as the number of subsequences of s that end up in position not later than i, match j first characters of abc and contain k question marks. Each transition is either to take the current character or to skip it, so they can be modeled in O(1), and overall this dynamic programming solution works in O(n).

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
	var mod, n, i, j, k, cntQ, ans uint
	var s string
	var pow3 []uint
	var dp [][][]uint

	mod = 1000000007

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		fmt.Fscan(reader, &s)
		cntQ = 0
		pow3 = make([]uint, n+1)

		for _, c := range s {
			if c == '?' {
				cntQ++
			}
		}

		pow3[0] = 1

		for i = 1; i <= n; i++ {
			pow3[i] = (pow3[i-1] * 3) % mod
		}

		dp = make([][][]uint, n+1)

		for i = 0; i <= n; i++ {
			dp[i] = make([][]uint, 4)

			for j = 0; j < 4; j++ {
				dp[i][j] = make([]uint, 4)
			}
		}

		dp[0][0][0] = 1

		for i = 0; i < n; i++ {
			for j = 0; j < 4; j++ {
				for k = 0; k < 4; k++ {
					if dp[i][j][k] > 0 {
						dp[i+1][j][k] = (dp[i+1][j][k] + dp[i][j][k]) % mod

						if j < 3 {
							if s[i] == '?' {
								dp[i+1][j+1][k+1] = (dp[i+1][j+1][k+1] + dp[i][j][k]) % mod
							}

							if s[i] == "abc"[j] {
								dp[i+1][j+1][k] = (dp[i+1][j+1][k] + dp[i][j][k]) % mod
							}
						}
					}
				}
			}
		}

		ans = 0

		for i = 0; i <= 3; i++ {
			if cntQ >= i {
				ans = (ans + ((dp[n][3][i] * pow3[cntQ-i]) % mod)) % mod
			}
		}

		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
