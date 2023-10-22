package main

import "strings"

func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
	res := 0
	for _, hour := range hours {
		if hour >= target {
			res++
		}
	}
	return res
}

func countCompleteSubarrays(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	res := 0
	left, right := 0, 0
	temp := make(map[int]int)
	for left <= right {
		for right < len(nums) && len(temp) < len(m) {
			temp[nums[right]]++
			right++
		}
		if len(temp) == len(m) {
			res += len(nums) - right + 1
		}
		temp[nums[left]]--
		if temp[nums[left]] == 0 {
			delete(temp, nums[left])
		}
		left++
	}
	return res
}

func minimumString(a string, b string, c string) string {
	ss := []string{a, b, c}
	res := a + b + c
	m := make(map[int]int)
	var dfs func(cur string)
	dfs = func(cur string) {
		if len(m) == len(ss) {
			if len(cur) < len(res) || (len(cur) == len(res) && cur < res) {
				res = cur
			}
			return
		}
		for i, s := range ss {
			if m[i] == 0 {
				//选中s加入cur中
				if strings.Contains(cur, s) {
					m[i] = 1
					dfs(cur)
					delete(m, i)
					continue
				}
				for j := 0; j <= len(cur); j++ {
					if len(cur)-j <= len(s) && cur[j:] == s[0:len(cur)-j] {
						m[i] = 1
						dfs(cur + s[len(cur)-j:])
						delete(m, i)
						break
					}
				}
			}
		}
	}
	dfs("")
	return res
}
