package main

import (
	"sort"
	"strings"
)

func trap(height []int) int {
	idx, mv := 0, height[0]
	for i, v := range height {
		if v > mv {
			mv = max(mv, v)
			idx = i
		}
	}
	temp := height[0]
	cidx := 0
	ans := 0
	for i := 1; i <= idx; i++ {
		if height[i] > temp {
			ans += temp * (i - cidx - 1)
			temp = height[i]
			cidx = i
		} else {
			ans -= height[i]
		}
	}
	temp = height[len(height)-1]
	cidx = len(height) - 1
	for i := len(height) - 1; i >= idx; i-- {
		if height[i] > temp {
			ans += temp * (cidx - i - 1)
			temp = height[i]
			cidx = i
		} else {
			ans -= height[i]
		}
	}
	return ans
}

func minimumIndex(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	nmc := nums[0]
	cnt := m[nums[0]]
	for k, v := range m {
		if v > cnt {
			nmc, cnt = k, v
		}
	}
	cnt = 0
	for i, v := range nums {
		if v == nmc {
			cnt++
		}
		if cnt*2 > (i+1) && (m[nmc]-cnt)*2 > (len(nums)-i-1) {
			return i
		}
	}
	return -1
}

func isGood(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for i := 1; i < len(nums); i++ {
		if i < len(nums)-1 && m[i] == 1 {
			continue
		}
		if i == len(nums)-1 && m[i] == 2 {
			continue
		}
		return false
	}
	return true
}

func maxArrayValue(nums []int) int64 {
	res := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		if res >= nums[i] {
			res += nums[i]
		} else {
			res = nums[i]
		}
	}
	return int64(res)
}

func sortVowels(s string) string {
	vowels := "AEIOUaeiou"
	idxs := make([]int, 0)
	strs := make([]rune, 0)
	for i, v := range s {
		if strings.ContainsRune(vowels, v) {
			idxs = append(idxs, i)
			strs = append(strs, v)
		}
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i] < strs[j]
	})
	sarr := []rune(s)
	i := 0
	for _, idx := range idxs {
		sarr[idx] = strs[i]
		i++
	}
	return string(sarr)
}
