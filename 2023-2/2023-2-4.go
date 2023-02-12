package main

import "sort"

func getMaximumConsecutive(coins []int) int {
	//关键是将其作为一个连续的区间看待 起始[0,0]
	sort.Ints(coins)
	start, end := 0, 0
	for _, v := range coins {
		if start+v > end+1 {
			// 不连续
			break
		}
		end = end + v
	}
	return end
}
