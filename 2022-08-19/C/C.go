// https://codeforces.com/problemset/problem/754/B

// 754B - Ilya and tic-tac-toe game

// Ilya is an experienced player in tic-tac-toe on the 4 × 4 field. He always starts and plays with Xs. He played a lot of games today with his friend Arseny. The friends became tired and didn't finish the last game. It was Ilya's turn in the game when they left it. Determine whether Ilya could have won the game by making single turn or not.
// The rules of tic-tac-toe on the 4 × 4 field are as follows. Before the first turn all the field cells are empty. The two players take turns placing their signs into empty cells (the first player places Xs, the second player places Os). The player who places Xs goes first, the another one goes second. The winner is the player who first gets three of his signs in a row next to each other (horizontal, vertical or diagonal).

// Input
// The tic-tac-toe position is given in four lines.
// Each of these lines contains four characters. Each character is '.' (empty cell), 'x' (lowercase English letter x), or 'o' (lowercase English letter o). It is guaranteed that the position is reachable playing tic-tac-toe, and it is Ilya's turn now (in particular, it means that the game is not finished). It is possible that all the cells are empty, it means that the friends left without making single turn.

// Output
// Print single line: "YES" in case Ilya could have won by making single turn, and "NO" otherwise.

// Examples

// input
// xx..
// .oo.
// x...
// oox.
// output
// YES

// input
// x.ox
// ox..
// x.o.
// oo.x
// output
// NO

// input
// x..x
// ..oo
// o...
// x.xo
// output
// YES

// input
// o.x.
// o...
// .x..
// ooxx
// output
// NO

// Note
// In the first example Ilya had two winning moves: to the empty cell in the left column and to the leftmost empty cell in the first row.
// In the second example it wasn't possible to win by making single turn.
// In the third example Ilya could have won by placing X in the last row between two existing Xs.
// In the fourth example it wasn't possible to win by making single turn.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var writer = bufio.NewWriter(os.Stdout)
	var t, i, j int
	var field [4]string
	var ans bool

	// Reads field until EOF
	for t, _ = fmt.Fscan(reader, &field[0], &field[1], &field[2], &field[3]); t == 4; t, _ = fmt.Fscan(reader, &field[0], &field[1], &field[2], &field[3]) {
		ans = false

		for i = 0; i < 4 && !ans; i++ {
			for j = 0; j < 4 && !ans; j++ {
				if field[i][j] == '.' {
					if (i >= 2 && field[i-2][j] == 'x' && field[i-1][j] == 'x') ||
						(i >= 1 && field[i-1][j] == 'x' && i <= 2 && field[i+1][j] == 'x') ||
						(i <= 1 && field[i+1][j] == 'x' && field[i+2][j] == 'x') ||
						(j >= 2 && field[i][j-2] == 'x' && field[i][j-1] == 'x') ||
						(j >= 1 && field[i][j-1] == 'x' && j <= 2 && field[i][j+1] == 'x') ||
						(j <= 1 && field[i][j+1] == 'x' && field[i][j+2] == 'x') ||
						(i >= 2 && j >= 2 && field[i-2][j-2] == 'x' && field[i-1][j-1] == 'x') ||
						(i >= 1 && j >= 1 && field[i-1][j-1] == 'x' && i <= 2 && j <= 2 && field[i+1][j+1] == 'x') ||
						(i <= 1 && j <= 1 && field[i+1][j+1] == 'x' && field[i+2][j+2] == 'x') ||
						(i >= 2 && j <= 1 && field[i-2][j+2] == 'x' && field[i-1][j+1] == 'x') ||
						(i >= 1 && j <= 2 && field[i-1][j+1] == 'x' && i <= 2 && j >= 1 && field[i+1][j-1] == 'x') ||
						(i <= 1 && j >= 2 && field[i+1][j-1] == 'x' && field[i+2][j-2] == 'x') {
						ans = true
					}
				}
			}
		}

		if ans {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}
	}

	writer.Flush()
}
