package main

import "sort"

func findClosestElements(arr []int, k int, x int) []int {
	pos := sort.SearchInts(arr, x)
	leftArr := arr[0:pos]
	rightArr := arr[pos:]
	left, right := pos-1, 0
	for k > 0 {
		switch {
		case left < 0:
			right++
		case right >= len(rightArr):
			left--
		case x-leftArr[left] <= rightArr[right]-x:
			left--
		default:
			right++
		}
		k--
	}
	return arr[left+1 : right+pos]
}
