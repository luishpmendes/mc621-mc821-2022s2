// https://codeforces.com/problemset/problem/1734/E

// 1734E - Rectangular Congruence

// You are given a prime number n, and an array of n integers b1,b2,…,bn, where 0≤bi<n for each 1≤i≤n.
// You have to find a matrix a of size n×n such that all of the following requirements hold:
//  • 0≤ai,j<n for all 1≤i,j≤n.
//  • ar1,c1+ar2,c2≢ar1,c2+ar2,c1(modn) for all positive integers r1, r2, c1, and c2 such that 1≤r1<r2≤n and 1≤c1<c2≤n.
//  • ai,i=bi for all 1≤i≤n.
// Here x≢y(modm) denotes that x and y give different remainders when divided by m.
// If there are multiple solutions, output any. It can be shown that such a matrix always exists under the given constraints.

// Input
// The first line contains a single positive integer n (2≤n<350).
// The second line contains n integers b1,b2,…,bn (0≤bi<n) — the required values on the main diagonal of the matrix.
// It is guaranteed that n is prime.

// Output
// Print n lines. On the i-th line, print n integers ai,1,ai,2,…,ai,n, each separated with a space.
// If there are multiple solutions, output any.

// Examples

// input
// 2
// 0 0
// output
// 0 1
// 0 0

// input
// 3
// 1 1 1
// output
// 1 2 2
// 1 1 0
// 1 0 1

// input
// 5
// 1 4 1 2 4
// output
// 1 0 1 3 4
// 1 4 3 1 0
// 2 4 1 0 2
// 1 2 2 2 2
// 2 2 0 1 4

// Note

// In the first example, the answer is valid because all entries are non-negative integers less than n=2, and a1,1+a2,2≢a1,2+a2,1(mod2) (because a1,1+a2,2=0+0≡0(mod2) and a1,2+a2,1=1+0≡1(mod2)). Moreover, the values on the main diagonals are equal to 0,0 as required.
// In the second example, the answer is correct because all entries are non-negative integers less than n=3, and the second condition is satisfied for all quadruplets (r1,r2,c1,c2). For example:
//  • When r1=1, r2=2, c1=1 and c2=2, a1,1+a2,2≢a1,2+a2,1(mod3) because a1,1+a2,2=1+1≡2(mod3) and a1,2+a2,1=2+1≡0(mod3).
//  • When r1=2, r2=3, c1=1, and c2=3, a2,1+a3,3≢a2,3+a3,1(mod3) because a2,1+a3,3=1+1≡2(mod3) and a2,3+a3,1=0+1≡1(mod3).
// Moreover, the values on the main diagonal are equal to 1,1,1 as required.

// Tutorial

// We say a matrix to be good if it satisfies the congruence condition (the second condition).
// When we have a good matrix, we can add any value c to a whole row while maintaining the congruence relation. The same is true for adding the same value to a whole column.
// Suppose we have any good matrix A, then by adding bi−ai,i to the i-th row for each of i=1,2,⋯,n, we obtain a good matrix that has the desired values on the diagonal.
// In fact, there are a lot of possible constructions. We present a few of them here:
//  • ai,j=i×j(modn)
//  • ai,j=(i+j)2(modn). This needs special handling when n=2.
//  • ai,j=(i+j)(i+j+1)2(modn).
//  • The coolest part is that all quadratic polynomials in the form ai2+bij+cj2+di+ej+f are valid for all integers a,b,c,d,e,f and b≢0(modn)
// As a bonus, we prove that the general quadratic polynomial gives a good construction.

// Proof
// Since we can add values to a whole row or column, and i2,j2,i,j,1 are also constant on the rows and columns, adding them by ai,j has no effect. So we may just assume a=c=d=e=f=0.
// So it suffices to show that ai,j=bij(modn) satisfies the condition. We can see directly that ar1,c1−ar2,c1−ar1,c2+ar2,c2=b(r1−r2)(c1−c2). As r1≠r2, c1≠c2, b≢0(modn), and n is a prime, this expression must be nonzero (modn).

// Here are some extra observations that may enable one to find a good matrix more quickly.

// Some extra observations

// For each j, the values ai,j−ai−1,j over all 0≤i<n must be a permutation of 0,1,⋯,n−1.
// A good matrix stays good if we permute the rows or permute the columns. Therefore, we can show that there exists some good matrix with a1,j=ai,1=0, and a2,j=j−1, ai,2=i−1. Assuming this, it should not be difficult to discover ai,j=(i−1)(j−1) yields one of the good matrices.
// Let b be the two dimensional difference array of a, that is, bi,j=ai,j−ai−1,j−ai,j−1+ai−1,j−1. Then, the condition becomes sum of any rectangles of b must be nonzero (modn). It is easy to see b=1 is valid. This corresponds to the solution that ai,j=i×j(modn).

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
	var n, i, j, aux uint
	var b []uint

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		b = make([]uint, n)

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &b[i])
			aux = (n + b[i] - ((i * i) % n)) % n

			for j = 0; j < n; j++ {
				fmt.Fprint(writer, (((i*j)%n)+aux)%n)

				if j < n-1 {
					fmt.Fprint(writer, " ")
				}
			}

			fmt.Fprintln(writer)
		}
	}

	writer.Flush()
}
