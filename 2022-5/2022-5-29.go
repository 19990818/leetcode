package main

import (
	"sort"
	"strings"
)

func digitCount(num string) bool {
	m := make(map[rune]int)
	for _, val := range num {
		m[val]++
	}
	for idx, val := range num {
		if int(val-'0') != m[rune(idx+'0')] {
			return false
		}
	}
	return true
}

func largestWordCount(messages []string, senders []string) string {
	m := make(map[string]int)
	for i := 0; i < len(messages); i++ {
		count := len(strings.Split(messages[i], " "))
		m[senders[i]] += count
	}
	maxCount := 0
	for _, val := range m {
		maxCount = max(maxCount, val)
	}
	ans := ""
	var isBigger func(des, src string) bool
	isBigger = func(des, src string) bool {
		i, j := 0, 0
		for i < len(des) && j < len(src) {
			if abs(int(des[i]-src[i])) == 'a'-'A' {
				if des[i]-src[i] < 0 {
					return false
				}
				return true
			}
			if des[i] > src[i] {
				return true
			}
			if des[i] < src[i] {
				return false
			}
			i++
			j++
		}
		return i >= len(des)
	}
	for key, val := range m {
		if val == maxCount {
			if ans == "" {
				ans = key
			} else if isBigger(key, ans) {
				ans = key
			}
		}
	}
	return ans
}

func maximumImportance(n int, roads [][]int) int64 {
	nodes := make([]int, n)
	for _, val := range roads {
		nodes[val[0]]++
		nodes[val[1]]++
	}
	ans := int64(0)
	sort.Ints(nodes)
	for i := len(nodes) - 1; i >= 0; i-- {
		ans += int64(nodes[i] * (i + 1))
	}
	return ans
}
