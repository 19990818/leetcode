package main

import (
	"math"
	"sort"
)

func findClosest(words []string, word1 string, word2 string) int {
	ans := math.MaxInt64
	a1, a2 := 0, 0
	for idx, val := range words {
		if val == word1 {
			a1 = idx
			ans = min(ans, a1-a2)
		}
		if val == word2 {
			a2 = idx
			ans = min(ans, a2-a1)
		}
	}
	return ans
}

func numMatchingSubseq(s string, words []string) int {
	m := make([][]int, 26)
	for idx, val := range s {
		m[val-'a'] = append(m[val-'a'], idx)
	}
	var isSubStr func(m [][]int, src string) bool
	isSubStr = func(m [][]int, src string) bool {
		if len(m[src[0]-'a']) == 0 {
			return false
		}
		cur := m[src[0]-'a'][0]
		for i := 1; i < len(src); i++ {
			if len(m[src[i]-'a']) == 0 {
				return false
			}
			pos := sort.SearchInts(m[src[i]-'a'], cur)
			if pos == len(m[src[i]-'a']) {
				return false
			}
			if m[src[i]-'a'][pos] == cur {
				if pos+1 >= len(m[src[i]-'a']) {
					return false
				}
				cur = m[src[i]-'a'][pos+1]
			} else {
				cur = m[src[i]-'a'][pos]
			}

		}
		return true
	}
	ans := 0
	for _, val := range words {
		if isSubStr(m, val) {
			ans++
		}
	}
	return ans
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	ans := 0
	maxNums := make([]int, len(nums))
	lessRight := make([]int, len(nums))
	for idx, val := range nums {
		if val >= left && val <= right {
			maxNums[idx] += 1
			if idx > 0 {
				maxNums[idx] += lessRight[idx-1]
			}
		}
		if val <= right {
			lessRight[idx] += 1
		}
		if idx > 0 {
			if val <= right {
				if val < left {
					maxNums[idx] += maxNums[idx-1]
				}
				lessRight[idx] += lessRight[idx-1]
			} else {
				maxNums[idx] = 0
				lessRight[idx] = 0
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] >= left && nums[i] <= right {
			ans += lessRight[i]
		} else {
			ans += maxNums[i]
		}
	}
	return ans
}

func numSubarrayBoundedMax2(nums []int, left int, right int) int {
	pos := -1
	temp := 0
	ans := 0
	for idx, val := range nums {
		if val > right {
			pos = idx
		}
		if val >= left {
			temp = idx - pos
		}
		ans += temp
	}
	return ans
}
