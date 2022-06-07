package main

func solveSudoku(board [][]byte) {
	//回溯 dfs主要问题是判断square的时候 开始开始结束位置首先写错了 修改只修改了部分 导致debug了很长时间
	solveSudoSuccess(board)
}
func isSatisfyCol(board [][]byte, i int, m map[byte]int) map[byte]int {
	for j := 0; j < 9; j++ {
		if board[i][j] != '.' {
			delete(m, board[i][j])
		}
	}
	return m
}
func isSatisfyRow(board [][]byte, j int, m map[byte]int) map[byte]int {
	for i := 0; i < 9; i++ {
		if board[i][j] != '.' {
			delete(m, board[i][j])
		}
	}
	return m
}
func isSatisfySquare(board [][]byte, i, j int, m map[byte]int) map[byte]int {
	starti, startj := i/3, j/3
	for i2 := starti * 3; i2 < starti*3+3; i2++ {
		for j2 := startj * 3; j2 < startj*3+3; j2++ {
			if board[i2][j2] != '.' {
				delete(m, board[i2][j2])
			}
		}
	}
	return m
}
func solveSudoSuccess(board [][]byte) bool {
	for i := range board {
		for j := range board[0] {
			if board[i][j] == '.' {
				tempM := make(map[byte]int)
				for k := 1; k <= 9; k++ {
					tempM[byte(k+'0')] = 1
				}
				tempM = isSatisfyCol(board, i, tempM)
				tempM = isSatisfyRow(board, j, tempM)
				tempM = isSatisfySquare(board, i, j, tempM)

				if len(tempM) == 0 {
					return false
				}
				for key := range tempM {
					board[i][j] = key
					//fmt.Println(int(key-'0'))
					if solveSudoSuccess(board) {
						return true
					}
					board[i][j] = '.'
				}
				return false
			}
		}
	}
	return true
}

func minEatingSpeed(piles []int, h int) int {
	//有上下边界，且存在可以选择哪部分的条件，使用二分查找能够快速找到符合条件的答案
	maxNum := 0
	for _, val := range piles {
		maxNum = max(maxNum, val)
	}
	left, right := 1, maxNum
	var countHours func(speed int) int
	countHours = func(speed int) int {
		res := 0
		for _, val := range piles {
			res += val / speed
			if val%speed > 0 {
				res += 1
			}
		}
		return res
	}
	for left < right {
		mid := (right-left)>>1 + left
		if countHours(mid) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func trap(height []int) int {
	//将极大值确定下来，那么两边的的值不大于极大值
	//可以确定下来一边，那么一边围成的面积实际上就是雨水
	//关键是围成面积，在这个围成面积的情况下，实际上就是两条边的抉择
	//再解决这样的问题下，尽可能减少变量，实际上也就是需要将一条边固定下来
	//在此题情况下，可以找到最长的边进行固定，将问题分为两部分
	//固定最长的边有一个好处，就是可以只用考虑一条边，木桶效应，面积由最短的边
	//决定，因此可以遍历完成面积的计算
	maxHeight, maxIndex := height[0], 0
	for i, val := range height {
		if val > maxHeight {
			maxIndex = i
			maxHeight = val
		}
	}
	temp := height[0]
	ans := 0
	for i := 0; i < maxIndex; i++ {
		if height[i] > temp {
			temp = height[i]
		} else {
			ans += temp - height[i]
		}
	}
	temp = height[len(height)-1]
	for i := len(height) - 1; i > maxIndex; i-- {
		if height[i] > temp {
			temp = height[i]
		} else {
			ans += temp - height[i]
		}
	}
	return ans
}
