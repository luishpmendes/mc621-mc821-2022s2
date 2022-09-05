// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=23&page=show_problem&problem=2117

// 11176 - Winning Streak

// “You can run on for a long time, sooner or later God’ll cut you down.” — Traditional folk song

// Mikael likes to gamble, and as you know, you can place bets on almost anything these days. A particular thing that has recently caught Mikael’s interest is the length of the longest winning streak of a team during a season (i.e. the highest number of consecutive games won). In order to be able to make smarter bets, Mikael has asked you to write a program to help him compute the expected value of the longest winning streak of his favourite teams.
// In general, the probability that a team wins a game depends on a lot of different factors, such as whether they’re the home team, whether some key player is injured, and so on. For the first prototype of the program, however, we simplify this, and assume that all games have the same fixed probability p of being won, and that the result of a game does not affect the win probability for subsequent games.
// The expected value of the longest streak is the average of the longest streak in all possible outcomes of all games in a season, weighted by their probability. For instance, assume that the season consists of only three games, and that p = 0.4. There are eight different outcomes, which we can represent by a string of ‘W’:s and ‘L’:s, indicating which games were won and which games were lost (for example, ‘WLW’ indicates that the team won the first and the third game, but lost the second). The possible results of the season are:
// Result LLL LLW LWL LWW WLL WLW WWL WWW
// Probability 0.216 0.144 0.144 0.096 0.144 0.096 0.096 0.064
// Streak 0 1 1 2 1 1 2 3
// In this case, the expected length of the longest winning streak becomes
// 0.216 · 0 + 0.144 · 1 + 0.144 · 1 + 0.096 · 2 + 0.144 · 1 + 0.096 · 1 + 0.096 · 2 + 0.064 · 3 = 1.104

// Input
// Several test cases (at most 40), each containing an integer 1 ≤ n ≤ 500 giving the number of games in a season, and a floating point number 0 ≤ p ≤ 1, the win probability. Input is terminated by a case where n = 0, which should not be processed.

// Output
// For each test case, give the expected length of the longest winning streak. The answer should be given as a floating point number with an absolute error of at most 10^(−4).

// Sample Input
// 3 0.4
// 10 0.75
// 0 0.5

// Sample Output
// 1.104000
// 5.068090

#include <iomanip>
#include <iostream>
#include <vector>

int main() {
    unsigned n, i, j;
    double p, ans;
    // pp[i] will have p^i
    std::vector<double> pp;
    // dp[i][j] will have the probability of having a maximum winning streak of j in i consecutive games
    std::vector<std::vector<double>> dp;

    while (std::cin >> n >> p && n != 0) {
        dp.resize(n + 1, std::vector<double>(n + 1, 1.0));
        dp.assign(n + 1, std::vector<double>(n + 1, 1.0));
        pp.resize(n + 1);

        pp[0] = 1.0;

        for (i = 1; i <= n; i++) {
            pp[i] = pp[i - 1] * p;
        }
 
        for (i = 1; i <= n; i++) {
            for (j = 0; j <= n; j++) {
                dp[i][j] = dp[i - 1][j];

                if (i == j + 1) {
                    // dp[i][j] = dp[i - 1][j] - p^(j + 1);
                    dp[i][j] -= pp[j + 1];
                } else if (i > j + 1) { 
                    // dp[i][j] = dp[i - 1][j] - dp[i - j - 2][j] * (1 - p) * p^(j + 1);
                    dp[i][j] -= dp[i - j - 2][j] * (1 - p) * pp[j + 1];
                }
            }
        }

        ans = 0.0;

        for (i = 1; i <= n; i++) {
            ans += i * (dp[n][i] - dp[n][i - 1]);
        }

        std::cout << std::fixed << std::setprecision(6) << ans << std::endl;
    }

	return 0;
}
