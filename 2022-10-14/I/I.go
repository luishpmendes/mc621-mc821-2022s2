// https://codeforces.com/problemset/problem/1428/B

// 1428B - Belted Rooms

// In the snake exhibition, there are n rooms (numbered 0 to n−1) arranged in a circle, with a snake in each room. The rooms are connected by n conveyor belts, and the i-th conveyor belt connects the rooms i and (i+1)modn. In the other words, rooms 0 and 1, 1 and 2, …, n−2 and n−1, n−1 and 0 are connected with conveyor belts.
// The i-th conveyor belt is in one of three states:
//  • If it is clockwise, snakes can only go from room i to (i+1)modn.
//  • If it is anticlockwise, snakes can only go from room (i+1)modn to i.
//  • If it is off, snakes can travel in either direction.
// Above is an example with 4 rooms, where belts 0 and 3 are off, 1 is clockwise, and 2 is anticlockwise.
// Each snake wants to leave its room and come back to it later. A room is returnable if the snake there can leave the room, and later come back to it using the conveyor belts. How many such returnable rooms are there?

// Input
// Each test contains multiple test cases. The first line contains a single integer t (1≤t≤1000): the number of test cases. The description of the test cases follows.
// The first line of each test case description contains a single integer n (2≤n≤300000): the number of rooms.
// The next line of each test case description contains a string s of length n, consisting of only '<', '>' and '-'.
//  • If si= '>', the i-th conveyor belt goes clockwise.
//  • If si= '<', the i-th conveyor belt goes anticlockwise.
//  • If si= '-', the i-th conveyor belt is off.
// It is guaranteed that the sum of n among all test cases does not exceed 300000.

// Output
// For each test case, output the number of returnable rooms.

// Example
// input
// 4
// 4
// -><-
// 5
// >>>>>
// 3
// <--
// 2
// <>
// output
// 3
// 5
// 3
// 0

// Note

// In the first test case, all rooms are returnable except room 2. The snake in the room 2 is trapped and cannot exit. This test case corresponds to the picture from the problem statement.
// In the second test case, all rooms are returnable by traveling on the series of clockwise belts.

// Tutorial

// Hint 1
// There are 2 cases to consider for a room to be returnable.

// Hint 2
// For a room to be returnable, either go one big round around all the rooms or move to an adjacent room and move back.

// Solution
// Let's consider two ways to return to the start point. The first is to go one big round around the circle. The second is to move 1 step to the side, and return back immediately.
// Going one big round is only possible if and only if:
//  • There are no clockwise belts OR
//  • There are no anticlockwise belts
// If we can go one big round, all rooms are returnable.
// If there are both clockwise and anticlockwise belts, then we can't go one big round. For any room to be returnable, it must have an off belt to the left or to the right.
// In summary, check if clockwise belts are absent or if anticlockwise belts are absent. If either is absent, the answer is n. Otherwise, we have to count the number of rooms with an off belt to the left or to the right.

// Other comments
// Sorry for the unclear statement for B, we should've explained each sample testcase more clearly with better diagrams. Additionally, we're also sorry for the weak pretests. We should've added more testcases of smaller length, and thanks to hackers for adding stronger tests.

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
	var t, n, i, ans uint
	var s string
	var hasCW, hasCCW bool

	for c, _ = fmt.Fscan(reader, &t); c == 1; c, _ = fmt.Fscan(reader, &t) {
		for ; t > 0; t-- {
			fmt.Fscan(reader, &n, &s)
			hasCW, hasCCW = false, false

			for i = 0; i < n; i++ {
				if s[i] == '<' {
					hasCCW = true
				} else if s[i] == '>' {
					hasCW = true
				}
			}

			if !hasCW || !hasCCW {
				ans = n
			} else { // hasCW && hasCCW
				ans = 0
				s += s[0:1]
				for i = 0; i < n; i++ {
					if s[i] == '-' || s[(i+1)%n] == '-' {
						ans++
					}
				}
			}

			fmt.Fprintln(writer, ans)
		}
	}

	writer.Flush()
}
