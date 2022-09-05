// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=242&page=show_problem&problem=3175

// 12024 - Hats

// John Hatman, the honest cloakroom attendant of the Royal Theatre of London, would like to know the solution to the following problem.
// When the show finishes, all spectators in the theatre are in a hurry to see the Final of the UEFA Championship. So, they run to the cloakroom to take their hats back.
// Some of them take a wrong hat. But, how likely is that everyone take a wrong hat?

// Input
// The first line of the input contains an integer, t, indicating the number of test cases. For each test case, one line appears, that contains a number n, 2 ≤ n ≤ 12, representing the number of people and hats

// Output
// For each test case, the output should contain a single line with the number representing the number of favourable cases (i.e., the number of cases where all people take a wrong hat), followed by a bar, ‘/’, and followed by a number representing the total number of possible cases.

// Sample Input
// 3
// 2
// 3
// 4

// Sample Output
// 1/2
// 2/6
// 9/24

#include <iostream>
#include <vector>

int main() {
    unsigned t, n, i;
    std::vector<unsigned> factorial(13),
                          derangement(13);
    
    factorial[0] = 1;
    factorial[1] = 1;
    derangement[0] = 1;
    derangement[1] = 0;

    for (i = 2; i <= 12; ++i) {
        factorial[i] = factorial[i - 1] * i;
        derangement[i] = (i - 1) * (derangement[i - 1] + derangement[i - 2]);
    }

    while (std::cin >> t) {
        while (t--) {
            std::cin >> n;

            std::cout << derangement[n] << '/' << factorial[n] << std::endl;
        }
    }

	return 0;
}
