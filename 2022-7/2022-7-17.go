package main

import (
	"sort"
)

func numberOfPairs(nums []int) []int {
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	ans := make([]int, 2)
	for _, val := range m {
		ans[0] += val / 2
		ans[1] += val % 2
	}
	return ans
}

func maximumSum(nums []int) int {
	m := make(map[int][]int)
	var getNumSumk func(a int) int
	getNumSumk = func(a int) int {
		res := 0
		for a > 0 {
			res += a % 10
			a /= 10
		}
		return res
	}
	for _, val := range nums {
		k := getNumSumk(val)
		m[k] = append(m[k], val)
	}
	ans := -1
	for _, val := range m {
		sort.Ints(val)
		if len(val) > 1 {
			ans = max(ans, val[len(val)-1]+val[len(val)-2])
		}
	}
	return ans
}

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	ans := make([]int, 0)
	temp := make([]int, len(nums))
	for i := range temp {
		temp[i] = i
	}
	for _, query := range queries {
		sort.Slice(temp, func(i, j int) bool {
			if nums[temp[i]][len(nums[temp[i]])-query[1]:] < nums[temp[j]][len(nums[temp[j]])-query[1]:] {
				return true
			}
			if nums[temp[i]][len(nums[temp[i]])-query[1]:] > nums[temp[j]][len(nums[temp[j]])-query[1]:] {
				return false
			}
			return temp[i] < temp[j]
		})
		//fmt.Println(temp)
		ans = append(ans, temp[query[0]-1])
	}
	return ans
}
