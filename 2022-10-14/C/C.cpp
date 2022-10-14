// https://onlinejudge.org/index.php?option=onlinejudge&Itemid=8&category=26&page=show_problem&problem=2470

// 11475 - Extend to Palindrome
// Your task is, given an integer N, to make a palidrome (word that reads the same when you reverse
// it) of length at least N. Any palindrome will do. Easy, isn’t it? That’s what you thought before you
// passed it on to your inexperienced team-mate. When the contest is almost over, you find out that
// that problem still isn’t solved. The problem with the code is that the strings generated are often not
// palindromic. There’s not enough time to start again from scratch or to debug his messy code. Seeing
// that the situation is desperate, you decide to simply write some additional code that takes the output
// and adds just enough extra characters to it to make it a palindrome and hope for the best. Your
// solution should take as its input a string and produce the smallest palindrome that can be formed by
// adding zero or more characters at its end.

// Input
// Input will consist of several lines ending in EOF. Each line will contain a non-empty string made up of
// upper case and lower case English letters (‘A’-‘Z’ and ‘a’-‘z’). The length of the string will be less than
// or equal to 100,000.

// Output
// For each line of input, output will consist of exactly one line. It should contain the palindrome formed
// by adding the fewest number of extra letters to the end of the corresponding input string.

// Sample Input
// aaaa
// abba
// amanaplanacanal
// xyz

// Sample Output
// aaaa
// abba
// amanaplanacanalpanama
// xyzyx

// ‘border’ of KMP

#include <algorithm>
#include <iostream>
#include <string>
#include <vector>

std::vector<int> kmp_preprocess(const std::string & p) {
    int i, j;
    std::vector<int> b (p.size(), 0);

    i = 1;
    j = -1;
    b[0] = -1;

    while (i < (int) p.size()) {
        while (j >= 0 && p[i] != p[j+1]) {
            j = b[j];
        }

        if (p[j+1] == p[i]) {
            j++;
        }
        
        b[i] = j;
        i++;
    }

    return b;
}

int kmp_search(const std::string & t, const std::string & p) {
    int i, j;
    std::vector<int> b;

    b = kmp_preprocess(p);

    i = 0;
    j = -1;

    while (i < (int) t.size()) {
        while (j >= 0 && t[i] != p[j+1]) {
            j = b[j]; 
        }

        if (t[i] == p[j+1]) {
            j++;
        }

        i++;
    }

    return j;
}

int main() {
    std::string s, r;
    int i;

    while (std::cin >> s) {
        r = s;
        std::reverse(r.begin(), r.end());
        i = kmp_search(s, r);
        s = r.substr(i+1);
        std::reverse(s.begin(), s.end());
        std::cout << s << r << std::endl;
    }

	return 0;
}
