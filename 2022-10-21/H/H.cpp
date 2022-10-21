// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=18&page=show_problem&problem=1576

// 10635 - Prince and Princess
// In an n×n chessboard, Prince and Princess plays a game. The squares in the chessboard are numbered
// 1, 2, 3, . . . , n ∗ n, as shown below:
// Prince stands in square 1, make p jumps and finally reach square n ∗ n. He enters a square at most
// once. So if we use xp to denote the p-th square he enters, then x1, x2, …, xp+1 are all different. Note
// that x1 = 1 and xp+1 = n ∗ n. Princess does the similar thing – stands in square 1, make q jumps and
// finally reach square n ∗ n. We use y1, y2, …, yq+1 to denote the sequence, and all q + 1 numbers are
// different.
// Figure 2 belows show a 3×3 square, a possible route for Prince and a different route for Princess.
// The Prince moves along the sequence: 1 –> 7 –> 5 –> 4 –> 8 –> 3 –> 9 (Black arrows), while the
// Princess moves along this sequence: 1 –> 4 –> 3 –> 5 > 6 –> 2 –> 8 –> 9 (White arrow).
// The King – their father, has just come. “Why move separately? You are brother and sister!” said
// the King, “Ignore some jumps and make sure that you’re always together.”
// For example, if the Prince ignores his 2nd, 3rd, 6th jump, he’ll follow the route: 1 –> 4 –> 8 –>
// 9. If the Princess ignores her 3rd, 4th, 5th, 6th jump, she’ll follow the same route: 1 –> 4 –> 8 –>
// 9, (The common route is shown in figure 3) thus satisfies the King, shown above. The King wants to
// know the longest route they can move together, could you tell him?

// Input
// The first line of the input contains a single integer t (1 ≤ t ≤ 10), the number of test cases followed.
// For each case, the first line contains three integers n, p, q (2 ≤ n ≤ 250, 1 ≤ p, q < n ∗ n). The second
// line contains p + 1 different integers in the range [1 . . . n ∗ n], the sequence of the Prince. The third line
// contains q + 1 different integers in the range [1 . . . n ∗ n], the sequence of the Princess.

// Output
// For each test case, print the case number and the length of longest route. Look at the output for sample
// input for details.

// Sample Input
// 1
// 3 6 7
// 1 7 5 4 8 3 9
// 1 4 3 5 6 2 8 9

// Sample Output
// Case 1: 4

// find LCS of two permutations

#include <iostream>
#include <map>
#include <vector>

int main() {
    unsigned t, n, p, q, c, value, position, i;
    std::vector<unsigned> lcs;
    std::map<unsigned, unsigned> val_to_pos;
    std::vector<unsigned>::iterator it;

    while (std::cin >> t) {
        for (c = 1; c <= t; c++) {
            std::cin >> n >> p >> q;
            lcs.clear();
            val_to_pos.clear();

            if (p >= q) {
                lcs.reserve(p);
            } else {
                lcs.reserve(q);
            }

            for (position = 0; position <= p; position++) {
                std::cin >> value;
                val_to_pos[value] = position;
            }

            for (i = 0; i <= q; i++) {
                std::cin >> value;
                position = val_to_pos[value];
                
                if (position > 0) {
                    it = std::lower_bound(lcs.begin(), lcs.end(), position);

                    if (it == lcs.end()) {
                        lcs.push_back(position);
                    } else {
                        *it = position;
                    }
                }
            }

            std::cout << "Case " << c << ": " << lcs.size() + 1 << std::endl;
        }
    }

	return 0;
}
