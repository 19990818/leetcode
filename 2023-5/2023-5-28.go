package main

import "strings"

func removeTrailingZeros(num string) string {
	return strings.TrimRight(num, "0")
}

func differenceOfDistinctValues(grid [][]int) [][]int {
	res := make([][]int, len(grid))
	for i := range grid {
		res[i] = make([]int, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			res[i][j] = abs(cntTp(i, j, grid) - cntBr(i, j, grid))
		}
	}
	return res
}
func cntTp(i, j int, grid [][]int) int {
	m := make(map[int]int)
	for x := 1; x <= i && x <= j; x++ {
		m[grid[i-x][j-x]] = 1
	}
	return len(m)
}
func cntBr(i, j int, grid [][]int) int {
	m := make(map[int]int)
	for x := 1; x < len(grid)-i && x < len(grid[0])-j; x++ {
		m[grid[i+x][j+x]] = 1
	}
	return len(m)
}

func minimumCost(s string) int64 {
	res := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			continue
		}
		res += min(i, len(s)-i)
	}
	return int64(res)
}
