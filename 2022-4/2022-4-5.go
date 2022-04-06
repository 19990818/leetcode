package main

import (
	"sort"
)

type customStrSort []string

func (m customStrSort) Len() int {
	return len(m)
}
func (m customStrSort) Less(i, j int) bool {
	if len(m[i]) > len(m[j]) {
		return true
	}
	if len(m[i]) < len(m[j]) {
		return false
	}
	k := 0
	for ; k < len(m[i]) && k < len(m[j]); k++ {
		if m[i][k] < m[j][k] {
			return true
		}
		if m[i][k] > m[j][k] {
			return false
		}
	}
	if k < len(m[i]) {
		return false
	}
	return true
}

func (m customStrSort) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func findLongestWord(s string, dictionary []string) string {
	sort.Sort(customStrSort(dictionary))
	for _, val := range dictionary {
		if isSubsuquence(val, s) {
			return val
		}
	}
	return ""
}

func findMaxLength(nums []int) int {
	sumMp := make(map[int]int)
	sum := 0
	ans := 0
	sumMp[0] = -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			sum += 1
		} else {
			sum -= 1
		}
		if _, ok := sumMp[sum]; ok {
			ans = max(ans, i-sumMp[sum])
		} else {
			sumMp[sum] = i
		}
	}
	return ans
}

func countArrangement(n int) int {
	m := make(map[int]int)
	var dfs func(m map[int]int, n, cur int)
	ans := 0
	dfs = func(m map[int]int, n, cur int) {
		if cur == n+1 {
			ans++
		}
		for i := 1; i <= n; i++ {
			if _, ok := m[i]; ok {
				continue
			}
			if i%cur == 0 || cur%i == 0 {
				m[i] = 1
				dfs(m, n, cur+1)
				delete(m, i)
			}
		}
	}
	dfs(m, n, 1)
	return ans
}
