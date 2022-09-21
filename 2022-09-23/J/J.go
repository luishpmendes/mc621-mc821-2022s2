// https://codeforces.com/contest/420/problem/A

// 420A - Start Up

// Recently, a start up by two students of a state university of city F gained incredible popularity. Now it's time to start a new company. But what do we call it?
// The market analysts came up with a very smart plan: the name of the company should be identical to its reflection in a mirror! In other words, if we write out the name of the company on a piece of paper in a line (horizontally, from left to right) with large English letters, then put this piece of paper in front of the mirror, then the reflection of the name in the mirror should perfectly match the line written on the piece of paper.
// There are many suggestions for the company name, so coming up to the mirror with a piece of paper for each name wouldn't be sensible. The founders of the company decided to automatize this process. They asked you to write a program that can, given a word, determine whether the word is a 'mirror' word or not.

// Input
// The first line contains a non-empty name that needs to be checked. The name contains at most 105 large English letters. The name will be written with the next sans serif font:

// Output
// Print 'YES' (without the quotes), if the given name matches its mirror reflection. Otherwise, print 'NO' (without the quotes).

// Examples

// input
// AHA
// output
// YES

// input
// Z
// output
// NO

// input
// XO
// output
// NO

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
	var name string
	var result bool
	var reflection = map[rune]rune{
		'A': 'A',
		'B': ' ',
		'C': ' ',
		'D': ' ',
		'E': ' ',
		'F': ' ',
		'G': ' ',
		'H': 'H',
		'I': 'I',
		'J': ' ',
		'K': ' ',
		'L': ' ',
		'M': 'M',
		'N': ' ',
		'O': 'O',
		'P': ' ',
		'Q': ' ',
		'R': ' ',
		'S': ' ',
		'T': 'T',
		'U': 'U',
		'V': 'V',
		'W': 'W',
		'X': 'X',
		'Y': 'Y',
		'Z': ' ',
	}

	for t, _ = fmt.Fscan(reader, &name); t == 1; t, _ = fmt.Fscan(reader, &name) {
		result = true

		for i := 0; i < len(name); i++ {
			if reflection[rune(name[i])] != rune(name[len(name)-1-i]) {
				result = false
				break
			}
		}

		if result {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}

	writer.Flush()
}
