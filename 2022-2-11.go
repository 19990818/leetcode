package main

func gameOfLife(board [][]int) {
	neighber := []int{-1, 0, 1}
	rows, cols := len(board), len(board[0])
	copyboard := make([][]int, rows)
	for i := 0; i < rows; i++ {
		copyboard[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			copyboard[i][j] = board[i][j]
		}
	}
	// fmt.Println(copyboard)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			liveNeighber := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if neighber[i] != 0 || neighber[j] != 0 {
						r, c := row+neighber[i], col+neighber[j]
						if (r >= 0 && r < rows) && (c >= 0 && c < cols) && copyboard[r][c] == 1 {
							liveNeighber += 1
						}
					}
				}
			}
			//fmt.Println(liveNeighber)
			if copyboard[row][col] == 1 && (liveNeighber < 2 || liveNeighber > 3) {
				board[row][col] = 0
			}
			if copyboard[row][col] == 0 && liveNeighber == 3 {
				board[row][col] = 1
			}
		}
	}
	//fmt.Println(copyboard)
}
