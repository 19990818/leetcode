package main

import "sort"

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	m := make(map[float64]int)
	for i := 0; i < len(nums)/2; i++ {
		temp := float64(nums[i]+nums[len(nums)-1-i]) / float64(2)
		m[temp] = 1
	}
	return len(m)
}

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1
	mod := int(1e9 + 7)
	for i := 1; i <= high; i++ {
		if i-zero >= 0 {
			dp[i] = (dp[i] + dp[i-zero]) % mod
		}
		if i-one >= 0 {
			dp[i] = (dp[i-one] + dp[i]) % mod
		}
	}
	res := 0
	for i := low; i <= high; i++ {
		res = (res + dp[i]) % mod
	}
	return res
}
