package main

func maximumXOR(nums []int) int {
	res := 0
	for _, val := range nums {
		res = res | val
	}
	return res
}

func checkXMatrix(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i+j == len(grid)-1 || i == j {
				if grid[i][j] == 0 {
					return false
				}
			} else {
				if grid[i][j] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func countHousePlacements(n int) int {
	mod := int(1e9 + 7)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[1][0] = 1
	dp[1][1] = 1
	for i := 2; i <= n; i++ {
		dp[i][0] = dp[i-1][1]
		dp[i][1] = (dp[i-1][0] + dp[i-1][1]) % mod
	}
	temp := (dp[n][0] + dp[n][1]) % mod
	return temp * temp % mod
}
