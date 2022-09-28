// https://codeforces.com/problemset/problem/525/A

// 525A - Vitaliy and Pie

// After a hard day Vitaly got very hungry and he wants to eat his favorite potato pie. But it's not that simple. Vitaly is in the first room of the house with n room located in a line and numbered starting from one from left to right. You can go from the first room to the second room, from the second room to the third room and so on — you can go from the (n - 1)-th room to the n-th room. Thus, you can go to room x only from room x - 1.
// The potato pie is located in the n-th room and Vitaly needs to go there.
// Each pair of consecutive rooms has a door between them. In order to go to room x from room x - 1, you need to open the door between the rooms with the corresponding key.
// In total the house has several types of doors (represented by uppercase Latin letters) and several types of keys (represented by lowercase Latin letters). The key of type t can open the door of type T if and only if t and T are the same letter, written in different cases. For example, key f can open door F.
// Each of the first n - 1 rooms contains exactly one key of some type that Vitaly can use to get to next rooms. Once the door is open with some key, Vitaly won't get the key from the keyhole but he will immediately run into the next room. In other words, each key can open no more than one door.
// Vitaly realizes that he may end up in some room without the key that opens the door to the next room. Before the start his run for the potato pie Vitaly can buy any number of keys of any type that is guaranteed to get to room n.
// Given the plan of the house, Vitaly wants to know what is the minimum number of keys he needs to buy to surely get to the room n, which has a delicious potato pie. Write a program that will help Vitaly find out this number.

// Input
// The first line of the input contains a positive integer n (2 ≤ n ≤ 105) — the number of rooms in the house.
// The second line of the input contains string s of length 2·n - 2. Let's number the elements of the string from left to right, starting from one.
// The odd positions in the given string s contain lowercase Latin letters — the types of the keys that lie in the corresponding rooms. Thus, each odd position i of the given string s contains a lowercase Latin letter — the type of the key that lies in room number (i + 1) / 2.
// The even positions in the given string contain uppercase Latin letters — the types of doors between the rooms. Thus, each even position i of the given string s contains an uppercase letter — the type of the door that leads from room i / 2 to room i / 2 + 1.

// Output
// Print the only integer — the minimum number of keys that Vitaly needs to buy to surely get from room one to room n.

// Examples

// input
// 3
// aAbB
// output
// 0

// input
// 4
// aBaCaB
// output
// 3

// input
// 5
// xYyXzZaZ
// output
// 2

// Tutorial
// To solve this problem we need to use array cnt[]. In this array we will store number of keys of every type, which we already found in rooms, but didn't use. Answer will store in variable ans.
// Now, we iterate on string. If current element of string si is lowercase letter (key), we make cnt[si]++. Else if current element of string si uppercase letter (door) and cnt[tolower(si)] > 0, we make cnt[tolower(si)]--, else we make ans++. It remains only to print ans.
// Asymptotic behavior of this solution — O(|s|), where |s| — length of string s.

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
	var n, ans, i uint
	var s string
	var cnt [26]uint

	for t, _ = fmt.Fscan(reader, &n); t == 1; t, _ = fmt.Fscan(reader, &n) {
		fmt.Fscan(reader, &s)
		ans = 0

		for i = 0; i < 26; i++ {
			cnt[i] = 0
		}

		for i = 0; i < 2*n-2; i++ {
			if i&1 == 0 {
				cnt[s[i]-'a']++
			} else if cnt[s[i]-'A'] > 0 {
				cnt[s[i]-'A']--
			} else {
				ans++
			}
		}

		fmt.Fprintln(writer, ans)
	}

	writer.Flush()
}