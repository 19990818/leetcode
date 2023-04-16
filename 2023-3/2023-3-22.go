package main

import "sort"

func bestTeamScore(scores []int, ages []int) int {
	p := make([][]int, len(scores))
	for i := range scores {
		p[i] = []int{scores[i], ages[i]}
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i][1] == p[j][1] {
			return p[i][0] < p[j][0]
		}
		return p[i][1] < p[j][1]
	})
	dp := make([]int, len(scores))
	dp[0] = p[0][0]
	res := dp[0]
	for i := 1; i < len(scores); i++ {
		dp[i] = p[i][0]
		for j := 0; j < i; j++ {
			if p[i][0] >= p[j][0] {
				dp[i] = max(dp[i], dp[j]+p[i][0])
			}
		}
		res = max(res, dp[i])
	}
	return res
}
