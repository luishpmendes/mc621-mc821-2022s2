// https://open.kattis.com/problems/doorman

// Kattis doorman - Doorman
// The doorman Bruno at the popular night club Heaven is having a hard time fulfilling his duties. He was told by the owner that when the club is full, the number of women and men let into the club should be roughly the same. When the night club opens, people wanting to enter the club are already lined up in a queue, and Bruno can only let them in one-by-one. He lets them in more-or-less in the order they are lined up. He can however decide to let the second person in the queue cut the line and slip into the club before the person in front. This will no doubt upset the person first in line, especially when this happens multiple times, but Bruno is quite a big guy and is capable of handling troublemakers.
// Unfortunately though, he is not that strong on mental calculations under these circumstances. He finds keeping track of the difference of the number of women and number of men let into the club a challenging task. As soon as the absolute difference gets too big, he loses track of his counting and must declare to the party people remaining in the queue that the club is full.

// Input
// The first line of input contains a positive integer X < 100 describing the largest absolute difference between the number of women and number of men let into the club that Bruno can handle.
// The second line contains a string consisting solely of the characters ’W’ and ’M’ of length at most 100, describing the genders of the people in the queue, in order of their arrrival. The leftmost character of the string is the gender of the person first in line.
// You may assume that the club is large enough to hold all the people in the queue.

// Output
// The maximum number of people Bruno can let into the club without losing track of his counting.

// Sample Input 1
// 1
// MWWMWMMWM
// Sample Output 1
// 9

// Sample Input 2
// 2
// WMMMMWWMMMWWMW
// Sample Output 2
// 8

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, i int
	var x, count_m, count_w uint
	var s string
	var chars []byte

	for t, _ = fmt.Fscan(reader, &x, &s); t == 2; t, _ = fmt.Fscan(reader, &x, &s) {
		count_m = 0
		count_w = 0
		chars = []byte(s)

		for i = 0; i < len(chars) && count_m <= count_w+x && count_w <= count_m+x; i++ {
			if i+1 < len(chars) && ((chars[i] == 'M' && count_m == count_w+x) || (chars[i] == 'W' && count_w == count_m+x)) {
				chars[i], chars[i+1] = chars[i+1], chars[i]
			}

			if chars[i] == 'M' {
				count_m++
			} else {
				count_w++
			}
		}

		if count_m > count_w+x || count_w > count_m+x {
			i--
		}

		fmt.Fprintln(writer, i)
	}

	writer.Flush()
}
