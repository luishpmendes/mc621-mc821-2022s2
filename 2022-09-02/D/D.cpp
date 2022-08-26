// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=226&page=show_problem&problem=3001

// 11879 - Multiple of 17

// Theorem: If you drop the last digit d of an integer n (n ≥ 10), subtract 5d from the remaining integer, then the difference is a multiple of 17 if and only if n is a multiple of 17.
// For example, 34 is a multiple of 17, because 3-20=-17 is a multiple of 17; 201 is not a multiple of 17, because 20-5=15 is not a multiple of 17.
// Given a positive integer n, your task is to determine whether it is a multiple of 17.

// Input
// There will be at most 10 test cases, each containing a single line with an integer n (1 ≤ n ≤ 10^100).
// The input terminates with n = 0, which should not be processed.

// Output
// For each case, print 1 if the corresponding integer is a multiple of 17, print 0 otherwise.

// Sample Input
// 34
// 201
// 2098765413
// 1717171717171717171717171717171717171717171717171718
// 0

// Sample Output
// 1
// 0
// 1
// 0

#include <algorithm>
#include <iostream>
#include <string>
#include <vector>

// maximum length bignum 
#define MAXDIGITS 123

// positive sign bit
#define PLUS 1
// negative sign bit
#define MINUS -1

typedef struct {
    // represent the number
    char digits[MAXDIGITS];
    // PLUS or MINUS
    int signbit;
    // index of high-order digit
    int lastdigit;
} bignum;

bignum string_to_bignum(const std::string & s);
void print_bignum(const bignum & n);
int compare_bignum(const bignum & a, const bignum & b);
void digit_shift(bignum & n, int d);
void zero_justify(bignum & n);
bignum add_bignum(bignum & a, bignum & b);
bignum subtract_bignum(bignum & a, bignum & b);
bignum multiply_bignum(const bignum & a, const bignum & b);
bignum divide_bignum(bignum & a, bignum & b);

bignum string_to_bignum(const std::string & s) {
    bignum n;
    unsigned i;
    bool is_negative = false;

    if (s.front() == '-') {
        n.signbit = MINUS;
        is_negative = true;
    } else {
        n.signbit = PLUS;
    }

    for (i = 0; i < MAXDIGITS; i++) {
        n.digits[i] = (char) 0;
    }

    n.lastdigit = -1;

    for (i = 0; (is_negative && i < s.size() - 1) || i < s.size(); i++) {
        n.lastdigit++;
        n.digits[n.lastdigit] = s[s.size() - i - 1] - '0';
    }

    return n;
}

void print_bignum(const bignum & n) {
    int i;

    if (n.signbit == MINUS) {
        std::cout << '-';
    }
    
    for (i = 0; i <= n.lastdigit; i++) {
        std::cout << char('0' + n.digits[n.lastdigit - i]);
    }

    std::cout << std::endl;
}

int compare_bignum(const bignum & a, const bignum & b) {
    // counter
    int i;

    if ((a.signbit == MINUS) && (b.signbit == PLUS)) {
        return PLUS;
    }

    if ((a.signbit == PLUS) && (b.signbit == MINUS)) {
        return MINUS;
    }

    if (b.lastdigit > a.lastdigit) {
        return PLUS * a.signbit;
    }

    if (a.lastdigit > b.lastdigit) {
        return MINUS * a.signbit;
    }

    for (i = a.lastdigit; i >= 0; i--) {
        if (a.digits[i] > b.digits[i]) {
            return MINUS * a.signbit;
        }

        if (b.digits[i] > a.digits[i]) {
            return PLUS * a.signbit;
        }
    }

    return 0;
}

// multiply n by 10^d
void digit_shift(bignum & n, int d) {
    // counter
    int i;
    
    if ((n.lastdigit == 0) && (n.digits[0] == 0)) {
        return;
    }
    
    for (i = n.lastdigit; i >= 0; i--) {
        n.digits[i + d] = n.digits[i];
    }
    
    for (i = 0; i < d; i++) {
        n.digits[i] = 0;
    }
    
    n.lastdigit += d;
}

