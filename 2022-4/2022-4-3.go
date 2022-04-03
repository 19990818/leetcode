package main

import (
	"sort"
	"strconv"
	"strings"
)

func convertTime(current string, correct string) int {
	currentInt, correctInt := 0, 0
	currentArr := strings.Split(current, ":")
	correctArr := strings.Split(correct, ":")
	h1, _ := strconv.Atoi(currentArr[0])
	m1, _ := strconv.Atoi(currentArr[1])
	h2, _ := strconv.Atoi(correctArr[0])
	m2, _ := strconv.Atoi(correctArr[1])
	currentInt = h1*60 + m1
	correctInt = h2*60 + m2
	ans := 0
	temp := correctInt - currentInt
	interval := []int{60, 15, 5, 1}
	//fmt.Println(temp)
	for _, val := range interval {
		for temp >= val {
			temp -= val
			ans++
		}
	}
	return ans
}

func findWinners(matches [][]int) [][]int {
	total := make(map[int]int)
	loser := make(map[int]int)
	for _, val := range matches {
		total[val[0]] = 1
		total[val[1]] = 1
		loser[val[1]]++
	}
	ans0 := make([]int, 0)
	ans1 := make([]int, 0)
	for key := range total {
		if loser[key] == 0 {
			ans0 = append(ans0, key)
		}
		if loser[key] == 1 {
			ans1 = append(ans1, key)
		}
	}
	sort.Ints(ans0)
	sort.Ints(ans1)
	return [][]int{ans0, ans1}
}

func maximumCandies(candies []int, k int64) int {
	l, r := 0, 10000000
	var f func(candies []int, m int) bool
	f = func(candies []int, m int) bool {
		count := int64(0)
		for _, val := range candies {
			count += int64(val) / int64(m)
		}
		return count >= k
	}
	var mid int
	for l < r {
		mid = (r-l)>>1 + l
		if f(candies, mid) {
			l = mid
		} else {
			r = mid - 1
		}
		if l == r-1 {
			if f(candies, r) {
				return r
			}
			break
		}
	}
	return l
}
