package sudoku

func DigitsCheck(table [9][9]int, row int, col int, n int) bool {
	//row & col coming through FindSolution func
	
	//row check
	for i := 0; i < 9; i++ {
		if table[row][i] == n{
			return false
		}
	}

	//column check
	for j := 0; j < 9; j++{
		if table[j][col] == n{
			return false
		}
	}

	//3x3 minitable check
	startRow := row-(row%3) //find start point by rows in 3x3
	startCol := col-(col%3) //find start point by columns in 3x3
	for i:= startRow; i < startRow+3; i++ {
		for j:= startCol; j < startCol+3; j++{
			if table[i][j] == n {
				return false
			}
		}
	}

	return true
}
