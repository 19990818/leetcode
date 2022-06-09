package main

func isMatch2(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(s) == 0 {
		for _, val := range p {
			if val != '*' {
				return false
			}
		}
		return true
	}
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for j := 1; j < len(p)+1; j++ {
		if p[j-1] == '*' {
			//i对应的字符不进行匹配
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '?' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				//其中dp[i][j-1]表示当前i对应的字符不进行匹配
				//dp[i-1][j]表示会对当前字符进行匹配
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	//printDp(dp)
	return dp[len(s)][len(p)]
}
