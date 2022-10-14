// https://onlinejudge.org/index.php?option=onlinejudge&Itemid=8&category=226&page=show_problem&problem=2988

// 11888 - Abnormal 89's
// A palindrome is a word that can be read the same way in either direction. More formally if a string is
// d (d > 0) characters length and the i-th character is ai, the string is palindrome if and only if ai equals
// a(d−i+1) for 1 ≤ i ≤ d. For example “abcba” is palindrome while “aaab” is not.
// It is known that everyone who gets to know palindromes, begin an emotional relationship with
// these beautiful strings. The harmony between the letters makes them artistic. But the 89’s (those
// who entered AUT at 1389) claim they love another kind of strings. It is called alindrome. Actually
// an alindrome is the result of concatenation of two palindromes. For example “abacc”=“aba”+“cc” is
// alindrome.
// Now you should write a program to distinguish alindromes, palindromes and other kind of strings.

// Input
// The first line contains T (T ≤ 50), the number of tests. Each test that comes in a separate line contains
// a string to be checked. Input strings contain only lower case letters (‘a’ to ‘z’) and their length will be
// at most 200000.

// Output
// For each test output a single word in a single line. If the input string can be made by concatenating
// two palindromes, output ‘alindrome’. Otherwise if it’s a palindrome output ‘palindrome’. In any other
// case print ‘simple’. (Quotes for clarity)

// Sample Input
// 4
// aaa
// aabaa
// aabaaa
// abc

// Sample Output
// alindrome
// palindrome
// alindrome
// simple

// to check ‘alindrome’, find reverse of s in s + s

#include <algorithm>
#include <iostream>
#include <string>

bool is_palindrome(const std::string &s, unsigned l, unsigned r) {
    bool result = true;

    while (l < r && result) {
        if (s[l++] != s[r--]) {
            result = false;
        }
    }

    return result;
}

bool is_alindrome(const std::string &s) {
    bool result = false;

    for (unsigned i = 0; i < s.size() - 1 && !result; i++) {
        if (is_palindrome(s, 0, i) && is_palindrome(s, i + 1, s.size() - 1)) {
            result = true;
        }
    }

    return result;
}

int main() {
    unsigned t;
    std::string s, s2, r;

    while (std::cin >> t) {
        while (t--) {
            std::cin >> s;
            
            if (is_alindrome(s)) {
                std::cout << "alindrome" << std::endl;
            } else if (is_palindrome(s, 0, s.size() - 1)) {
                std::cout << "palindrome" << std::endl;
            } else {
                std::cout << "simple" << std::endl;
            }
        }
    }

	return 0;
}
