package main

import "math"

func kSimilarity(s1 string, s2 string) int {
	ans := math.MaxInt64
	var dfs func(s1, s2 []byte, cur, curans int)
	dfs = func(s1, s2 []byte, cur, curans int) {
		if curans > ans {
			return
		}
		if string(s1) == string(s2) {
			ans = curans
			return
		}
		if s1[cur] == s2[cur] {
			dfs(s1, s2, cur+1, curans)
		} else {
			for i := cur + 1; i < len(s1); i++ {
				if s1[i] != s2[i] && s1[i] == s2[cur] {
					s1[i], s1[cur] = s1[cur], s1[i]
					dfs(s1, s2, cur+1, curans+1)
					s1[i], s1[cur] = s1[cur], s1[i]
				}
			}
		}
	}
	dfs([]byte(s1), []byte(s2), 0, 0)
	return ans
}
