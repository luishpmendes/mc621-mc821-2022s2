// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=14&page=show_problem&problem=1153

// 10212 - The Last Non-zero Digit.
// In this problem you will be given two decimal integer number N, M. You will have to find the last
// non-zero digit of the P^{N}_{M}. This means no of permutations of N things taking M at a time.

// Input
// The input file contains several lines of input. Each line of the input file contains two integers N
// (0 ≤ N ≤ 20000000), M (0 ≤ M ≤ N). Input is terminated by end-of-file.

// Output
// For each line of the input file you should output a single digit, which is the last non-zero digit of P^{N}_{M}.
// For example, if P^{N}_{M} is 720 then the last non-zero digit is 2. So in this case your output should be 2.

// Sample Input
// 10 10
// 10 5
// 25 6

// Sample Output
// 8
// 4
// 2

// there is a modulo arithmetic
// solution: multiply numbers from N down to N − M + 1; repeatedly use /10
// to discard the trailing zero(es), and then use %1 Billion to only memorize
// the last few (maximum 9) non zero digits

#include <iostream>
#include <vector>

int main() {
    unsigned long long n, m, ans, i;

    while (std::cin >> n >> m) {
        ans = 1;
        for (i = n; i + m > n; i--) {
            ans *= i;

            while (ans % 10 == 0) {
                ans /= 10;
            }

            ans %= 1000000000;
        }

        std::cout << ans % 10 << std::endl;
    }

	return 0;
}
