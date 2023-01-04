package main

func maxValue(n int, index int, maxSum int) int {
	left, right := 0, maxSum+1
	for left+1 < right {
		mid := (left + right) >> 1
		sum := compute(mid-1, index) + mid + compute(mid-1, n-index-1)
		if sum <= maxSum {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
func compute(start, l int) int {
	sum := 0
	if l >= start {
		sum += (1+start)*start/2 + (l - start)
	} else {
		sum += (start + start + 1 - l) * l / 2
	}
	return sum
}
