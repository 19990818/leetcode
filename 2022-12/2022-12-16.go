package main

import "math"

func minElements(nums []int, limit int, goal int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return int(math.Ceil(float64(abs(sum-goal))/float64(limit)))
}