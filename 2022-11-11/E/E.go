// https://codeforces.com/contest/1251/problem/A

// 1251A - Broken Keyboard
// Recently Polycarp noticed that some of the buttons of his keyboard are malfunctioning. For simplicity, we assume that Polycarp's keyboard contains 26 buttons (one for each letter of the Latin alphabet). Each button is either working fine or malfunctioning.
// To check which buttons need replacement, Polycarp pressed some buttons in sequence, and a string s appeared on the screen. When Polycarp presses a button with character c, one of the following events happened:
//  • if the button was working correctly, a character c appeared at the end of the string Polycarp was typing;
//  • if the button was malfunctioning, two characters c appeared at the end of the string.
// For example, suppose the buttons corresponding to characters a and c are working correctly, and the button corresponding to b is malfunctioning. If Polycarp presses the buttons in the order a, b, a, c, a, b, a, then the string he is typing changes as follows: a → abb → abba → abbac → abbaca → abbacabb → abbacabba.
// You are given a string s which appeared on the screen after Polycarp pressed some buttons. Help Polycarp to determine which buttons are working correctly for sure (that is, this string could not appear on the screen if any of these buttons was malfunctioning).
// You may assume that the buttons don't start malfunctioning when Polycarp types the string: each button either works correctly throughout the whole process, or malfunctions throughout the whole process.

// Input
// The first line contains one integer t (1≤t≤100) — the number of test cases in the input.
// Then the test cases follow. Each test case is represented by one line containing a string s consisting of no less than 1 and no more than 500 lowercase Latin letters.

// Output
// For each test case, print one line containing a string res. The string res should contain all characters which correspond to buttons that work correctly in alphabetical order, without any separators or repetitions. If all buttons may malfunction, res should be empty.

// Example
// input
// 4
// a
// zzaaz
// ccff
// cbddbb
// output
// a
// z
//
// bc
//

// Tutorial

// If a key malfunctions, each sequence of presses of this key gives a string with even number of characters. So, if there is a substring consisting of odd number of equal characters c, such that it cannot be extended to the left or to the right without adding other characters, then it could not be produced by presses of button c if c was malfunctioning.
// The only thing that's left is to find all maximal by inclusion substrings consisting of the same character.

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
	var t uint
	var s string
	var is_working []bool
	var ans []byte

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &s)

			is_working = make([]bool, 26)
			ans = make([]byte, 0, 26)

			for i := 0; i < len(s); i++ {
				j := i + 1

				for j < len(s) && s[j] == s[i] {
					j++
				}

				if (j-i)%2 == 1 {
					is_working[s[i]-'a'] = true
				}

				i = j - 1
			}

			for i := 0; i < 26; i++ {
				if is_working[i] {
					ans = append(ans, byte(i)+'a')
				}
			}

			fmt.Fprintln(writer, string(ans))
		}
	}

	writer.Flush()
}
