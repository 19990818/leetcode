package main

import "strconv"

func numberOfCuts(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return n / 2
	}
	return n
}

func onesMinusZeros(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	oR, oC := make([]int, m), make([]int, n)
	zR, zC := make([]int, m), make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				oR[i]++
				oC[j]++
			} else {
				zR[i]++
				zC[j]++
			}
		}
	}
	diff := make([][]int, m)
	for i := range diff {
		diff[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diff[i][j] = oR[i] + oC[j] - zR[i] - zC[j]
		}
	}
	return diff
}

func bestClosingTime(customers string) int {
	sum := 0
	for _, val := range customers {
		if val == 'Y' {
			sum++
		}
	}
	res := 0
	minNum := sum
	for i := 1; i <= len(customers); i++ {
		if customers[i-1] == 'Y' {
			sum--
		} else {
			sum++
		}
		if sum < minNum {
			res = i
			minNum = sum
		}
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countPalindromes(s string) int {
	mod := int(1e9 + 7)
	ps := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		str := strconv.Itoa(i)
		if len(str) == 1 {
			str = "00" + str
		}
		if len(str) == 2 {
			str = "0" + str
		}
		str = str + string(str[1]) + string(str[0])
		ps[i] = str
	}
	res := 0
	for _, v := range ps {
		res = (res + numDistinct(s, v)) % mod
	}
	return res
}
func reverseStr(s []byte) string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

// 找到子序列的个数
func numDistinct(s, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	mod := int(1e9 + 7)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][n] = 1
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = (dp[i+1][j+1] + dp[i+1][j]) % mod
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}
