package main

import "math"

func countPrefixes(words []string, s string) int {
	ans := 0
	for _, val := range words {
		if len(val) <= len(s) && val == s[0:len(val)] {
			ans++
		}
	}
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func minimumAverageDifference(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	minNum := math.MaxInt64
	left := 0
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		left += nums[i]
		sum -= nums[i]
		var cur int
		if i == n-1 {
			cur = abs(left / (i + 1))
		} else {
			cur = abs(left/(i+1) - sum/(n-i-1))
		}

		if cur < minNum {
			ans = i
			minNum = cur
		}
	}
	return ans
}

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	source := make([][]int, m)
	var dfs func(point []int, src [][]int)
	dfs = func(point []int, src [][]int) {
		x, y := point[0], point[1]
		for i := x - 1; i >= 0; i-- {
			if src[i][y] == 2 {
				break
			}
			src[i][y] = 1
		}
		for i := x + 1; i < len(src); i++ {
			if src[i][y] == 2 {
				break
			}
			src[i][y] = 1
		}
		for j := y - 1; j >= 0; j-- {
			if src[x][j] == 2 {
				break
			}
			src[x][j] = 1
		}
		for j := y + 1; j < len(src[0]); j++ {
			if src[x][j] == 2 {
				break
			}
			src[x][j] = 1
		}
	}
	for i := 0; i < m; i++ {
		source[i] = make([]int, n)
	}
	for _, val := range guards {
		source[val[0]][val[1]] = 2
	}
	for _, val := range walls {
		source[val[0]][val[1]] = 2
	}
	for _, val := range guards {
		dfs(val, source)
	}
	ans := 0
	for _, val := range source {
		for _, val2 := range val {
			if val2 == 0 {
				ans++
			}
		}
	}
	return ans
}

func findTheWinner(n int, k int) int {
	//从一个人出发 最后剩下的位置必然是0
	//那么之前一轮淘汰的人为k-1，也就是上一轮结束 开始位置为k
	//那么前一轮最后剩下的位置为(curpos+k)%curlen
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (ans + k) % i
	}
	return ans + 1
}
