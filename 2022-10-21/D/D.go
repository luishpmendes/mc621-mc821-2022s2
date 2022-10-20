// https://codeforces.com/problemset/problem/1737/A

// 1737A - Ela Sorting Books

// Ela loves reading a lot, just like her new co-workers in DTL! On her first day after becoming an engineer in DTL, she is challenged by a co-worker to sort a heap of books into different compartments on the shelf.

// n books must be split into k compartments on the bookshelf (n is divisible by k). Each book is represented by a lowercase Latin letter from 'a' to 'y' inclusively, which is the beginning letter in the title of the book.
// Ela must stack exactly n/k books in each compartment. After the books are stacked, for each compartment indexed from 1 to k, she takes the minimum excluded (MEX) letter of the multiset of letters formed by letters representing all books in that compartment, then combines the resulting letters into a string. The first letter of the resulting string is the MEX letter of the multiset of letters formed by the first compartment, the second letter of the resulting string is the MEX letter of the multiset of letters formed by the second compartment, ... and so on. Please note, under the constraint of this problem, MEX letter can always be determined for any multiset found in this problem because 'z' is not used.
// What is the lexicographically greatest resulting string possible that Ela can create?
// A string a is lexicographically greater than a string b if and only if one of the following holds:
//  • b is a prefix of a, but b≠a;
//  • in the first position where a and b differ, the string a has a letter that appears later in the alphabet than the corresponding letter in b.
// The minimum excluded (MEX) letter of a multiset of letters is the letter that appears earliest in the alphabet and is not contained in the multiset. For example, if a multiset of letters contains 7 letters 'b', 'a', 'b', 'c', 'e', 'c', 'f' respectively, then the MEX letter of this compartment is 'd', because 'd' is not included in the multiset, and all letters comes before 'd' in the alphabet, namely 'a', 'b' and 'c', are included in the multiset.

// Input
// Each test contains multiple test cases. The first line contains the number of test cases t (1≤t≤100). Description of the test cases follows.
// The first line of each test case contains two integers n and k (1≤n≤200; 1≤k≤n). It is guaranteed that n is divisible by k.
// The second line of each test case contains a string of n lowercase Latin letters from 'a' to 'y' inclusively. Each letter represents the starting letter of the title of a book in the initial heap.
// It is guaranteed that the sum of n over all test cases does not exceed 1000.

// Output
// For each test case, output a string of k letters which is the most optimal string that Ela can find.

// Example
// input
// 5
// 12 3
// cabccadabaac
// 12 6
// cabccadabaac
// 12 12
// cabccadabaac
// 25 1
// abcdefghijklmnopqrstuvwxy
// 10 5
// bcdxedbcfg
// output
// edb
// ccbbba
// bbbbbaaaaaaa
// z
// aaaaa

// Note

// In the first test case, the books can be divided into 3 compartments as below:
//  • the first compartment contains the books with indices 1,2,3,7: multiset1={'c', 'a', 'b', 'd'} → MEX(multiset1)= 'e'
//  • the second compartment contains the books with indices 4,5,6,9 : multiset2={'c', 'c', 'a', 'b'} → MEX(multiset2)= 'd'
//  • the third compartment contains the remaining books 8,10,11,12 : multiset3={ 'a', 'a', 'a', 'c'} → MEX(multiset3)= 'b'
// Therefore, the answer is 'edb'. It can be proven that there is no way that Ela can arrange the books so that it results in a lexicographically greater string.

// Tutorial

// We'll iterate through compartments from 1 to K. we'll try to put 1 'a' book in it, so the MEX in that compartment will not be 'a'. If any compartment can't be filled by 'a' because we ran out of 'a', the MEX of that compartment will have to be 'a'.
// The same logic applies to 'b', 'c', 'd', ..., where for each compartment from 1, if the MEX of that compartment hasn't been determined, we'll fill the corresponding letter there.
// We'll stop when reaching the end of the alphabet, or we've reached the (N/K)-th letter in the alphabet (because the first compartment is already full).

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var q int
	var t, n, k, i, j uint
	var s string
	var count_char []uint
	var ans []byte
	var c rune
	var char byte

	for q, _ = fmt.Fscan(reader, &t); q == 1; q, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &k, &s)

			count_char = make([]uint, 26)
			ans = make([]byte, 0)

			for _, c = range s {
				count_char[c-'a']++
			}

			for i = 0; i < 25 && i*k < n; i++ {
				for k > count_char[i]+uint(len(ans)) {
					ans = append(ans, byte('a'+i))
				}
			}

			if n < 25*k {
				char = byte('a' + n/k)
			} else {
				char = byte('a' + 25)
			}

			for uint(len(ans)) < k {
				ans = append(ans, char)
			}

			for i, j = 0, uint(len(ans)); i+1 < j; i, j = i+1, j-1 {
				ans[i], ans[j-1] = ans[j-1], ans[i]
			}

			fmt.Fprintln(writer, string(ans))
		}
	}

	writer.Flush()
}
