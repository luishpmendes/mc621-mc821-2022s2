// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=22&page=show_problem&problem=1962

// 11021 - Tribles

// GRAVITATION, n.
// “The tendency of all bodies to approach one another with a strength proportion to the quantity of matter they contain – the quantity of matter they contain being ascertained by the strength of their tendency to approach one another. This is a lovely and edifying illustration of how science, having made A the proof of B, makes B the proof of A.”
// Ambrose Bierce

// You have a population of k Tribbles. This particular species of Tribbles live for exactly one day and then die. Just before death, a single Tribble has the probability Pi of giving birth to i more Tribbles. What is the probability that after m generations, every Tribble will be dead?

// Input
// The first line of input gives the number of cases, N. N test cases follow. Each one starts with a line containing n (1 ≤ n ≤ 1000), k (0 ≤ k ≤ 1000) and m (0 ≤ m ≤ 1000). The next n lines will give the probabilities P0, P1, . . . , Pn−1.

// Output
// For each test case, output one line containing ‘Case #x:’ followed by the answer, correct up to an absolute or relative error of 10^(−6).

// Sample Input
// 4
// 3 1 1
// 0.33
// 0.34
// 0.33
// 3 1 2
// 0.33
// 0.34
// 0.33
// 3 1 2
// 0.5
// 0.0
// 0.5
// 4 2 2
// 0.5
// 0.0
// 0.0
// 0.5

// Sample Output
// Case #1: 0.3300000
// Case #2: 0.4781370
// Case #3: 0.6250000
// Case #4: 0.3164062

#include <iomanip>
#include <iostream>
#include <vector>

double power (double base, unsigned exponent) {
    double result = 1.0;

    while (exponent > 0) {
        if ((exponent & 1) == 1) {
            result *= base;
        }
        base *= base;
        exponent >>= 1;
    }

    return result;
}

int main() {
    unsigned t, c, n, k, m, i, j;
    std::vector<double> p, dp;
    double aux, ans;

    while (std::cin >> t) {
        for (c = 1; c <= t; c++) {
            std::cin >> n >> k >> m;

            // p[i] : probability of a single Tribble giving birth to i more Tribbles
            p.resize(n);
            // dp[i] : probability that all offspring of a single Tribble will be dead on day i
            dp.resize(m + 1);

            for (i = 0; i < n; i++) {
                std::cin >> p[i];
            }

            dp[0] = 0.0;

            for (i = 1; i <= m; i++) {
                // dp[i] = p[0] + \sum_{j = 1}^{n - 1}{p[j] * dp[i - 1]^j}
                dp[i] = p[0];
                aux = dp[i - 1];

                for (j = 1; j < n; j++) {
                    dp[i] += p[j] * aux;
                    aux *= dp[i - 1];
                }
            }

            ans = power(dp[m], k);

            std::cout << "Case #" << c << ": " << std::fixed << std::setprecision(7) << ans << std::endl;
        }
    }

	return 0;
}
