package main

import (
	"sort"
	"strconv"
)

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	x, y := tomatoSlices/2-cheeseSlices, 2*cheeseSlices-tomatoSlices/2
	if tomatoSlices%2 == 0 && x >= 0 && y >= 0 {
		return []int{x, y}
	}
	return []int{}
}

func sequentialDigits(low int, high int) []int {
	res := make([]int, 0)
	for i := cntLen(low); i <= cntLen(high); i++ {
		start, interval := getOriginAndInterval(i)
		if start <= high && start >= low {
			res = append(res, start)
		}
		for j := 0; j < 9-i; j++ {
			start = start + interval
			if start <= high && start >= low {
				res = append(res, start)
			}
		}
	}
	return res
}
func cntLen(a int) int {
	return len(strconv.Itoa(a))
}
func getOriginAndInterval(l int) (int, int) {
	res, interval := 0, 0
	for i := 1; i <= l; i++ {
		res = res*10 + i
		interval = interval*10 + 1
	}
	return res, interval
}

func isPossibleDivide(nums []int, k int) bool {
	m := make(map[int]int)
	sort.Ints(nums)
	for _, val := range nums {
		m[val]++
	}
	for _, val := range nums {
		if m[val] != 0 {
			for i := 0; i < k; i++ {
				if m[val+i] == 0 {
					return false
				}
				m[val+i]--
			}
		}
	}
	return true
}
