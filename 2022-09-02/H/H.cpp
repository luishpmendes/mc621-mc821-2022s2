// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=9&page=show_problem&problem=654

// 713 - Adding Reversed Numbers

// The Antique Comedians of Malidinesia prefer comedies to tragedies. Unfortunately, most of the ancient plays are tragedies. Therefore the dramatic advisor of ACM has decided to transfigure some tragedies into comedies. Obviously, this work is very hard because the basic sense of the play must be kept intact, although all the things change to their opposites. For example the numbers: if any number appears in the tragedy, it must be converted to its reversed form before being accepted into the comedy play.
// Reversed number is a number written in arabic numerals but the order of digits is reversed. The first digit becomes last and vice versa. For example, if the main hero had 1245 strawberries in the tragedy, he has 5421 of them now. Note that all the leading zeros are omitted. That means if the number ends with a zero, the zero is lost by reversing (e.g. 1200 gives 21). Also note that the reversed number never has any trailing zeros.
// ACM needs to calculate with reversed numbers. Your task is to add two reversed numbers and output their reversed sum. Of course, the result is not unique because any particular number is a reversed form of several numbers (e.g. 21 could be 12, 120 or 1200 before reversing). Thus we must assume that no zeros were lost by reversing (e.g. assume that the original number was 12).

// Input
// The input consists of N cases. The first line of the input contains only positive integer N. Then follow the cases. Each case consists of exactly one line with two positive integers separated by space. These are the reversed numbers you are to add. Numbers will be at most 200 characters long.

// Output
// For each case, print exactly one line containing only one integer — the reversed sum of two reversed numbers. Omit any leading zeros in the output.

// Sample Input
// 3
// 24 1
// 4358 754
// 305 794

// Sample Output
// 34
// 1998
// 1

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

bignum string_to_bignum(const std::string & s);
std::string bignum_to_string(const bignum & n);
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

    zero_justify(n);

    return n;
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
    unsigned n;
    std::string s;
    bignum a, b, c;

    while (std::cin >> n) {
        while (n--) {
            std::cin >> s;
            std::reverse(s.begin(), s.end());
            a = string_to_bignum(s);
            std::cin >> s;
            std::reverse(s.begin(), s.end());
            b = string_to_bignum(s);
            c = add_bignum(a, b);
            s = bignum_to_string(c);
            std::reverse(s.begin(), s.end());
            c = string_to_bignum(s);
            std::cout << bignum_to_string(c) << std::endl;
        }
    }

	return 0;
}
