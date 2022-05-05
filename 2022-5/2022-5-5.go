package main

import (
	"strconv"
	"strings"
)

func removeDigit(number string, digit byte) string {
	var pos int
	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			pos = i
			if i+1 < len(number) && number[i+1] > digit {
				break
			}
		}
	}
	var temp strings.Builder
	temp.WriteString(number[0:pos])
	temp.WriteString(number[pos+1:])
	return temp.String()
}

func minimumCardPickup(cards []int) int {
	ans := -1
	m := make(map[int]int)
	for idx, val := range cards {
		if _, ok := m[val]; ok {
			if ans == -1 {
				ans = idx - m[val] + 1
			} else {
				ans = min(ans, idx-m[val]+1)
			}
		}
		m[val] = idx
	}
	return ans
}

func countDistinct(nums []int, k int, p int) int {
	ans := 0
	m := make(map[string]int)
	var trans func(arr []int) string
	trans = func(arr []int) string {
		var temp strings.Builder
		for idx, val := range arr {
			if idx > 0 {
				temp.WriteString(",")
			}
			temp.WriteString(strconv.Itoa(val))
		}
		return temp.String()
	}
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i; j < len(nums); j++ {
			if nums[j]%p == 0 {
				count++
			}
			if count <= k {
				str := trans(nums[i : j+1])
				if _, ok := m[str]; !ok {
					ans++
					m[str] = 1
				}
			}
		}
	}
	return ans
}
