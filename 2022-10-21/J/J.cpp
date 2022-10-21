// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=22&page=show_problem&problem=1963

// 11022 - String Factoring
// Spotting patterns in seemingly random strings is a problem with many applications. E.g. in our efforts
// to understand the genome we investigate the structure of DNA strings. In data compression we are
// interested in finding repetitions, so the data can be represented more efficiently with pointers. Another
// plausible example arises from the area of artificial intelligence, as how to interpret information given
// to you in a language you do not know. The natural thing to do in order to decode the information
// message would be to look for recurrences in it. So if the SETI project (the Search for Extra Terrestrial
// Intelligence) ever get a signal in the H21-spectra, we need to know how to decompose it.
// One way of capturing the redundancy of a string is to find its factoring. If two or more identical
// substrings A follow each other in a string S, we can represent this part of S as the substring A, embraced
// by parentheses, raised to the power of the number of its recurrences. E.g. the string DOODOO can
// be factored as (DOO)2, but also as (D(O)2)2. Naturally, the latter factoring is considered better since
// it cannot be factored any further. We say that a factoring is irreducible if it does not contain any
// consecutive repetition of a substring. A string may have several irreducible factorings, as seen by the
// example string POPPOP. It can be factored as (POP)2, as well as PO(P)2OP. The first factoring has
// a shorter representation and motivates the following definition. The weight of a factoring, equals the
// number of characters in it, excluding the parentheses and the exponents. Thus the weight of (POP)2
// is 3, whereas PO(P)2OP has weight 5. A maximal factoring is a factoring with the smallest possible
// weight. It should be clear that a maximal factoring is always an irreducible one, but there may still be
// several maximal factorings. E.g. the string ABABA has two maximal factorings (AB)2A and A(BA)2

// Input
// The input consists of several rows. The rows each hold one string of at least one, but less than 80
// characters from the capital alphabet A-Z. The input is terminated by a row containing the character ‘*’
// only. There will be no white space characters in the input.

// Output
// For each string in the input, output one line containing the weight of a maximal factoring of the string.

// Sample Input
// PRATTATTATTIC
// GGGGGGGGG
// PRIME
// BABBABABBABBA
// ARPARPARPARPAR
// *

// Sample Output
// 6
// 1
// 5
// 6
// 5

// s: the min weight of substring [i..j]

#include <iostream>
#include <string>
#include <vector>

unsigned solve(const std::string & s, std::vector<std::vector<unsigned>> & dp, unsigned l, unsigned r) {
    unsigned result, i, j, k;

    if (l == r) {
        result = 1;
    } else if (dp[l][r]) {
        result = dp[l][r];
    } else {
        result = 100;

        for (i = l; i < r; ++i) {
            result = std::min(result, solve(s, dp, l, i) + solve(s, dp, i + 1, r));
        }

        for (i = 1; i + l < r + 1; ++i) {
            if ((r-l+1) % i == 0) {
                for (k = l, j = 0; k <= r && s[k] == s[j+l]; k++) {
                    j++;

                    if (j >= i) {
                        j = 0;
                    }
                }

                if (k == r + 1 && r + 1 != l + i) {
                    result = std::min(result, solve(s, dp, l, l + i - 1));
                }
            }
        }
    }

    return dp[l][r] = result;
}

int main() {
    std::string s;
    std::vector<std::vector<unsigned>> dp;

    while (std::cin >> s && s != "*") {
        dp.resize(s.size(), std::vector<unsigned>(s.size(), 0));
        dp.assign(s.size(), std::vector<unsigned>(s.size(), 0));

        std::cout << solve(s, dp, 0, s.size() - 1) << std::endl;
    }

	return 0;
}
