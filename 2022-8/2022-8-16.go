package main

func setZeroes(matrix [][]int) {
	x, y := make(map[int]int), make(map[int]int)
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				x[i] = 1
				y[j] = 1
			}
		}
	}
	for k := range x {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[k][j] = 0
		}
	}
	for k := range y {
		for i := 0; i < len(matrix); i++ {
			matrix[i][k] = 0
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
