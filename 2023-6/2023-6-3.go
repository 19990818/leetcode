package main

func maxRepOpt1(text string) int {
	m := make(map[byte][]int)
	for i := range text {
		if m[text[i]] == nil {
			m[text[i]] = []int{i}
		} else {
			m[text[i]] = append(m[text[i]], i)
		}
	}
	res := 1
	for _, v := range m {
		dp := make([][]int, len(v))
		for j := range dp {
			dp[j] = make([]int, 2)
		}
		dp[0][0], dp[0][1] = 1, 1
		for i := 1; i < len(v); i++ {
			if v[i] == v[i-1]+1 {
				dp[i][0] = dp[i-1][0] + 1
				dp[i][1] = dp[i-1][1] + 1
			} else if v[i] == v[i-1]+2 {
				dp[i][0] = 1
				dp[i][1] = dp[i-1][0] + 2
				res = max(res, min(max(dp[i-1][0]+1, dp[i-1][1]), len(v)))
			} else {
				res = max(res, min(max(dp[i-1][0]+1, dp[i-1][1]), len(v)))
				dp[i][0] = 1
				dp[i][1] = 1
			}
		}
		res = max(res, min(max(dp[len(v)-1][0]+1, dp[len(v)-1][1]), len(v)))
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
