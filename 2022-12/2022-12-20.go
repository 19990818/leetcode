package main

func minimumSize(nums []int, maxOperations int) int {
	left, right := 1, 1
	for _, num := range nums {
		right = max(num, right)
	}
	for left < right {
		cnt := 0
		mid := (right + left) >> 1
		for _, num := range nums {
			cnt += num / mid
			if num%mid == 0 {
				cnt--
			}
		}
		if cnt <= maxOperations {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
