package main

func minOperations(nums []int, x int) int {
	m := make(map[int]int)
	preSum := 0
	m[preSum] = 0
	for i, v := range nums {
		preSum = preSum + v
		m[preSum] = i + 1
	}
	res := -1
	if x > preSum {
		return res
	}
	preSum = 0
	for i := 0; i < len(nums); i++ {
		if _, ok := m[x-preSum]; ok {
			if res == -1 {
				res = m[x-preSum] + i
			} else {
				res = min(res, m[x-preSum]+i)
			}
		}
		preSum = preSum + nums[len(nums)-1-i]
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
