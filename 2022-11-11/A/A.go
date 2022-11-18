// https://www.spoj.com/problems/ONP/

// ONP - Transform the Expression
// Transform the algebraic expression with brackets into RPN form (Reverse Polish Notation). Two-argument operators: +, -, *, /, ^ (priority from the lowest to the highest), brackets ( ). Operands: only letters: a,b,...,z. Assume that there is only one RPN form (no expressions like a*b*c).

// Input
// t [the number of expressions <= 100]
// expression [length <= 400]
// [other expressions]
// Text grouped in [ ] does not appear in the input file.

// Output
// The expressions in RPN form, one per line.

// Example
// input
// 3
// (a+(b*c))
// ((a+b)*(z+x))
// ((a+t)*((b+(a+c))^(c+d)))
// output
// abc*+
// ab+zx+*
// at+bac++cd+^*

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
	var stack []byte
	var ans []byte

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &s)
			stack = make([]byte, 0)
			ans = make([]byte, 0)

			for _, char := range []byte(s) {
				if char == '(' {
					stack = append(stack, char)
				} else if char == ')' {
					ans = append(ans, stack[len(stack)-1])
					stack = stack[:len(stack)-2]
				} else if char >= 'a' && char <= 'z' {
					ans = append(ans, char)
				} else {
					stack = append(stack, char)
				}
			}

			fmt.Fprintln(writer, string(ans))
		}
	}

	writer.Flush()
}
