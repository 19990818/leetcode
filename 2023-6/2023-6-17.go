package main

import (
	"sort"
)

func sumDistance(nums []int, s string, d int) int {
	for i := range nums {
		if s[i] == 'R' {
			nums[i] += d
		} else {
			nums[i] -= d
		}
	}
	sort.Ints(nums)

	preSum := 0
	res := 0
	mod := int(1e9 + 7)
	for i, v := range nums {
		res = (res + i*v - preSum) % mod
		preSum = (preSum + v) % mod
	}
	return res
}

func buyChoco(prices []int, money int) int {
	sort.Ints(prices)
	t := prices[0] + prices[1]
	if t > money {
		return money
	}
	return money - t
}

func minExtraChar(s string, dictionary []string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
		dp[i][i] = 0
		for j := i + 1; j < len(s)+1; j++ {
			dp[i][j] = 500
		}
	}
	m := make(map[string]int)
	for _, v := range dictionary {
		m[v] = 1
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s)+1; j++ {
			for k := i; k < j; k++ {
				if m[s[k:j]] == 1 {
					dp[i][j] = min(dp[i][j], dp[i][k])
				} else {
					dp[i][j] = min(dp[i][j], dp[i][k]+j-k)
				}
			}
		}
	}
	return dp[0][len(s)]
}
