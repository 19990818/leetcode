package main

import "sort"

func countPairs(nums []int, target int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				res++
			}
		}
	}
	return res
}

func canMakeSubsequence(str1 string, str2 string) bool {
	i, j := 0, 0
	for i < len(str1) && j < len(str2) {
		if str2[j]-str1[i] >= 0 && str2[j]-str1[i] < 2 || (str2[j] == 'a' && str1[i] == 'z') {
			j++
		}
		i++
	}
	return j >= len(str2)
}

func minimumOperations(nums []int) int {
	cnts := make([][]int, len(nums))
	for i := range cnts {
		cnts[i] = make([]int, 3)
	}
	for i, v := range nums {
		cnts[i][v-1]++
	}
	res := len(nums)
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			cnt1 := i + 1 - cnts[i][1]
			cnt2 := j - i + 1 - cnts[j][2] + cnts[i][2]
			cnt3 := len(nums) - j - cnts[len(nums)-1][3] + cnts[j][3]
			res = min(res, cnt1+cnt2+cnt3)
		}
	}
	return res
}

func isAcronym(words []string, s string) bool {
	res := ""
	for _, word := range words {
		res += string(word[0])
	}
	return res == s
}

func minimumSum(n int, k int) int {
	res := 0
	cnt := 0
	m := make(map[int]int)
	for i := 1; cnt < n; i++ {
		if i < k && m[k-i] == 1 {
			continue
		}
		res += i
		m[i] = 1
		cnt++
	}
	return res
}

func maximizeTheProfit(n int, offers [][]int) int {
	dp := make([]int, n+1)
	sort.Slice(offers, func(i, j int) bool {
		if offers[i][1] == offers[j][1] {
			return offers[i][0] < offers[j][0]
		}
		return offers[i][1] < offers[j][1]
	})
	k := 0
	for i := 0; i < n; i++ {
		dp[i+1] = max(dp[i], dp[i+1])
		for k < len(offers) && i == offers[k][1] {
			dp[i+1] = max(dp[i+1], dp[offers[k][0]]+offers[k][2])
			k++
		}
	}
	return dp[n]
}
