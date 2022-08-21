// https://codeforces.com/problemset/problem/117/A

// 117A - Elevator

// And now the numerous qualifying tournaments for one of the most prestigious Russian contests Russian Codec Cup are over. All n participants who have made it to the finals found themselves in a huge m-floored 108-star hotel. Of course the first thought to come in a place like this is "How about checking out the elevator?".
// The hotel's elevator moves between floors according to one never changing scheme. Initially (at the moment of time 0) the elevator is located on the 1-st floor, then it moves to the 2-nd floor, then — to the 3-rd floor and so on until it reaches the m-th floor. After that the elevator moves to floor m - 1, then to floor m - 2, and so on until it reaches the first floor. This process is repeated infinitely. We know that the elevator has infinite capacity; we also know that on every floor people get on the elevator immediately. Moving between the floors takes a unit of time.
// For each of the n participant you are given si, which represents the floor where the i-th participant starts, fi, which represents the floor the i-th participant wants to reach, and ti, which represents the time when the i-th participant starts on the floor si.
// For each participant print the minimum time of his/her arrival to the floor fi.
// If the elevator stops on the floor si at the time ti, then the i-th participant can enter the elevator immediately. If the participant starts on the floor si and that's the floor he wanted to reach initially (si = fi), then the time of arrival to the floor fi for this participant is considered equal to ti.

// Input
// The first line contains two space-separated integers n and m (1 ≤ n ≤ 105, 2 ≤ m ≤ 108).
// Next n lines contain information about the participants in the form of three space-separated integers si fi ti (1 ≤ si, fi ≤ m, 0 ≤ ti ≤ 108), described in the problem statement.

// Output
// Print n lines each containing one integer — the time of the arrival for each participant to the required floor.

// Examples

// input
// 7 4
// 2 4 3
// 1 2 0
// 2 2 0
// 1 2 1
// 4 3 5
// 1 2 2
// 4 2 0
// output
// 9
// 1
// 0
// 7
// 10
// 7
// 5

// input
// 5 5
// 1 5 4
// 1 3 1
// 1 3 4
// 3 1 5
// 4 2 5
// output
// 12
// 10
// 10
// 8
// 7

// Note

// Let's consider the first sample. The first participant starts at floor s = 2 by the time equal to t = 3. To get to the floor f = 4, he has to wait until the time equals 7, that's the time when the elevator will go upwards for the second time. Then the first participant should get on the elevator and go two floors up. In this case the first participant gets to the floor f at time equal to 9. The second participant starts at the time t = 0 on the floor s = 1, enters the elevator immediately, and arrives to the floor f = 2. The third participant doesn't wait for the elevator, because he needs to arrive to the same floor where he starts.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var c, n, m, i, s, f, t, d int

	for c, _ = fmt.Fscan(reader, &n, &m); c == 2; c, _ = fmt.Fscan(reader, &n, &m) {
		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &s, &f, &t)

			if s == f {
				fmt.Fprintln(writer, t)
			} else {
				d = t / (2*m - 2)
				t %= (2*m - 2)

				if t > 2*m-1-s {
					t -= 2*m - 2
					d++
				}

				if t > s-1 {
					t = 2*m - 1 - s
				} else { // t <= s-1
					t = s - 1
				}

				if t > 2*m-1-f {
					t = 2*m - 2 + f - 1
				} else if t > f-1 {
					t = 2*m - 1 - f
				} else { // t <= 2*m-1-f && t <= f-1
					t = f - 1
				}

				fmt.Fprintln(writer, t+d*(2*m-2))
			}
		}
	}

	writer.Flush()
}
