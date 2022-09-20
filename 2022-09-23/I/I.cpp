// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=20&page=show_problem&problem=1833

// 10892 - LCM Cardinality

// A pair of numbers has a unique LCM but a single number can be the LCM of more than one possible pairs. For example 12 is the LCM of (1, 12), (2, 12), (3,4) etc. For a given positive integer N, the number of different integer pairs with LCM is equal to N can be called the LCM cardinality of that number N. In this problem your job is to find out the LCM cardinality of a number.

// Input
// The input file contains at most 101 lines of inputs. Each line contains an integer N (0 < N ≤ 2 ∗ 109). Input is terminated by a line containing a single zero. This line should not be processed.

// Output
// For each line of input except the last one produce one line of output. This line contains two integers N and C. Here N is the input number and C is its cardinality. These two numbers are separated by a single space.

// Sample Input
// 2
// 12
// 24
// 101101291
// 0

// Sample Output
// 2 2
// 12 8
// 24 11
// 101101291 5

#include <algorithm>
#include <iostream>
#include <vector>

unsigned long gcd(unsigned long a, unsigned long b) {
    return b == 0 ? a : gcd(b, a % b);
}

unsigned long lcm(unsigned long a, unsigned long b) {
    return a * (b / gcd(a, b));
}

unsigned LCM_Cardinality(unsigned long n) { 
    std::vector<unsigned long> factors;
    unsigned i, j, cardinality = 0;

    for (i = 1; i * i <= n; i++) {
        if (n % i == 0) {
            factors.push_back(i);

            if (i * i != n) {
                factors.push_back(n / i);
            }
        }
    }

    for (i = 0; i < factors.size(); i++) {
        for (j = i; j < factors.size(); j++) {
            if (lcm(factors[i], factors[j]) == n) {
                cardinality++;
            }
        }
    }

    return cardinality;
}

int main() {
    unsigned long n;

    while (std::cin >> n && n) {
        std::cout << n << ' ' << LCM_Cardinality(n) << std::endl;
    }

	return 0;
}
