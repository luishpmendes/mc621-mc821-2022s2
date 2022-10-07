// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=1456

// 10515 - Powers Et Al.
// Finding the exponent of any number can be very troublesome as it grows exponentially. But in this
// problem you will have to do a very simple task. Given two non-negative numbers m and n, you will
// have to find the last digit of m^n in decimal number system.

// Input
// The input file contains less than 100000 lines. Each line contains two integers m and n (Less than 10^101.
// Input is terminated by a line containing two zeroes. This line should not be processed.

// Output
// For each set of input you must produce one line of output which contains a single digit. This digit is
// the last digit of m^n.

// Sample Input
// 2 2
// 2 5
// 0 0

// Sample Output
// 4
// 2

// concentrate on the last digit

#include <iostream>
#include <string>

int main() {
    std::string m, n;
    unsigned last_digit_m, last_digits_n;

    while (std::cin >> m >> n && (m != "0" || n != "0")) {
        last_digit_m = m.back() - '0';

        if (n == "0") {
            last_digit_m = 1;
        } else {
            last_digits_n = (n[n.size()-2] - '0')*10 + n.back() - '0';
            if (last_digit_m == 2) {
                switch (last_digits_n & 3) {
                    case 0:
                        last_digit_m = 6;
                        break;
                    case 1:
                        last_digit_m = 2;
                        break;
                    case 2:
                        last_digit_m = 4;
                        break;
                    case 3:
                        last_digit_m = 8;
                        break;
                }
            } else if (last_digit_m == 3) {
                switch (last_digits_n & 3) {
                    case 0:
                        last_digit_m = 1;
                        break;
                    case 1:
                        last_digit_m = 3;
                        break;
                    case 2:
                        last_digit_m = 9;
                        break;
                    case 3:
                        last_digit_m = 7;
                        break;
                }
            } else if (last_digit_m == 4) {
                switch (last_digits_n & 1) {
                    case 0:
                        last_digit_m = 6;
                        break;
                    case 1:
                        last_digit_m = 4;
                        break;
                }
            } else if (last_digit_m == 7) {
                switch (last_digits_n & 3) {
                    case 0:
                        last_digit_m = 1;
                        break;
                    case 1:
                        last_digit_m = 7;
                        break;
                    case 2:
                        last_digit_m = 9;
                        break;
                    case 3:
                        last_digit_m = 3;
                        break;
                }
            } else if (last_digit_m == 8) {
                switch (last_digits_n & 3) {
                    case 0:
                        last_digit_m = 6;
                        break;
                    case 1:
                        last_digit_m = 8;
                        break;
                    case 2:
                        last_digit_m = 4;
                        break;
                    case 3:
                        last_digit_m = 2;
                        break;
                }
            } else if (last_digit_m == 9) {
                switch (last_digits_n & 1) {
                    case 0:
                        last_digit_m = 1;
                        break;
                    case 1:
                        last_digit_m = 9;
                        break;
                }
            }
        }

        std::cout << last_digit_m << std::endl;
    }

	return 0;
}
