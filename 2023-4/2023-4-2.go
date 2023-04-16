package main

import (
	"math"
)

func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[n-3][n-1] = values[n-3] * values[n-2] * values[n-1]
	for i := n - 3; i >= 0; i-- {
		for j := i + 2; j < n; j++ {
			dp[i][j] = math.MaxInt64
			for k := i + 1; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j]+values[k]*values[i]*values[j])
			}
		}
	}
	return dp[0][n-1]
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minNumber(nums1 []int, nums2 []int) int {
	m := make([]int, 10)
	m1, m2 := 9, 9
	for _, v := range nums1 {
		m[v]++
		m1 = min(m1, v)
	}
	for _, v := range nums2 {
		m[v]++
		m2 = min(m2, v)
	}
	for i := 1; i < 10; i++ {
		if m[i] == 2 {
			return i
		}
	}
	return min(m1, m2)*10 + max(m1, m2)
}

func maximumCostSubstring(s string, chars string, vals []int) int {
	costs := make([]int, 26)
	for i := range costs {
		costs[i] = i + 1
	}
	for i, v := range chars {
		costs[int(v-'a')] = vals[i]
	}
	dp := make([]int, len(s)+1)
	dp[0] = costs[int(s[0]-'a')]
	res := max(dp[0], 0)
	for i := 1; i < len(s); i++ {
		dp[i] = max(dp[i-1]+costs[int(s[i]-'a')], costs[int(s[i]-'a')])
		res = max(res, dp[i])
	}

	return res
}

func prevPermOpt1(arr []int) []int {
	h := make([]int, 0)
	h = append(h, len(arr)-1)
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] <= arr[h[len(h)-1]] {
			h = append(h, i)
			continue
		}
		pos := -1
		for len(h) > 0 && arr[i] > arr[h[len(h)-1]] {
			if pos == -1 || arr[pos] < arr[h[len(h)-1]] {
				pos = h[len(h)-1]
			}
			h = h[0 : len(h)-1]
		}
		arr[i], arr[pos] = arr[pos], arr[i]
		return arr
	}
	return arr
}



