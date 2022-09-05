package main

import "sort"

func findLongestChainDP(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[i][1]
	})
	//dp记录每个开始值开始的最长数队
	dp := make(map[int]int)
	n := len(pairs)
	dp[pairs[n-1][0]] = 1
	ans := 1
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if pairs[i][1] < pairs[j][0] {
				dp[pairs[i][0]] = max(dp[pairs[i][0]], dp[pairs[j][0]]+1)
			}
			ans = max(ans, dp[pairs[i][0]])
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 贪心
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	ans := 1
	end := pairs[0][1]
	for i := 2; i < len(pairs); i++ {
		if pairs[i][0] > end {
			ans++
			end = pairs[i][1]
		}
	}
	return ans
}
