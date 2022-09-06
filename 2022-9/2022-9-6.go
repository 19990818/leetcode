package main

import "math/bits"

func findSubarrays(nums []int) bool {
	m := make(map[int]int)
	for i := 1; i < len(nums); i++ {
		if m[nums[i]+nums[i-1]] == 1 {
			return true
		}
		m[nums[i]+nums[i-1]] = 1
	}
	return false
}

// 需要形成用二进制数表示状态的思路
func maximumRows(matrix [][]int, numSelect int) int {
	m := len(matrix[0])
	n := len(matrix)
	maxX := 1 << m
	ans := 0
	for x := 0; x < maxX; x++ {
		cnt := bits.OnesCount(uint(x))
		temp := 0
		if cnt == numSelect {
			for i := 0; i < n; i++ {
				flag := true
				for j := 0; j < m && flag; j++ {
					if matrix[i][j] == 1 && (x>>j&1) == 0 {
						flag = false
					}
				}
				if flag {
					temp++
				}
			}
		}
		ans = max(ans, temp)
	}
	return ans
}
