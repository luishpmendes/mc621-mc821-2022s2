// https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=12&page=show_problem&problem=951

// 10010 - Where's Waldorf?
// Given a m by n grid of letters, (1 ≤ m, n ≤ 50), and a list of words, find the location in the grid at
// which the word can be found.
// A word matches a straight, uninterrupted line of letters in the grid. A word can match the letters
// in the grid regardless of case (i.e. upper and lower case letters are to be treated as the same). The
// matching can be done in any of the eight directions either horizontally, vertically or diagonally through
// the grid.

// Input
// The input begins with a single positive integer on a line by itself indicating the number
// of the cases following, each of them as described below. This line is followed by a blank
// line, and there is also a blank line between two consecutive inputs.
// The input begins with a pair of integers, m followed by n, 1 ≤ m, n ≤ 50 in decimal notation on a
// single line. The next m lines contain n letters each; this is the grid of letters in which the words of the
// list must be found. The letters in the grid may be in upper or lower case. Following the grid of letters,
// another integer k appears on a line by itself (1 ≤ k ≤ 20). The next k lines of input contain the list of
// words to search for, one word per line. These words may contain upper and lower case letters only (no
// spaces, hyphens or other non-alphabetic characters).

// Output
// For each test case, the output must follow the description below. The outputs of two
// consecutive cases will be separated by a blank line.
// For each word in the word list, a pair of integers representing the location of the corresponding
// word in the grid must be output. The integers must be separated by a single space. The first integer
// is the line in the grid where the first letter of the given word can be found (1 represents the topmost
// line in the grid, and m represents the bottommost line). The second integer is the column in the grid
// where the first letter of the given word can be found (1 represents the leftmost column in the grid, and
// n represents the rightmost column in the grid). If a word can be found more than once in the grid,
// then the location which is output should correspond to the uppermost occurence of the word (i.e. the
// occurence which places the first letter of the word closest to the top of the grid). If two or more words
// are uppermost, the output should correspond to the leftmost of these occurences. All words can be
// found at least once in the grid.

// Sample Input
// 1
//
// 8 11
// abcDEFGhigg
// hEbkWalDork
// FtyAwaldORm
// FtsimrLqsrc
// byoArBeDeyv
// Klcbqwikomk
// strEBGadhrb
// yUiqlxcnBjf
// 4
// Waldorf
// Bambi
// Betty
// Dagbert

// Sample Output
// 2 5
// 2 3
// 1 2
// 7 8

#include <ctype.h>
#include <iostream>
#include <string>
#include <vector>

int x[] = {-1, -1, -1,  0, 0,  1, 1, 1};
int y[] = {-1,  0,  1, -1, 1, -1, 0, 1};

bool match(const std::vector<std::string> & grid, const std::string & word, int i, int j, int di, int dj) {
    bool result = true;
    int k;

    for (k = 0; i < (int) grid.size() && j < (int) grid.front().size() && i >= 0 && j >= 0 && k < (int) word.size() && result; i += di, j += dj, k++) {
        if (tolower(grid[i][j]) != tolower(word[k])) {
            result = false;
        }
    }

    if (k != (int) word.size()) {
        result = false;
    }

    return result;
}

int main() {
    int t, m, n, k, u, i, j, l;
    std::vector<std::string> grid;
    std::string word;
    bool found;

    while (std::cin >> t) {
        for (u = 0; u < t; u++) {
            std::cin >> m >> n;
            grid.resize(m);

            for (auto& s : grid) {
                std::cin >> s;
            }

            std::cin >> k;

            while (k--) {
                std::cin >> word;

                found = false;

                for (i = 0; i < m && !found; i++) {
                    for (j = 0; j < n && !found; j++) {
                        if (toupper(grid[i][j]) == toupper(word[0])) {
                            for (l = 0; l < 8 && !found; l++) {
                                if (match(grid, word, i, j, x[l], y[l])) {
                                    std::cout << i + 1 << " " << j + 1 << std::endl;
                                    found = true;
                                }
                            }
                        }
                    }
                }
            }

            if (u < t - 1) {
                std::cout << std::endl;
            }
        }
    }

	return 0;
}
