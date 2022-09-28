// https://codeforces.com/problemset/problem/1553/A

// 1553A - Digits Sum

// Let's define S(x) to be the sum of digits of number x written in decimal system. For example, S(5)=5, S(10)=1, S(322)=7.
// We will call an integer x interesting if S(x+1)<S(x). In each test you will be given one integer n. Your task is to calculate the number of integers x such that 1≤x≤n and x is interesting.

// Input
// The first line contains one integer t (1≤t≤1000)  — number of test cases.
// Then t lines follow, the i-th line contains one integer n (1≤n≤109) for the i-th test case.

// Output
// Print t integers, the i-th should be the answer for the i-th test case.

// Example
// input
// 5
// 1
// 9
// 10
// 34
// 880055535
// output
// 0
// 1
// 1
// 3
// 88005553

// Note

// The first interesting number is equal to 9.

// Tutorial

// Let's think: what properties do all interesting numbers have? Well, if a number x does not end with 9, we can say for sure that f(x+1)=f(x)+1, because the last digit will get increased. What if the number ends with 9? Then the last digit will become 0, so, no matter what happens to other digits, we can say that f(x+1) will surely be less than f(x).
// So the problem asks us to count all numbers 1≤x≤n with the last digit equal to 9. It is not hard to see that the answer is equal to ⌊n+1/10⌋.
// This concludes the solution, as we are now able to answer all testcases in O(1), resulting in total O(t) runtime.

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
	var t, n uint

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			fmt.Fprintln(writer, (n+1)/10)
		}
	}

	writer.Flush()
}
