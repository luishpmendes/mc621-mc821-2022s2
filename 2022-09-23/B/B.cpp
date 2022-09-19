// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=7&page=show_problem&problem=484

// 543 - Goldbach's Conjecture

// In 1742, Christian Goldbach, a German amateur mathematician, sent a letter to Leonhard Euler in which he made the following conjecture:
// Every number greater than 2 can be written as the sum of three prime numbers.
// Goldbach was considering 1 as a primer number, a convention that is no longer followed. Later on, Euler re-expressed the conjecture as:
// Every even number greater than or equal to 4 can be expressed as the sum of two prime numbers.
// For example:
//  • 8 = 3 + 5. Both 3 and 5 are odd prime numbers.
//  • 20 = 3 + 17 = 7 + 13.
//  • 42 = 5 + 37 = 11 + 31 = 13 + 29 = 19 + 23.
// Today it is still unproven whether the conjecture is right. (Oh wait, I have the proof of course, but it is too long to write it on the margin of this page.)
// Anyway, your task is now to verify Goldbach’s conjecture as expressed by Euler for all even numbers less than a million.

// Input
// The input file will contain one or more test cases.
// Each test case consists of one even integer n with 6 ≤ n < 1000000.
// Input will be terminated by a value of 0 for n.

// Output
// For each test case, print one line of the form n = a + b, where a and b are odd primes. Numbers and operators should be separated by exactly one blank like in the sample output below. If there is more than one pair of odd primes adding up to n, choose the pair where the difference b − a is maximized.
// If there is no such pair, print a line saying ‘Goldbach's conjecture is wrong.’

// Sample Input
// 8
// 20
// 42
// 0

// Sample Output
// 8 = 3 + 5
// 20 = 3 + 17
// 42 = 5 + 37

#include <iomanip>
#include <iostream>
#include <vector>
#include <bitset> // compact STL for Sieve, better than vector<bool>!


// create list of primes in [0..upperbound]
unsigned long sieve(unsigned long upperbound,
                    std::bitset<10000010> & bs,
                    std::vector<unsigned long> & primes) {
    // add 1 to include upperbound
    unsigned long _sieve_size = upperbound + 1;
    // set all bits to 1
    bs.set();
    // except index 0 and 1
    bs[0] = bs[1] = 0;
    primes.clear();

    for (unsigned long i = 2; i <= _sieve_size; i++) {
        if (bs[i]) {
            // cross out multiples of i starting from i * i!
            for (unsigned long j = i * i; j <= _sieve_size; j += i) {
                bs[j] = 0;
            }

            // add this prime to the list of primes
            primes.push_back(i);
        }
    }

    return _sieve_size;
}

// a good enough deterministic prime tester
// note: only work for N <= (last prime in vi "primes")^2
bool isPrime(unsigned long N,
             const unsigned long & _sieve_size,
             const std::bitset<10000010> & bs,
             const std::vector<unsigned long> & primes) {
    if (N <= _sieve_size) {
        // O(1) for small primes
        return bs[N]; 
    }
    
    for (unsigned i = 0; i < primes.size(); i++) {
        if (N % primes[i] == 0) {
            return false;
        }
    }

    // it takes longer time if N is a large prime!
    return true; 
} 

bool goldbachsConjecture(unsigned n,
                         const unsigned long & _sieve_size,
                         const std::bitset<10000010> & bs,
                         const std::vector<unsigned long> & primes,
                         unsigned & a,
                         unsigned & b) {
    bool result = false;

    for (unsigned i = 0; i < primes.size() && primes[i] <= n && !result; i++) {
        if (isPrime(n - primes[i], _sieve_size, bs, primes)) {
            a = primes[i];
            b = n - primes[i];
            result = true;
        }
    }

    return result;
}

int main() {
    unsigned n, max_n = 0;
    unsigned long _sieve_size;
    std::bitset<10000010> bs;
    std::vector<unsigned long> primes;
    std::vector<unsigned> ns;

    

    while (std::cin >> n && n) {
        ns.push_back(n);

        if (max_n < n) {
            max_n = n;
        }
    }

    _sieve_size = sieve(max_n, bs, primes);

    for (const unsigned & x : ns) {
        unsigned a, b;

        if (goldbachsConjecture(x, _sieve_size, bs, primes, a, b)) {
            std::cout << x << " = " << a << " + " << b << std::endl;
        } else {
            std::cout << "Goldbach's conjecture is wrong." << std::endl;
        }
    }

	return 0;
}
