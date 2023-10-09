package sudoku

//recursive function to find solution to given sudoku table
func FindSolution(table [9][9]int) bool {
	//assign row and col variables that FindEmptyCell returns
	row, col := FindEmptyCell(table)
	
	//solved sudoku is printed if empty cells not found
	if row == -1 && col == -1 {
		PrintSudoku(table)
		return true
	}
	//trying to place digits into empty cell
	for n := 1; n <= 9; n++ {
		//function DigitsCheck is called to check if digit is safe to use in this cell
		if DigitsCheck(table, row, col, n) {
			//if digit is Ok to use assign it to an empty cell and call FindSolution again with updated table
			table[row][col] = n
			if FindSolution(table) {
				return true
			}
			//backtracking if this solution does not work
			table[row][col] = 0
		} 
	}
	return false
}