void zero_justify(bignum & n) {
    while ((n.lastdigit > 0) && (n.digits[n.lastdigit] == 0)) {
        n.lastdigit--;
    }
    
    if ((n.lastdigit == 0) && (n.digits[0] == 0)) {
        // hack to avoid -0
        n.signbit = PLUS;
    }
}

bignum add_bignum(bignum & a, bignum & b) {
    bignum c;
    // carry digit
    int carry;
    // counter
    int i;

    c = string_to_bignum("0");
    
    if (a.signbit == b.signbit) {
        c.signbit = a.signbit;
    } else {
        if (a.signbit == MINUS) {
            a.signbit = PLUS;
            c = subtract_bignum(b, a);
            a.signbit = MINUS;
        } else {
            b.signbit = PLUS;
            c = subtract_bignum(a, b);
            b.signbit = MINUS;
        }

        return c;
    }

    c.lastdigit = std::max(a.lastdigit, b.lastdigit) + 1;
    carry = 0;

    for (i = 0; i <= c.lastdigit; i++) {
        c.digits[i] = (char) (carry + a.digits[i] + b.digits[i]) % 10;
        carry = (carry + a.digits[i] + b.digits[i]) / 10;
    }

    zero_justify(c);

    return c;
}

bignum subtract_bignum(bignum & a, bignum & b) {
    bignum c;
    // anything borrowed?
    int borrow;
    // placeholder digit
    int v;
    // counter
    int i;

    c = string_to_bignum("0");

    if ((a.signbit == MINUS) || (b.signbit == MINUS)) {
        b.signbit = -1 * b.signbit;
        c = add_bignum(a, b);
        b.signbit = -1 * b.signbit;

        return c;
    }

    if (compare_bignum(a, b) == PLUS) {
        c = subtract_bignum(b, a);
        c.signbit = MINUS;

        return c;
    }

    c.lastdigit = std::max(a.lastdigit, b.lastdigit);
    borrow = 0;

    for (i = 0; i <= (c.lastdigit); i++) {
        v = (a.digits[i] - borrow - b.digits[i]);

        if (a.digits[i] > 0) {
            borrow = 0;
        }

        if (v < 0) {
            v = v + 10;
            borrow = 1;
        }

        c.digits[i] = (char) v % 10;
    }

    zero_justify(c);

    return c;
}

bignum multiply_bignum(const bignum & a, const bignum & b) {
    bignum c;
    // represent shifted row
    bignum row;
    // placeholder bignum
    bignum tmp;
    // counters
    int i,j;

    c = string_to_bignum("0");

    row = a;

    for (i = 0; i <= b.lastdigit; i++) {
        for (j = 1; j <= b.digits[i]; j++) {
            tmp = add_bignum(c, row);
            c = tmp;
        }

        digit_shift(row, 1);
    }

    c.signbit = a.signbit * b.signbit;
    
    zero_justify(c);

    return c;
}

bignum divide_bignum(bignum & a, bignum & b) {
    bignum c;
    // represent shifted row
    bignum row;
    // placeholder bignum
    bignum tmp;
    // temporary signs
    int asign, bsign;
    // counter
    int i;

    c = string_to_bignum("0");

    c.signbit = a.signbit * b.signbit;

    asign = a.signbit;
    bsign = b.signbit;

    a.signbit = PLUS;
    b.signbit = PLUS;

    row = string_to_bignum("0");
    tmp = string_to_bignum("0");

    c.lastdigit = a.lastdigit;

    for (i = a.lastdigit; i >= 0; i--) {
        digit_shift(row, 1);
        row.digits[0] = a.digits[i];
        c.digits[i] = 0;

        while (compare_bignum(row, b) != PLUS) {
            c.digits[i] ++;
            tmp = subtract_bignum(row, b);
            row = tmp;
        }
    }

    zero_justify(c);

    a.signbit = asign;
    b.signbit = bsign;

    return c;
}

int main() {
    std::string s;
    bignum n, q, p, seventeen;
    seventeen = string_to_bignum("17");

    while (std::cin >> s && s != "0") {
        n = string_to_bignum(s);
        q = divide_bignum(n, seventeen);
        p = multiply_bignum(q, seventeen);

        if (compare_bignum(n, p) == 0) {
            std::cout << 1 << std::endl;
        } else {
            std::cout << 0 << std::endl;
        }
    }

	return 0;
}
