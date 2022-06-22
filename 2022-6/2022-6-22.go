package main

import "strings"

func soupServings(n int) float64 {
	if n == 0 {
		return float64(0.5)
	}
	if n >= 5000 {
		return float64(1)
	}
	if n%25 != 0 {
		n = n/25 + 1
	} else {
		n = n / 25
	}
	dp := make([][]float64, n+1)
	for i := range dp {
		dp[i] = make([]float64, n+1)
	}
	dp[0][0] = 0.5
	for i := 1; i <= n; i++ {
		dp[0][i] = 1
		dp[i][0] = 0
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = (dp[max(i-4, 0)][j] + dp[max(i-3, 0)][j-1] + dp[max(i-2, 0)][max(j-2, 0)] + dp[i-1][max(j-3, 0)]) / 4
		}
	}
	return dp[n][n]
}

func shiftingLetters(s string, shifts []int) string {
	sum := 0
	mod := 26
	for i := len(shifts) - 1; i >= 0; i-- {
		sum = (sum + shifts[i]) % mod
		shifts[i] = sum
	}
	//fmt.Println(shifts)
	var ans strings.Builder
	for i, val := range s {
		pos := (int(val-'a') + shifts[i]) % mod
		//fmt.Println(pos)
		ans.WriteRune(rune(pos + 'a'))
	}
	return ans.String()
}
