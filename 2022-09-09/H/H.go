// https://codeforces.com/problemset/problem/1166/A

// 1166A - Silent Classroom

// There are n students in the first grade of Nlogonia high school. The principal wishes to split the students into two classrooms (each student must be in exactly one of the classrooms). Two distinct students whose name starts with the same letter will be chatty if they are put in the same classroom (because they must have a lot in common). Let x be the number of such pairs of students in a split. Pairs (a,b) and (b,a) are the same and counted only once.
// For example, if there are 6 students: "olivia", "jacob", "tanya", "jack", "oliver" and "jessica", then:
//  • splitting into two classrooms ("jack", "jacob", "jessica", "tanya") and ("olivia", "oliver") will give x=4 (3 chatting pairs in the first classroom, 1 chatting pair in the second classroom),
//  • splitting into two classrooms ("jack", "tanya", "olivia") and ("jessica", "oliver", "jacob") will give x=1 (0 chatting pairs in the first classroom, 1 chatting pair in the second classroom).
// You are given the list of the n names. What is the minimum x we can obtain by splitting the students into classrooms?
// Note that it is valid to place all of the students in one of the classrooms, leaving the other one empty.

// Input
// The first line contains a single integer n (1≤n≤100) — the number of students.
// After this n lines follow.
// The i-th line contains the name of the i-th student.
// It is guaranteed each name is a string of lowercase English letters of length at most 20. Note that multiple students may share the same name.

// Output
// The output must consist of a single integer x — the minimum possible number of chatty pairs.

// Examples

// input
// 4
// jorge
// jose
// oscar
// jerry
// output
// 1

// input
// 7
// kambei
// gorobei
// shichiroji
// kyuzo
// heihachi
// katsushiro
// kikuchiyo
// output
// 2

// input
// 7
// 5
// mike
// mike
// mike
// mike
// mike
// output
// 4

// Note

// In the first sample the minimum number of pairs is 1. This can be achieved, for example, by putting everyone except jose in one classroom, and jose in the other, so jorge and jerry form the only chatty pair.
// In the second sample the minimum number of pairs is 2. This can be achieved, for example, by putting kambei, gorobei, shichiroji and kyuzo in one room and putting heihachi, katsushiro and kikuchiyo in the other room. In this case the two pairs are kambei and kyuzo, and katsushiro and kikuchiyo.
// In the third sample the minimum number of pairs is 4. This can be achieved by placing three of the students named mike in one classroom and the other two students in another classroom. Thus there will be three chatty pairs in one classroom and one chatty pair in the other classroom.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, n, i, ans int
	var s string
	var v []int

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		v = make([]int, 'z'-'a'+1)

		for i = 0; i < n; i++ {
			fmt.Fscan(reader, &s)
			v[s[0]-'a']++
		}

		ans = 0

		for i = 0; i < len(v); i++ {
			if (v[i] & 1) == 0 {
				ans += (v[i] * (v[i] - 2)) / 4
			} else {
				ans += ((v[i] - 1) * (v[i] - 1)) / 4
			}
		}

		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}
