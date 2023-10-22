package main

import (
	"sort"
	"strings"
)

func isValid(s string) bool {
	pre := s
	for strings.ReplaceAll(pre, "abc", "") != pre {
		pre = strings.ReplaceAll(pre, "abc", "")
	}
	return pre == ""
}

func maximizeSum(nums []int, k int) int {
	sort.Ints(nums)
	return k*nums[len(nums)-1] + k*(k-1)/2
}

func findThePrefixCommonArray(A []int, B []int) []int {
	res := make([]int, len(A))
	ma, mb := make(map[int]int), make(map[int]int)
	for i := range A {
		ma[A[i]] = 1
		mb[B[i]] = 1
		res[i] = cntsuffix(ma, mb)
	}
	return res
}
func cntsuffix(ma, mb map[int]int) int {
	cnt := 0
	for k := range ma {
		if mb[k] == 1 {
			cnt++
		}
	}
	return cnt
}

func findMaxFish(grid [][]int) int {
	res := 0
	travel := make(map[int]int)
	for i, g := range grid {
		for j, v := range g {
			if v > 0 && travel[i*len(grid[0])+j] == 0 {
				travel[i*len(grid[0])+j] = 1
				res = max(res, bfs(i, j, grid, travel))
			}
		}
	}
	return res
}
func bfs(i, j int, grid [][]int, travel map[int]int) int {
	tos := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	q := [][]int{{i, j}}
	res := 0
	for len(q) > 0 {
		temp := q
		q = nil
		for i := range temp {
			res += grid[temp[i][0]][temp[i][1]]
			for _, to := range tos {
				curx, cury := temp[i][0]+to[0], temp[i][1]+to[1]
				if check(curx, cury, len(grid), len(grid[0]), travel, grid) {
					travel[curx*len(grid[0])+cury] = 1
					q = append(q, []int{curx, cury})
				}
			}
		}
	}
	return res
}
func check(x, y, m, n int, travel map[int]int, grid [][]int) bool {
	return x >= 0 && x < m && y >= 0 && y < n && travel[x*n+y] == 0 && grid[x][y] > 0
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
