package main

import "strings"

func maximumNumberOfStringPairs(words []string) int {
	res := 0
	m := make(map[string]int)
	for _, v := range words {
		var temp strings.Builder
		for i := len(v) - 1; i >= 0; i-- {
			temp.WriteByte(v[i])
		}
		if m[temp.String()] == 1 {
			res++
		}
		m[v] = 1
	}
	return res
}

func longestString(x int, y int, z int) int {
	return 4*min(x, y) + 2*z + 2*min(1, min(z, abs(x-y)))
}

func minimizeConcatenatedLength(words []string) int {
	if len(words) == 1 {
		return len(words[0])
	}
	dp := make([][][]int, len(words))
	for i := range dp {
		dp[i] = make([][]int, 26)
		for j := range dp[i] {
			dp[i][j] = make([]int, 26)
			for k := range dp[i][j] {
				dp[i][j][k] = 100000
			}
		}
	}
	dp[0][int(words[0][0]-'a')][int(words[0][len(words[0])-1]-'a')] = len(words[0])
	res := 1000000
	for i := 1; i < len(words); i++ {
		for j := 0; j < 26; j++ {
			for k := 0; k < 26; k++ {
				// stri-1在前
				cur := dp[i-1][j][k] + len(words[i])
				if k == int(words[i][0]-'a') {
					cur -= 1
				}
				if i == len(words)-1 {
					res = min(res, cur)
				}
				end := int(words[i][len(words[i])-1] - 'a')
				dp[i][j][end] = min(dp[i][j][end], cur)
				// words[i]在前
				cur = dp[i-1][j][k] + len(words[i])
				if j == int(words[i][len(words[i])-1]-'a') {
					cur -= 1
				}
				if i == len(words)-1 {
					res = min(res, cur)
				}
				start := int(words[i][0] - 'a')
				dp[i][start][k] = min(dp[i][start][k], cur)
			}
		}
	}
	return res
}
