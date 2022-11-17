package main

func splitArraySameAverage(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
}
