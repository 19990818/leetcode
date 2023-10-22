package main

import "strconv"

func maxSumAfterPartitioning(arr []int, k int) int {
	// dp[i]表示i以后能够得到的最大值
	n := len(arr)
	dp := make([]int, n+1)
	dp[n] = 0
	dp[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		ma := arr[i]
		for j := i; j < n && j < i+k; j++ {
			ma = max(ma, arr[j])
			dp[i] = max(dp[i], (j-i+1)*ma+dp[j+1])
		}
	}
	return dp[0]
}

func findColumnWidth(grid [][]int) []int {
	res := make([]int, len(grid[0]))
	for j := 0; j < len(grid[0]); j++ {
		for i := 0; i < len(grid); i++ {
			res[j] = max(res[j], len(strconv.Itoa(grid[i][j])))
		}
	}
	return res
}

func findPrefixScore(nums []int) []int64 {
	ma := 0
	for i, v := range nums {
		ma = max(ma, v)
		nums[i] = v + ma
	}
	res := make([]int64, len(nums))
	res[0] = int64(nums[0])
	for i := 1; i < len(nums); i++ {
		res[i] = int64(res[i-1]) + int64(nums[i])
	}
	return res
}
