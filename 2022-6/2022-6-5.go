package main

import "sort"

func minMaxGame(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	temp := make([]int, 0)
	for i := 0; i < len(nums)/2; i++ {
		if i%2 == 0 {
			temp = append(temp, min(nums[2*i], nums[2*i+1]))
		} else {
			temp = append(temp, max(nums[2*i], nums[2*i+1]))
		}
	}
	return minMaxGame(temp)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func partitionArray(nums []int, k int) int {
	sort.Ints(nums)
	cur := nums[0]
	ans := 1
	for i := 1; i < len(nums); i++ {
		if nums[i]-cur <= k {
			continue
		} else {
			ans++
			cur = nums[i]
		}
	}
	return ans
}

func arrayChange(nums []int, operations [][]int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		m[val] = i
	}
	for _, val := range operations {
		nums[m[val[0]]] = val[1]
		m[val[1]] = m[val[0]]
		delete(m, val[0])
	}
	return nums
}
