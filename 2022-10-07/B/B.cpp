// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=27&page=show_problem&problem=2544

// 11549 - Calculator Conundrum
// Alice got a hold of an old calculator that can display n digits. She was bored enough to come up with
// the following time waster.
// She enters a number k then repeatedly squares it until the result overflows. When the result
// overflows, only the n most significant digits are displayed on the screen and an error flag appears. Alice
// can clear the error and continue squaring the displayed number. She got bored by this soon enough,
// but wondered:
// “Given n and k, what is the largest number I can get by wasting time in this manner?”

// Input
// The first line of the input contains an integer t (1 ≤ t ≤ 200), the number of test cases. Each test case
// contains two integers n (1 ≤ n ≤ 9) and k (0 ≤ k < 10^n) where n is the number of digits this calculator
// can display k is the starting number.

// Output
// For each test case, print the maximum number that Alice can get by repeatedly squaring the starting
// number as described.

// Sample Input
// 2
// 1 6
// 2 99

// Sample Output
// 9
// 99

// repeat squaring with limited digits until
// it cycles; that is, the Floyd’s cycle-finding algorithm is only used to detect the
// cycle, we do not use the value of µ or λ; instead, we keep track the largest iterated
// function value found before any cycle is encountered

#include <iostream>
#include <utility>
#include <vector>

unsigned buffer[100];

unsigned long f(unsigned long n, unsigned long x) {
    unsigned long y = x * x,
                  tmp = 0, i;

    while (y > 0) {
        buffer[tmp++] = y % 10;
        y /= 10;
    }

    if (n > tmp) {
        n = tmp;
    }

    y = 0;

    for (i = 0; i < n; i++) {
        y = y * 10 + buffer[--tmp];
    }

    return y;
}

unsigned long floydCycleFinding(unsigned long n, unsigned long k) {
    // 1st part: finding k*mu, hare’s speed is 2x tortoise’s
    unsigned long tortoise = k,
                  hare = k, // f(x0) is the node next to x0
                  result = k;

    do {
        tortoise = f(n, tortoise);

        if (result < tortoise) {
            result = tortoise;
        }

        hare = f(n, hare);

        if (result < hare) {
            result = hare;
        }

        hare = f(n, hare);

        if (result < hare) {
            result = hare;
        }
    } while (tortoise != hare);

    return result;
}

int main() {
    unsigned long t, n, k;

    while (std::cin >> t) {
        while (t--) {
            std::cin >> n >> k;
            std::cout << floydCycleFinding(n, k) << std::endl;
        }
    }

	return 0;
}
