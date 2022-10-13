package main

import "strings"

func hardestWorker(n int, logs [][]int) int {
	ans := logs[0][0]
	temp := logs[0][1]
	for i := 1; i < len(logs); i++ {
		if temp < logs[i][1]-logs[i-1][1] {
			temp = logs[i][1] - logs[i-1][1]
			ans = logs[i][0]
		} else if temp == logs[i][1]-logs[i-1][1] {
			ans = min(ans, logs[i][0])
		}
	}
	return ans
}

func findArray(pref []int) []int {
	ans := make([]int, len(pref))
	ans[0] = pref[0]
	for i := 1; i < len(pref); i++ {
		ans[i] = pref[i-1] ^ pref[i]
	}
	return ans
}

func numberOfPaths(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]map[int]int, m+1)
	for i := range dp {
		dp[i] = make([]map[int]int, n+1)
	}
	dp[1][0] = map[int]int{0: 1}
	mod := int(1e9 + 7)
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			temp := make(map[int]int)
			for key, v := range dp[i-1][j] {
				temp[(key+grid[i-1][j-1])%k] = (v + temp[(key+grid[i-1][j-1])%k]) % mod
			}
			for key, v := range dp[i][j-1] {
				temp[(key+grid[i-1][j-1])%k] = (v + temp[(key+grid[i-1][j-1])%k]) % mod
			}
			dp[i][j] = temp
		}
	}
	// fmt.Println(dp)
	return dp[m][n][0]
}

func robotWithString(s string) string {
	//我们统计每一个字符 比其小的字符数目
	digits := make([]int, 26)
	for _, val := range s {
		digits[int(val-'a')]++
	}
	smaller := make([]int, 26)
	smaller[0] = 0
	for i := 1; i < 26; i++ {
		smaller[i] = smaller[i-1] + digits[i-1]
	}
	var res strings.Builder
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
		}
		for j := int(s[i] - 'a'); j < 26; j++ {
			smaller[j]--
		}
		for len(stack) > 0 && smaller[int(stack[len(stack)-1]-'a')] <= 0 {
			res.WriteByte(stack[len(stack)-1])
			stack = stack[0 : len(stack)-1]
		}
	}
	for len(stack) > 0 {
		res.WriteByte(stack[len(stack)-1])
		stack = stack[0 : len(stack)-1]
	}
	return res.String()
}
