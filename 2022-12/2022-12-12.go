package main

func beautySum(s string) int {
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, 26)
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < 26; j++ {
			dp[i+1][j] = dp[i][j]
		}
		dp[i+1][int(s[i]-'a')]++
	}
	return cntRes(len(s), dp)
}
func cntRes(n int, dp [][]int) int {
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			maxNum, minNum := 0, n
			for k := 0; k < 26; k++ {
				maxNum = max(maxNum, dp[j+1][k]-dp[i][k])
				if dp[j+1][k]-dp[i][k] != 0 {
					minNum = min(minNum, dp[j+1][k]-dp[i][k])
				}
			}
			if maxNum > minNum {
				res += maxNum - minNum
			}
		}
	}
	return res
}
