package piscine

import "github.com/01-edu/z01"

const n = 8

func EightQueens() {
	board := make([]int, n)
	solveQueens(board, 0)
}

func solveQueens(board []int, row int) {
	if row == n {
		printSolution(board)
		return
	}

	for col := 0; col < n; col++ {
		if isSafe(board, row, col) {
			board[row] = col
			solveQueens(board, row+1)
		}
	}
}

func isSafe(board []int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col || board[i]+i == col+row || board[i]-i == col-row {
			return false
		}
	}
	return true
}

func printSolution(board []int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i] == j {
				z01.PrintRune(rune(j + 49))
			}
		}
	}
	z01.PrintRune('\n')
}
