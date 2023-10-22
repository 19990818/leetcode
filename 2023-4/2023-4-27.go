package main

import "sort"

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	m := make(map[string]int)
	for i, word := range words {
		m[word] = i
	}
	dp := make([]int, len(words))
	dp[0] = 1
	res := 1
	for i := 1; i < len(words); i++ {
		dp[i] = 1
		for j := 0; j < len(words[i]); j++ {
			preStr := words[i][0:j] + words[i][j+1:]
			if _, ok := m[preStr]; ok {
				dp[i] = max(dp[i], dp[m[preStr]]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}
