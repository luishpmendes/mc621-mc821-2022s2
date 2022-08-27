// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=17&page=show_problem&problem=1464

// 10523 - Very Easy !!!

// Most of the times, the students of Computer Science & Engineering of BUET deal with bogus, tough and very complex formulae. That is why, sometimes, even for a easy problem they think very hard and make the problem much complex to solve. But, the team members of the team “BUET PESSIMISTIC” are the only exceptions. Just like the opposite manner, they treat every hard problem as easy and so cannot do well in any contest. Today, they try to solve a series but fail for treating it as hard. Let them help.

// Input
// Just try to determine the answer for the following series
// ∑_{i = 1}^{N}{i*A^i}
// You are given the value of integers N and A (1 ≤ N ≤ 150, 0 ≤ A ≤ 15)

// Output
// For each line of the input, your correct program should output the integer value of the sum in separate lines for each pair of values of N and A.

// Sample Input
// 3 3
// 4 4

// Sample Output
// 102
// 1252

#include <algorithm>
#include <iostream>
#include <sstream>
#include <string>
#include <vector>

// maximum length bignum 
#define MAXDIGITS 250

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

bignum int_to_bignum(const int & n);
std::string bignum_to_string(const bignum & n);
int compare_bignum(const bignum & a, const bignum & b);
void digit_shift(bignum & n, int d);
void zero_justify(bignum & n);
bignum add_bignum(bignum & a, bignum & b);
bignum subtract_bignum(bignum & a, bignum & b);
bignum multiply_bignum(const bignum & a, const bignum & b);
bignum divide_bignum(bignum & a, bignum & b);
bignum power_bignum(const bignum & base, bignum & exponent);

bignum int_to_bignum(const int & n) {
    bignum m;
    int i, t;

    if (n >= 0) {
        m.signbit = PLUS;
        t = n;
    } else {
        m.signbit = MINUS;
        t = -n;
    }

    for (i = 0; i < MAXDIGITS; i++) {
        m.digits[i] = (char) 0;
    }
    
    m.lastdigit = -1;

    while (t > 0) {
        m.lastdigit++;
        m.digits[m.lastdigit] = (char) (t % 10);
        t /= 10;
    }

    if (n == 0) {
        m.lastdigit = 0;
    }

    return m;
}

std::string bignum_to_string(const bignum & n) {
    std::stringstream ss;
    int i;

    if (n.signbit == MINUS) {
        ss << '-';
    }
    
    for (i = 0; i <= n.lastdigit; i++) {
        ss << char('0' + n.digits[n.lastdigit - i]);
    }

    return ss.str();
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

    c = int_to_bignum(0);
    
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

    c = int_to_bignum(0);

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

    c = int_to_bignum(0);

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

    c = int_to_bignum(0);

    c.signbit = a.signbit * b.signbit;

    asign = a.signbit;
    bsign = b.signbit;

    a.signbit = PLUS;
    b.signbit = PLUS;

    row = int_to_bignum(0);
    tmp = int_to_bignum(0);

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

bignum power_bignum(const bignum & base, bignum & exponent) {
    bignum power,
           zero = int_to_bignum(0),
           one = int_to_bignum(1),
           two = int_to_bignum(2);

    if (compare_bignum(exponent, zero) == 0) {
        power = int_to_bignum(1);
    } else if ((exponent.digits[0] & 1) == 0) { // exponent % 2 == 0
        exponent = divide_bignum(exponent, two);
        power = power_bignum(base, exponent);
        power = multiply_bignum(power, power);
    } else { // exponent % 2 == 1
        exponent = subtract_bignum(exponent, one);
        exponent = divide_bignum(exponent, two);
        power = power_bignum(base, exponent);
        power = multiply_bignum(power, power);
        power = multiply_bignum(power, base);
    }

    return power;
}

int main() {
    int n, a;
    bignum big_n, big_a,
           one,
           two,
           a_minus_one,
           n_plus_one,
           a_to_n_plus_one,
           numerator,
           denominator,
           ans;

    one = int_to_bignum(1);
    two = int_to_bignum(2);

    while (std::cin >> n >> a) {
        big_n = int_to_bignum(n);
        n_plus_one = add_bignum(big_n, one);

        if (a > 1) {
            big_a = int_to_bignum(a);
            a_minus_one = subtract_bignum(big_a, one);
            a_to_n_plus_one = power_bignum(big_a, n_plus_one);
            numerator = multiply_bignum(a_minus_one, big_n);
            numerator = subtract_bignum(numerator, one);
            numerator = multiply_bignum(numerator, a_to_n_plus_one);
            numerator = add_bignum(numerator, big_a);
            denominator = multiply_bignum(a_minus_one, a_minus_one);
            ans = divide_bignum(numerator, denominator);
        } else if (a == 1) {
            ans = multiply_bignum(big_n, n_plus_one);
            ans = divide_bignum(ans, two);
        } else { // a == 0
            ans = int_to_bignum(0);
        }

        std::cout << bignum_to_string(ans) << std::endl;
    }

	return 0;
}
