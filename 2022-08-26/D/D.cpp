// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=279&page=show_problem&problem=3914

// 12470 - Tribonacci

// Everybody knows about Fibonacci, they guy who invented the famous sequence where the first two terms are 0 and 1 and from then on every term is the sum of the previous two.
// What most people don’t know is that he had a somewhat mentally retarded brother named Tribonacci. In a desperate attempt to surpass his brother and achieve eternal glory, Tribonacci invented his own sequence: the first three terms are 0, 1, 2 and from then on every term is the sum of the previous three.
// Sadly, regardless of enormous courage and dedication, Tribonacci was never able to compute more than the first 3 terms of his sequence. Even more sadly, one cold night he performed an extraordinary mental effort that dilated one of the blood vessels in his brain, causing severe hemorrhage and killing him instantly. This is clinically known as an aneurysm, but of course Tribonacci did not know this (it is suspected that even pronouncing the word aneurysm would have been an impossible task for him).
// Write a program that changes history and finds the n-th term in the Tribonacci sequence modulo 1,000,000,009.

// Input
// The input contains several test cases (at most 400).
// Each test case contains a single integer n (1 ≤ n ≤ 1016), the desired term in the Tribonacci sequence.
// The last line of the input contains a single ‘0’ and should not be processed.

// Output
// For each test case, output the n–th term in the Tribonacci sequence on a single line. This number might be huge, so output the number modulo 1,000,000,009.

// Sample Input
// 1
// 2
// 3
// 4
// 5
// 6
// 7
// 8
// 9
// 10
// 100
// 10000
// 10000000
// 100000000000
// 10000000000000000
// 0

// Sample Output
// 0
// 1
// 2
// 3
// 6
// 11
// 20
// 37
// 68
// 125
// 616688122
// 335363379
// 272924536
// 48407255
// 163452242

#include <iostream>
#include <vector>

// O(n^3)
std::vector<std::vector<unsigned long>> mat_mul (const std::vector<std::vector<unsigned long>> & a,
                                            const std::vector<std::vector<unsigned long>> & b) {
    std::vector<std::vector<unsigned long>> ans (a.size(), std::vector<unsigned long> (b.front().size()));
    unsigned i, j, k;

    for (i = 0; i < ans.size(); i++) {
        for (j = 0; j < ans[i].size(); j++) {
            ans[i][j] = 0;

            for (k = 0; k < a[i].size(); k++) {
                ans[i][j] += a[i][k] * b[k][j];
            }

            ans[i][j] %= 1000000009;
        }
    }

    return ans;
}

// O(n^3 log p)
std::vector<std::vector<unsigned long>> mat_pow (std::vector<std::vector<unsigned long>> base, unsigned long p) {
    std::vector<std::vector<unsigned long>> ans (base.size(), std::vector<unsigned long> (base.front().size()));
    unsigned i, j;

    // prepare identity matrix
    for (i = 0; i < ans.size(); i++) {
        for (j = 0; j < ans[i].size(); j++) {
            ans[i][j] = (i == j);
        }
    }

    // iterative version of Divide & Conquer exponentiation
    while (p) {
        // if p is odd (last bit is on)
        if (p & 1) {
            ans = mat_mul(ans, base);
        }

        // square the base
        base = mat_mul(base, base);
        // divide p by 2
        p >>= 1;
    }

    return ans;
}

int main() {
    unsigned long n;
    std::vector<std::vector<unsigned long>> base (3, std::vector<unsigned long> (3)), ans;
    base[0][0] = 1; base[0][1] = 1; base[0][2] = 1;
    base[1][0] = 1; base[1][1] = 0; base[1][2] = 0;
    base[2][0] = 0; base[2][1] = 1; base[2][2] = 0;

    while (std::cin >> n && n) {
        if (n <= 3) {
            std::cout << n-1 << std::endl;
        } else { // n >= 3
            ans = mat_pow(base, n - 1);
            std::cout << (2 * ans[2][0] + ans[2][1]) % 1000000009 << std::endl;
        }
    }

	return 0;
}
