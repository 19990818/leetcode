package main

import "sort"

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	check := func(d int) bool {
		cnt, x0 := 1, price[0]
		for _, x := range price {
			if x >= x0+d {
				cnt++
				x0 = x
			}
		}
		return cnt >= k
	}
	left, right := 0, price[len(price)-1]-price[0]+1
	for left+1 < right {
		mid := (left + right) >> 1
		if check(mid) {
			left = mid
		} else {
			right = mid
		}
	}
	return left
}
