// https://codeforces.com/problemset/problem/1714/E

// 1714E - Add Modulo 10

// You are given an array of n integers a1,a2,…,an
// You can apply the following operation an arbitrary number of times:
//  • select an index i (1≤i≤n) and replace the value of the element ai with the value ai+(aimod10), where aimod10 is the remainder of the integer dividing ai by 10.
// For a single index (value i), this operation can be applied multiple times. If the operation is applied repeatedly to the same index, then the current value of ai is taken into account each time. For example, if ai=47 then after the first operation we get ai=47+7=54, and after the second operation we get ai=54+4=58.
// Check if it is possible to make all array elements equal by applying multiple (possibly zero) operations.
// For example, you have an array [6,11].
//  • Let's apply this operation to the first element of the array. Let's replace a1=6 with a1+(a1mod10)=6+(6mod10)=6+6=12. We get the array [12,11].
//  • Then apply this operation to the second element of the array. Let's replace a2=11 with a2+(a2mod10)=11+(11mod10)=11+1=12. We get the array [12,12].
// Thus, by applying 2 operations, you can make all elements of an array equal.

// Input
// The first line contains one integer t (1≤t≤104) — the number of test cases. What follows is a description of each test case.
// The first line of each test case contains one integer n (1≤n≤2⋅105) — the size of the array.
// The second line of each test case contains n integers ai (0≤ai≤109) — array elements.
// It is guaranteed that the sum of n over all test cases does not exceed 2⋅105.

// Output
// For each test case print:
//  • YES if it is possible to make all array elements equal;
//  • NO otherwise.
// You can print YES and NO in any case (for example, the strings yEs, yes, Yes and YES will be recognized as a positive answer) .

// Example
// input
// 10
// 2
// 6 11
// 3
// 2 18 22
// 5
// 5 10 5 10 5
// 4
// 1 2 4 8
// 2
// 4 5
// 3
// 93 96 102
// 2
// 40 6
// 2
// 50 30
// 2
// 22 44
// 2
// 1 5
// output
// Yes
// No
// Yes
// Yes
// No
// Yes
// No
// No
// Yes
// No

// Note

// The first test case is clarified above.
// In the second test case, it is impossible to make all array elements equal.
// In the third test case, you need to apply this operation once to all elements equal to 5.
// In the fourth test case, you need to apply this operation to all elements until they become equal to 8.
// In the fifth test case, it is impossible to make all array elements equal.
// In the sixth test case, you need to apply this operation to all elements until they become equal to 102.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, t, n, i int
	var a []int
	var ans bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n)
			a = make([]int, n)

			for i = 0; i < n; i++ {
				fmt.Fscan(reader, &a[i])

				for a[i]%10 != 2 && a[i]%10 != 0 {
					a[i] += a[i] % 10
				}
			}

			ans = true

			if a[0]%10 == 2 {
				for i = 1; i < n && ans; i++ {
					ans = ans && (a[i]%10 == 2 && a[i]%20 == a[i-1]%20)
				}
			} else {
				for i = 1; i < n && ans; i++ {
					ans = ans && (a[i] == a[i-1])
				}
			}

			if ans {
				fmt.Fprintln(writer, "Yes")
			} else {
				fmt.Fprintln(writer, "No")
			}
		}
	}

	writer.Flush()
}
