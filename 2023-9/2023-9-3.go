package main

import (
	"sort"
	"strings"
)

func eliminateMaximum(dist []int, speed []int) int {
	n := len(dist)
	t := make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = (dist[i]-1)/speed[i] + 1
	}
	sort.Ints(t)
	for i := 0; i < n; i++ {
		if t[i] < i+1 {
			return i
		}
	}
	return n
}

func canBeEqual(s1 string, s2 string) bool {
	for i := 0; i < len(s1); i++ {
		flag := 1
		if i > 1 {
			flag = -1
		}
		if s1[i] == s2[i] || (s1[i] == s2[i+2*flag] && s1[i+2*flag] == s2[i]) {
			continue
		}
		return false
	}
	return true
}

func checkStrings(s1 string, s2 string) bool {
	var os1, os2, es1, es2 strings.Builder
	for i := 0; i < len(s1); i++ {
		if i%2 == 0 {
			os1.WriteByte(s1[i])
			os2.WriteByte(s2[i])
		} else {
			es1.WriteByte(s1[i])
			es2.WriteByte(s2[i])
		}
	}
	osb1 := []byte(os1.String())
	osb2 := []byte(os2.String())
	esb1 := []byte(es1.String())
	esb2 := []byte(es2.String())
	sort.Slice(osb1, func(i, j int) bool {
		return osb1[i] < osb1[j]
	})
	sort.Slice(osb2, func(i, j int) bool {
		return osb2[i] < osb2[j]
	})
	sort.Slice(esb1, func(i, j int) bool {
		return esb1[i] < esb1[j]
	})
	sort.Slice(esb2, func(i, j int) bool {
		return esb2[i] < esb2[j]
	})
	return string(osb1) == string(osb2) && string(esb1) == string(esb2)
}

func maxSum(nums []int, m int, k int) int64 {
	ma := make(map[int]int)
	res := 0
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
		ma[nums[i]]++
	}
	if len(ma) >= m {
		res = sum
	}
	for i := k; i < len(nums); i++ {
		ma[nums[i-k]]--
		if ma[nums[i-k]] == 0 {
			delete(ma, nums[i-k])
		}
		ma[nums[i]]++
		sum = sum - nums[i-k] + nums[i]
		if len(ma) >= m {
			res = max(res, sum)
		}
	}
	return int64(res)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
