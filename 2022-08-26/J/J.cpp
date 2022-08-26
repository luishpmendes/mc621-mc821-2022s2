// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=17&page=show_problem&problem=1459

// 10518 - How Many Calls?

// The fibonacci number is defined by the following recurrence:
//  • fib(0) = 0
//  • fib(1) = 1
//  • fib(n) = fib(n − 1) + fib(n − 2)
// But we’re not interested in the fibonacci numbers here. We would like to know how many calls does it take to evaluate the n-th fibonacci number if we follow the given recurrence. Since the numbers are going to be quite large, we’d like to make the job a bit easy for you. We’d only need the last digit of the number of calls, when this number is represented in base b.

// Input
// Input consists of several test cases. For each test you’d be given two integers n (0 ≤ n < 2^63 − 1), b (0 < b ≤ 10000). Input is terminated by a test case where n = 0 and b = 0, you must not process this test case

// Output
// For each test case, print the test case number first. Then print n, b and the last digit (in base b) of the number of calls. There would be a single space in between the two numbers of a line.
// Note that the last digit has to be represented in decimal number system.

// Sample Input
// 0 100
// 1 100
// 2 100
// 3 100
// 10 10
// 0 0

// Sample Output
// Case 1: 0 100 1
// Case 2: 1 100 1
// Case 3: 2 100 3
// Case 4: 3 100 5
// Case 5: 10 10 7

#include <iostream>
#include <vector>

// O(n^3)
std::vector<std::vector<unsigned long>> mat_mul (const std::vector<std::vector<unsigned long>> & a,
                                                 const std::vector<std::vector<unsigned long>> & b,
                                                 unsigned long m) {
    std::vector<std::vector<unsigned long>> ans (a.size(), std::vector<unsigned long> (b.front().size()));
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
std::vector<std::vector<unsigned long>> mat_pow (std::vector<std::vector<unsigned long>> base,
                                                 unsigned long p,
                                                 unsigned long m) {
    std::vector<std::vector<unsigned long>> ans (base.size(),
                                                 std::vector<unsigned long> (base.front().size()));
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
    unsigned long n, b, t;
    std::vector<std::vector<unsigned long>> base (2, std::vector<unsigned long> (2)), ans;

    t = 0;
    base[0][0] = 1; base[0][1] = 1;
    base[1][0] = 1; base[1][1] = 0;

    while (std::cin >> n >> b && (n || b)) {
        t++;
        if (n <= 1) {
            std::cout << "Case " << t << ": " << n << ' ' << b << ' ' << 1 << std::endl;
        } else { // n >= 2
            ans = mat_pow(base, n + 1, b);
            std::cout << "Case " << t << ": " << n << ' ' << b << ' ' << (2*ans[1][0]-1+b)%b << std::endl;
        }
    }

	return 0;
}
