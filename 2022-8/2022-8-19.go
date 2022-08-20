package main

func validTicTacToe(board []string) bool {
	//不能有两组三个相等的 xcnt>=ocnt
	xcnt, ocnt := 0, 0
	win := 0
	m := make([][]byte, 3)
	for i := range m {
		m[i] = make([]byte, 3)
	}

	for i, val := range board {
		for j, val2 := range val {
			if val2 == 'O' {
				ocnt++
			}
			if val2 == 'X' {
				xcnt++
			}
			m[i][j] = byte(val2)
		}
	}
	if !(xcnt == ocnt || xcnt == ocnt+1) {
		return false
	}
	colM, rowM := make(map[int]int), make(map[int]int)
	var isStatisfycol func(i int) bool
	isStatisfycol = func(i int) bool {
		flag := m[i][0] == m[i][1] && m[i][1] == m[i][2] && m[i][0] != ' '
		if flag {
			if m[i][0] == 'X' {
				rowM[i] = 1
			} else {
				rowM[i] = 2
			}

		}
		return flag
	}
	var isStatisfyrow func(j int) bool
	isStatisfyrow = func(j int) bool {
		flag := m[0][j] == m[1][j] && m[1][j] == m[2][j] && m[0][j] != ' '
		if flag {
			if m[0][j] == 'X' {
				colM[j] = 1
			} else {
				colM[j] = 2
			}

		}

		return flag
	}
	var isStatisfy func(i, j int) bool
	isStatisfy = func(i, j int) bool {
		if i == 1 && j == 1 {
			flag := (m[i-1][j-1] == m[i][j] && m[i+1][j+1] == m[i][j] || m[i+1][j-1] == m[i][j] && m[i][j] == m[i-1][j+1]) && m[i][j] != ' '
			if flag {
				if m[i][j] == 'X' {
					win = 1
				} else {
					win = 2
				}
			}

			return flag
		}
		return false
	}
	sum := 0
	for i := range m {
		for j := range m[i] {
			flag := false
			if colM[j] != 0 || rowM[i] != 0 {
				if colM[j] != 0 {
					win = colM[j]
				} else {
					win = rowM[i]
				}
				sum -= 1
			}
			flag = isStatisfyrow(j) || flag
			flag = isStatisfycol(i) || flag

			flag = flag || isStatisfy(i, j)
			//fmt.Println(i,j,flag)
			if flag {
				sum += 1
			}
		}
	}
	//fmt.Println(colM,rowM,win)
	if sum > 1 {
		return false
	}
	if win == 1 && xcnt == ocnt {
		return false
	}
	if win == 2 && xcnt == ocnt+1 {
		return false
	}
	return true
}
