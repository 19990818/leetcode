package main

import "bytes"

func new21Game(n int, k int, maxPts int) float64 {
	low := min(n, k+maxPts-1)
	high := max(n, k+maxPts-1)
	dp := make([]float64, high+1)
	for x := k; x <= low; x++ {
		dp[x] = 1
	}
	for x := low + 1; x < high; x++ {
		dp[x] = 0
	}
	for i := k - 1; i >= 0; i-- {
		if i == k-1 {
			for j := 1; j <= maxPts; j++ {
				dp[i] += dp[i+j] / float64(maxPts)
			}
		} else {
			dp[i] = dp[i+1] + (dp[i+1]-dp[i+1+maxPts])/float64(maxPts)
		}
	}
	return dp[0]
}

func pushDominoes(dominoes string) string {
	stack := make([]byte, 0)
	ans := make([]byte, 0)
	dominoes += "R"
	for i := range dominoes {
		if dominoes[i] == '.' {
			stack = append(stack, dominoes[i])
		} else if dominoes[i] == 'R' {
			if len(stack) > 0 {
				ans = append(ans, bytes.Repeat([]byte{stack[0]}, len(stack))...)
			}
			stack = []byte{}
			stack = append(stack, dominoes[i])
		} else {
			if len(stack) == 0 {
				ans = append(ans, 'L')
			} else if stack[0] == '.' {
				ans = append(ans, bytes.Repeat([]byte{'L'}, len(stack)+1)...)
				stack = []byte{}
			} else {
				ans = append(ans, bytes.Repeat([]byte{'R'}, (len(stack)+1)/2)...)
				if (len(stack)+1)%2 == 1 {
					ans = append(ans, '.')
				}
				ans = append(ans, bytes.Repeat([]byte{'L'}, (len(stack)+1)/2)...)
				stack = []byte{}
			}
		}
	}
	return string(ans)
}

func numMagicSquaresInside(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m < 3 || n < 3 {
		return 0
	}
	var isTrue func(i, j int) bool
	isTrue = func(i, j int) bool {
		target := grid[i][j] + grid[i][j+1] + grid[i][j+2]
		m := make(map[int]int)
		for y := j; y < j+3; y++ {
			temp := 0
			for x := i; x < i+3; x++ {
				temp += grid[x][y]
				m[grid[x][y]] = 1
			}
			if temp != target {
				return false
			}
		}
		for i := 1; i <= 9; i++ {
			if m[i] == 0 {
				return false
			}
		}
		for x := i + 1; x < i+3; x++ {
			temp := 0
			for y := j; y < j+3; y++ {
				temp += grid[x][y]
			}
			if temp != target {
				return false
			}
		}
		return (grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] == target) && (grid[i+2][j]+grid[i+1][j+1]+grid[i][j+2] == target)
	}
	ans := 0
	for x := 0; x <= m-3; x++ {
		for y := 0; y <= n-3; y++ {
			if isTrue(x, y) {
				ans++
			}
		}
	}
	return ans
}
