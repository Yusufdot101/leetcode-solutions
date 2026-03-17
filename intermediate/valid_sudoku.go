package intermediate

import (
	"slices"
)

/**
You are given a 9 x 9 Sudoku board board. A Sudoku board is valid if the following rules are followed:
	1. Each row must contain the digits 1-9 without duplicates.
	2. Each column must contain the digits 1-9 without duplicates.
	3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without duplicates.
Return true if the Sudoku board is valid, otherwise return false

Note: A board does not need to be full or be solvable to be valid.

Example 1:
Input: board =
[["1","2",".",".","3",".",".",".","."],
 ["4",".",".","5",".",".",".",".","."],
 [".","9","8",".",".",".",".",".","3"],
 ["5",".",".",".","6",".",".",".","4"],
 [".",".",".","8",".","3",".",".","5"],
 ["7",".",".",".","2",".",".",".","6"],
 [".",".",".",".",".",".","2",".","."],
 [".",".",".","4","1","9",".",".","8"],
 [".",".",".",".","8",".",".","7","9"]]

Output: true

Example 2:
Input: board =
[["1","2",".",".","3",".",".",".","."],
 ["4",".",".","5",".",".",".",".","."],
 [".","9","1",".",".",".",".",".","3"],
 ["5",".",".",".","6",".",".",".","4"],
 [".",".",".","8",".","3",".",".","5"],
 ["7",".",".",".","2",".",".",".","6"],
 [".",".",".",".",".",".","2",".","."],
 [".",".",".","4","1","9",".",".","8"],
 [".",".",".",".","8",".",".","7","9"]]

Output: false
Explanation: There are two 1's in the top-left 3x3 sub-box.

Constraints:

    board.length == 9
    board[i].length == 9
    board[i][j] is a digit 1-9 or '.'.
**/

func IsValidSudoku(board [][]byte) bool {
	rows := make(map[int][]byte, len(board))
	columns := make(map[int][]byte, len(board))
	squares := make(map[[2]int][]byte, len(board))

	for r := range 9 {
		for c := range 9 {
			val := board[r][c]
			if val == '.' {
				continue
			}
			if slices.Contains(rows[r], val) {
				return false
			}
			if slices.Contains(columns[c], val) {
				return false
			}
			if slices.Contains(squares[[2]int{r / 3, c / 3}], val) {
				return false
			}
			rows[r] = append(rows[r], val)
			columns[c] = append(columns[c], val)
			squares[[2]int{r / 3, c / 3}] = append(squares[[2]int{r / 3, c / 3}], val)
		}
	}

	return true
}
