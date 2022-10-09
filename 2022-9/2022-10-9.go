package main

import (
	"math/bits"
)

func commonFactors(a int, b int) int {
	ans := 0
	for i := 1; i <= min(a, b); i++ {
		if a%i == 0 && b%i == 0 {
			ans++
		}
	}
	return ans
}

func maxSum(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			to := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}}
			temp := grid[i][j]
			for _, t := range to {
				temp += grid[i+t[0]][j+t[1]]
			}
			ans = max(ans, temp)
		}
	}
	return ans
}

func minimizeXor(num1 int, num2 int) int {
	cnt := bits.OnesCount(uint(num2))
	a := make([]int, 32)
	used := make(map[int]int)
	for i := 31; i >= 0; i-- {
		if num1 >= 1<<i {
			a[i] = 1
			num1 -= 1 << i
		}
	}
	for i := 31; i >= 0 && cnt > 0; i-- {
		if a[i] == 1 {
			used[i] = 1
			cnt--
		}
	}
	for i := 0; i < 32 && cnt > 0; i++ {
		if used[i] == 0 {
			cnt--
			used[i] = 1
		}
	}
	ans := 0
	for i := 0; i < 32; i++ {
		if used[i] == 1 {
			ans += 1 << i
		}
	}
	return ans
}
