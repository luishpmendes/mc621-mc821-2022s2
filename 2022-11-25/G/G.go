// https://open.kattis.com/problems/primepath

// Kattis primepath - Prime Path
// The ministers of the cabinet were quite upset by the message from the Chief of Security stating that they would all have to change the four-digit room numbers on their offices.
// — It is a matter of security to change such things every now and then, to keep the enemy in the dark.
// — But look, I have chosen my number 1033 for good reasons. I am the Prime minister, you know!
// — I know, so therefore your new number 8179 is also a prime. You will just have to paste four new digits over the four old ones on your office door.
// — No, it’s not that simple. Suppose that I change the first digit to an 8, then the number will read 8033 which is not a prime!
// — I see, being the prime minister you cannot stand having a non-prime number on your door even for a few seconds.
// — Correct! So I must invent a scheme for going from 1033 to 8179 by a path of prime numbers where only one digit is changed from one prime to the next prime.
// Now, the minister of finance, who had been eavesdropping, intervened.
// — No unnecessary expenditure, please! I happen to know that the price of a digit is one pound.
// — Hmm, in that case I need a computer program to minimize the cost. You don’t know some very cheap software gurus, do you?
// — In fact, I do. You see, there is this programming contest going on
// ...
// Help the prime minister to find the cheapest prime path between any two given four-digit primes! The first digit must be nonzero, of course. Here is a solution in the case above.
//     1033
//     1733
//     3733
//     3739
//     3779
//     8779
//     8179
// The cost of this solution is 6 pounds. Note that the digit 1 which got pasted over in step 2 can not be reused in the last step – a new 1 must be purchased.

// Input
// One line with a positive number: the number of test cases (at most 100). Then for each test case, one line with two numbers separated by a blank. Both numbers are four-digit primes (without leading zeros).

// Output
// One line for each case, either with a number stating the minimal cost or containing the word “Impossible”.

// Sample Input 1
// 3
// 1033 8179
// 1373 8017
// 1033 1033
// Sample Output 1
// 6
// 7
// 0

package main

import (
	"bufio"
	"fmt"
	"os"
)

func sieve(n int) []bool {
	var is_prime []bool
	var i, j int

	is_prime = make([]bool, n+1)

	for i = 2; i <= n; i++ {
		is_prime[i] = true
	}

	for i = 2; i <= n; i++ {
		if is_prime[i] {
			for j = i * i; j <= n; j += i {
				is_prime[j] = false
			}
		}
	}

	return is_prime
}

func are_adjacent(a, b int) bool {
	var diff int
	diff = 0

	for a > 0 && b > 0 {
		if a%10 != b%10 {
			diff++
		}
		a /= 10
		b /= 10
	}

	return diff == 1
}

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var q, c, s, t, u, i, j int
	var is_prime []bool
	var id_to_num, queue, dist []int
	var num_to_id map[int]int
	var adj [][]int

	is_prime = sieve(9999)

	id_to_num = make([]int, 0)
	num_to_id = make(map[int]int)

	for u = 1000; u <= 9999; u++ {
		if is_prime[u] {
			num_to_id[u] = len(id_to_num)
			id_to_num = append(id_to_num, u)
		}
	}

	adj = make([][]int, len(id_to_num))

	for i = 0; i < len(id_to_num); i++ {
		for j = i + 1; j < len(id_to_num); j++ {
			if are_adjacent(id_to_num[i], id_to_num[j]) {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}

	for q, _ = fmt.Fscan(reader, &c); q == 1; q, _ = fmt.Fscan(reader, &c) {
		for ; c > 0; c-- {
			fmt.Fscan(reader, &s, &t)
			dist = make([]int, len(id_to_num))
			queue = make([]int, 0, len(id_to_num))

			for i = 0; i < len(id_to_num); i++ {
				dist[i] = len(id_to_num)
			}

			dist[num_to_id[s]] = 0
			queue = append(queue, num_to_id[s])

			for len(queue) > 0 {
				u = queue[0]
				queue = queue[1:]

				for _, v := range adj[u] {
					if dist[v] > dist[u]+1 {
						dist[v] = dist[u] + 1
						queue = append(queue, v)
					}
				}
			}

			fmt.Fprintln(writer, dist[num_to_id[t]])
		}
	}

	writer.Flush()
}
