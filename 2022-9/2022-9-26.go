package main

import "math"

func missingTwo(nums []int) []int {
	n := len(nums) + 2
	sum := (n*n + n) / 2
	psum := n * (n + 1) * (2*n + 1) / 6
	cursum := 0
	curPsum := 0
	for _, val := range nums {
		curPsum += val * val
		cursum += val
	}
	x := psum - curPsum
	y := sum - cursum
	return []int{(y + int(math.Sqrt(float64(2*x-y*y)))) / 2, (y - int(math.Sqrt(float64(2*x-y*y)))) / 2}
}
