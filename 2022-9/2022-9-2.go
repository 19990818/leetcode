package main

import (
	"sort"
)

func flipgame(fronts []int, backs []int) int {
	ans := 0
	n := len(fronts)
	realM := make(map[int]int)
	for i := 0; i < n; i++ {
		if fronts[i] == backs[i] {
			realM[fronts[i]] = 1
		}
	}
	for i := 0; i < n; i++ {
		if fronts[i] == backs[i] {
			continue
		}
		bs := []int{fronts[i], backs[i]}
		for _, temp := range bs {
			if realM[temp] == 0 && (ans == 0 || ans > temp) {
				ans = temp
			}
		}

	}
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	ans := 0
	mod := int(1e9 + 7)
	dp := make(map[int]int)
	for i, val := range arr {
		dp[val] = 1
		for j := 0; j < i; j++ {
			if val%arr[j] == 0 {
				dp[val] = (dp[val] + dp[arr[j]]*dp[val/arr[j]]) % mod
			}
		}
		ans = (ans + dp[val]) % mod
	}
	return ans
}

func numFriendRequests(ages []int) int {
	sort.Ints(ages)
	//确定x
	ans := 0
	m := make(map[int]int)
	for i, val := range ages {
		m[val] = i
	}
	for i := len(ages) - 1; i >= 0; i-- {
		low := ages[i]/2 + 7
		pos := sort.SearchInts(ages, low+1)
		//fmt.Println(pos,m[ages[i]])
		if m[ages[i]]-pos > 0 {
			ans += m[ages[i]] - pos
		}
	}
	return ans
}
