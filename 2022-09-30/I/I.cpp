// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=22&page=show_problem&problem=1970

// 11029 - Leading and Trailing
// Apart from the novice programmers, all others know that you can’t exactly represent numbers raised
// to some high power. For example, the C function pow(125456, 455) can be represented in double data
// type format, but you won’t get all the digits of the result. However we can get at least some satisfaction
// if we could know few of the leading and trailing digits. This is the requirement of this problem.

// Input
// The first line of input will be an integer T < 1001, where T represents the number of test cases. Each
// of the next T lines contains two positive integers, n and k. n will fit in 32 bit integer and k will be less
// than 10000001.

// Output
// For each line of input there will be one line of output. It will be of the format LLL . . . T T T, where LLL
// represents the first three digits of n^k and T T T represents the last three digits of n^k. You are assured
// that n^k will contain at least 6 digits.

// Sample Input
// 2
// 123456 1
// 123456 2

// Sample Output
// 123...456
// 152...936

// combination of logarithmic trick to get
// the first three digits and ‘big mod’ trick to get the last three digits

#include <iomanip>
#include <iostream>
#include <cmath>

unsigned long modular_pow(unsigned long base, unsigned long exp, unsigned long modulus) {
    unsigned long result = 1;

    while (exp > 0) {
        if (exp & 1) {
            result = (result * base) % modulus;
        }

        exp = exp >> 1;
        base = (base * base) % modulus;
    }

    return result;
}

int main() {
    unsigned long t, n, k, leading, trailing;
    long double power;

    while (std::cin >> t) {
        while (t--) {
            std::cin >> n >> k;

            power = (double) k * log10(n);
            leading = pow(10, power - floor(power) + 2);
            trailing = modular_pow(n, k, 1000);

            std::cout << leading << "..." << std::setw(3) << std::setfill('0') << trailing << "\n";
        }
    }

	return 0;
}
