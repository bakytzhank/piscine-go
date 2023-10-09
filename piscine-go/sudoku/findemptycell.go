package sudoku

//function searches and returns row and column of the first empty cell found in sudoku table
func FindEmptyCell(table [9][9]int) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == 0 {
				//if found return i and j as row and column of empty cell
				return i, j
			}
		}
	}
	//if no empty cell return -1 for both row and column
	return -1, -1
}