// https://codeforces.com/problemset/problem/1425/H

// 1425H - Huge Boxes of Animal Toys

// Chaneka has a hobby of playing with animal toys. Every toy has a different fun value, a real number. Chaneka has four boxes to store the toys with specification:
//  • The first box stores toys with fun values in range of (−∞,−1].
//  • The second box stores toys with fun values in range of (−1,0).
//  • The third box stores toys with fun values in range of (0,1).
//  • The fourth box stores toys with fun value in range of [1,∞).
// Chaneka has A, B, C, D toys in the first, second, third, and fourth box, respectively. One day she decides that she only wants one toy, a super toy. So she begins to create this super toy by sewing all the toys she has.
// While the number of toys Chaneka has is more than 1, she takes two different toys randomly and then sews them together, creating a new toy. The fun value of this new toy is equal to the multiplication of fun values of the sewn toys. She then puts this new toy in the appropriate box. She repeats this process until she only has one toy. This last toy is the super toy, and the box that stores this toy is the special box.
// As an observer, you only know the number of toys in each box initially but do not know their fun values. You also don't see the sequence of Chaneka's sewing. Determine which boxes can be the special box after Chaneka found her super toy.

// Input
// The first line has an integer T (1≤T≤5⋅104), the number of test cases.
// Every case contains a line with four space-separated integers A B C D (0≤A,B,C,D≤10^6,A+B+C+D>0), which denotes the number of toys in the first, second, third, and fourth box, respectively.

// Output
// For each case, print four space-separated strings. Each string represents the possibility that the first, second, third, and fourth box can be the special box from left to right.
// For each box, print "Ya" (Without quotes, Indonesian for yes) if that box can be the special box. Print "Tidak" (Without quotes, Indonesian for No) otherwise.

// Example
// input
// 2
// 1 2 0 1
// 0 1 0 0
// output
// Ya Ya Tidak Tidak
// Tidak Ya Tidak Tidak

// Note

// For the first case, here is a scenario where the first box is the special box:
//  • The first box had toys with fun values {−3}.
//  • The second box had toys with fun values {−0.5,−0.5}
//  • The fourth box had toys with fun values {3}
// The sewing sequence:
//  1. Chaneka sews the toy with fun −0.5 and −0.5 to a toy with fun 0.25 and then put it in the third box.
//  2. Chaneka sews the toy with fun −3 and 0.25 to a toy with fun −0.75 and then put it in the second box.
//  3. Chaneka sews the toy with fun −0.75 and 3 to a toy with fun −1.25 and then put it in the first box, which then became the special box.
// Here is a scenario where the second box ends up being the special box:
//  • The first box had toys with fun values {−3}
//  • The second box had toys with fun values {−0.33,−0.25}.
//  • The fourth box had toys with fun values {3}.
// The sewing sequence:
//  1. Chaneka sews the toy with fun −3 and −0.33 to a toy with fun 0.99 and then put it in the third box.
//  2. Chaneka sews the toy with fun 0.99 and 3 to a toy with fun 2.97 and then put in it the fourth box.
//  3. Chaneka sews the toy with fun 2.97 and −0.25 to a toy with fun −0.7425 and then put it in the second box, which then became the special box.
// There is only one toy for the second case, so Chaneka does not have to sew anything because that toy, by definition, is the super toy.

// Tutorial

// The main observation of this problem is that the sewing order does not matter. The final fun value is the multiplication of fun values of all toys. There are two separate cases we must consider:
//  • Cek the sign of the toy. If (A+B) is even, the sign is positive. Else, it's negative.
//  • Cek whether it is possible to make a fun value with absolute value lower than 1 and greater than 1. The former is possible if (B+C)>0, the latter is possible if (A+D)>0.
// Combining these two cases will give you the solution.
// Time Complexity: O(T)

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var s int
	var t, a, b, c, d uint

	for s, _ = fmt.Fscan(reader, &t); s == 1; s, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &a, &b, &c, &d)

			if a == 0 && d == 0 {
				if b&1 == 1 {
					fmt.Fprintln(writer, "Tidak Ya Tidak Tidak")
				} else {
					fmt.Fprintln(writer, "Tidak Tidak Ya Tidak")
				}
			} else if b == 0 && c == 0 {
				if a&1 == 1 {
					fmt.Fprintln(writer, "Ya Tidak Tidak Tidak")
				} else {
					fmt.Fprintln(writer, "Tidak Tidak Tidak Ya")
				}
			} else {
				if (a+b)&1 == 1 {
					fmt.Fprintln(writer, "Ya Ya Tidak Tidak")
				} else {
					fmt.Fprintln(writer, "Tidak Tidak Ya Ya")
				}
			}
		}
	}

	writer.Flush()
}
