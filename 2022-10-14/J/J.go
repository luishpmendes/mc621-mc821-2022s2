// https://codeforces.com/problemset/problem/1679/A

// 1679A - AvtoBus

// Spring has come, and the management of the AvtoBus bus fleet has given the order to replace winter tires with summer tires on all buses.
// You own a small bus service business and you have just received an order to replace n tires. You know that the bus fleet owns two types of buses: with two axles (these buses have 4 wheels) and with three axles (these buses have 6 wheels).
// You don't know how many buses of which type the AvtoBus bus fleet owns, so you wonder how many buses the fleet might have. You have to determine the minimum and the maximum number of buses that can be in the fleet if you know that the total number of wheels for all buses is n.

// Input
// The first line contains an integer t (1≤t≤1000) — the number of test cases. The following lines contain description of test cases.
// The only line of each test case contains one integer n (1≤n≤10^18) — the total number of wheels for all buses.

// Output
// For each test case print the answer in a single line using the following format.
// Print two integers x and y (1≤x≤y) — the minimum and the maximum possible number of buses that can be in the bus fleet.
// If there is no suitable number of buses for the given n, print the number −1 as the answer.

// Example
// input
// 4
// 4
// 7
// 24
// 998244353998244352
// output
// 1 1
// -1
// 4 6
// 166374058999707392 249561088499561088

// Note

// In the first test case the total number of wheels is 4. It means that there is the only one bus with two axles in the bus fleet.
// In the second test case it's easy to show that there is no suitable number of buses with 7 wheels in total.
// In the third test case the total number of wheels is 24. The following options are possible:
//  • Four buses with three axles.
//  • Three buses with two axles and two buses with three axles.
//  • Six buses with two axles.
// So the minimum number of buses is 4 and the maximum number of buses is 6.

// Tutorial

// Let the number of buses with two axles is x and the number of buses with three axles is y. Then the equality 4x+6y=n must be true. If n is odd, there is no answer, because the left part of the equality is always even. Now we can divide each part of the equality by two: 2x+3y=n/2.
// Let's maximize the number of buses. Then we should make x as large as possible. So, we will get 2+…+2+2=n2 if n2 is even, and 2+…+2+3=n/2 otherwise. In both cases the number of buses is ⌊n/2⌋.
// Now let's minimize the number of buses. So, we should make y as large is possible. We will get 3+…+3+3+3=n/2 if n/2 is divisible by 3, 3+…+3+3+2=n/2 if n≡2(mod 3), and 3+…+3+2+2=n/2 if n≡1(mod 3). In all cases the number of buses is ⌈n/3⌉.
// Also don't forget the case n=2 — each bus has at least four wheels, so in this case there is no answer.
// Time complexity: O(1).

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
	var n uint64

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			if n&1 == 1 || n < 4 {
				fmt.Fprintln(writer, -1)
			} else {
				fmt.Fprintln(writer, (n+5)/6, n/4)
			}
		}
	}

	writer.Flush()
}
