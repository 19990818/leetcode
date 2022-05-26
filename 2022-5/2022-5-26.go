package main

import "sort"

func escapeGhosts(ghosts [][]int, target []int) bool {
	player := abs(target[0]) + abs(target[1])
	for _, val := range ghosts {
		if abs(val[0]-target[0])+abs(val[1]-target[1]) <= player {
			return false
		}
	}
	return true
}

func numTilings(n int) int {
	//L型的只能成对出现，并且length=3 种类为两种
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	mod := int(1e9 + 7)
	for i := 2; i <= n; i++ {
		if i == 2 {
			dp[i] = 2
		} else {
			dp[i] = (2*dp[i-1] + dp[i-3]) % mod
		}
	}
	return dp[n]
}

func customSortString(order string, s string) string {
	mLetter := make(map[byte]int)
	for i := 0; i < len(order); i++ {
		mLetter[order[i]] = i
	}
	ans := []byte(s)
	sort.Sort(strBytes{ans, mLetter})
	return string(ans)
}

type strBytes struct {
	bytes []byte
	ma    map[byte]int
}

func (m strBytes) Len() int {
	return len(m.bytes)
}
func (m strBytes) Less(i, j int) bool {
	return m.ma[m.bytes[i]] < m.ma[m.bytes[j]]
}
func (m strBytes) Swap(i, j int) {
	m.bytes[i], m.bytes[j] = m.bytes[j], m.bytes[i]
}
