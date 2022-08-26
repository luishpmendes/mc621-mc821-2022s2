// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=20&page=show_problem&problem=1811

// 10870 - Recurrences

// Consider recurrent functions of the following form:
// f(n) = a1f(n − 1) + a2f(n − 2) + a3f(n − 3) + ... + adf(n − d), for n > d,
// where a1, a2, ... , ad are arbitrary constants.
// A famous example is the Fibonacci sequence, defined as: f(1) = 1, f(2) = 1, f(n) = f(n − 1) + f(n − 2). Here d = 2, a1 = 1, a2 = 1.
// Every such function is completely described by specifying d (which is called the order of recurrence), values of d coefficients: a1, a2, ... , ad, and values of f(1), f(2), ... , f(d). You’ll be given these numbers, and two integers n and m. Your program’s job is to compute f(n) modulo m.

// Input
// Input file contains several test cases. Each test case begins with three integers: d, n, m, followed by two sets of d non-negative integers. The first set contains coefficients: a1, a2, ... , ad. The second set gives values of f(1), f(2), ... , f(d).
// You can assume that: 1 ≤ d ≤ 15, 1 ≤ n ≤ 2^31 − 1, 1 ≤ m ≤ 46340. All numbers in the input will fit in signed 32-bit integer.
// Input is terminated by line containing three zeroes instead of d, n, m. Two consecutive test cases are separated by a blank line.

// Output
// For each test case, print the value of f(n)( mod m) on a separate line. It must be a non-negative integer, less than m.

// Sample Input
// 1 1 100
// 2
// 1
//
// 2 10 100
// 1 1
// 1 1
//
// 3 2147483647 12345
// 12345678 0 12345
// 1 2 3
//
// 0 0 0

// Sample Output
// 1
// 55
// 423

#include <iostream>
#include <vector>

// O(n^3)
std::vector<std::vector<long>> mat_mul (const std::vector<std::vector<long>> & a,
                                        const std::vector<std::vector<long>> & b,
                                        unsigned m) {
    std::vector<std::vector<long>> ans (a.size(), std::vector<long> (b.front().size()));
    unsigned i, j, k;

    for (i = 0; i < ans.size(); i++) {
        for (j = 0; j < ans[i].size(); j++) {
            ans[i][j] = 0;

            for (k = 0; k < a[i].size(); k++) {
                ans[i][j] += a[i][k] * b[k][j];
            }

            ans[i][j] %= m;
        }
    }

    return ans;
}

// O(n^3 log p)
std::vector<std::vector<long>> mat_pow (std::vector<std::vector<long>> base,
                                        unsigned p,
                                        unsigned m) {
    std::vector<std::vector<long>> ans (base.size(), std::vector<long> (base.front().size()));
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
            ans = mat_mul(ans, base, m);
        }

        // square the base
        base = mat_mul(base, base, m);
        // divide p by 2
        p >>= 1;
    }

    return ans;
}

int main() {
    unsigned d, n, m, i, ans;
    std::vector<long> a, f;
    std::vector<std::vector<long>> base, matrix;

    while (std::cin >> d >> n >> m && (d || n || m)) {
        a.resize(d);
        f.resize(d);
        base.resize(d, std::vector<long>(d));
        base.assign(d, std::vector<long>(d, 0));

        for (i = 0; i < d; i++) {
            std::cin >> a[i];
            a[i] %= m;
        }

        for (i = 0; i < d; i++) {
            std::cin >> f[i];
        }

        for (i = 0; i < d; i++) {
            base[0][i] = a[i];
        }

        for (i = 1; i < d; i++) {
            base[i][i-1] = 1;
        }

        if (n <= d) {
            std::cout << f[n - 1] << std::endl;
        } else { // n >= 3
            matrix = mat_pow(base, n - 1, m);
            ans = 0;

            for (i = 0; i < d; i++) {
                ans += (matrix.back()[d-1-i] * f[i]) % m;
                ans %= m;
            }
            
            std::cout << ans << std::endl;
        }
    }

	return 0;
}
