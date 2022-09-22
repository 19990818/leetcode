package main

func countSubstrings(s string, t string) int {
	//数据量为100 属于比较小的 动态规划
	//其中含有两个字符串
	//需要得到两个子字符串有多少个连续的相同数量
	var cnt func(s1, s2 string) int
	cnt = func(s1, s2 string) int {
		if len(s1) == 0 || len(s2) == 0 {
			return 0
		}
		dp := make([][]int, len(s1))
		for i := range dp {
			dp[i] = make([]int, len(s2))
		}
		for i := 0; i < len(s1); i++ {
			if s1[i] == s2[0] {
				dp[i][0] = 1
			}
		}
		for j := 0; j < len(s2); j++ {
			if s1[0] == s2[j] {
				dp[0][j] = 1
			}
		}
		for i := 1; i < len(s1); i++ {
			for j := 1; j < len(s2); j++ {
				if s1[i] == s2[j] {
					dp[i][j] = 1 + dp[i-1][j-1]
				} else {
					dp[i][j] = 0
				}
			}
		}
		return dp[len(s1)-1][len(s2)-1]
	}
	reverseStr := func(s string) string {
		res := make([]byte, len(s))
		for i := len(s) - 1; i >= 0; i-- {
			res[len(s)-i-1] = s[i]
		}
		return string(res)
	}
	ans := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] != t[j] {
				left, right := cnt(s[0:i], t[0:j]), cnt(reverseStr(s[i+1:]), reverseStr(t[j+1:]))
				//fmt.Println(i,j,left,right)
				ans += (left + 1) * (right + 1)
			}
		}
	}
	return ans
}
