package main

import (
	"sort"
	"strconv"
	"strings"
)

func alphabetBoardPath(target string) string {
	var res strings.Builder
	startx, starty := 0, 0
	for _, v := range target {
		x, y := int(v-'a')/5, int(v-'a')%5
		xo, yo := x-startx, y-starty
		m := map[int]byte{1: 'D', -1: 'U', 2: 'R', -2: 'L'}
		if v == 'z' {
			res = writeXandY(yo, &res, m, 2)
			res = writeXandY(xo, &res, m, 1)
		} else {
			res = writeXandY(xo, &res, m, 1)
			res = writeXandY(yo, &res, m, 2)
		}
		res.WriteString("!")
		startx, starty = x, y
	}
	return res.String()
}
func writeXandY(o int, res *strings.Builder, m map[int]byte, kind int) strings.Builder {
	if o < 0 {
		kind *= -1
	}
	for i := 0; i < abs(o); i++ {
		res.WriteByte(m[kind])
	}
	return *res
}

func findTheArrayConcVal(nums []int) int64 {
	n := len(nums) - 1
	cnt := 0
	res := int64(0)
	for i := 0; i < n; i++ {
		if cnt < n {
			cnt += 2
			res += int64(concatInt(nums[i], nums[n-i]))
		} else if cnt == n {
			res += int64(nums[i])
			cnt += 1
		}
		// fmt.Println(i,res)
	}
	return res
}
func concatInt(a, b int) int {
	temp := strconv.Itoa(a) + strconv.Itoa(b)
	res, _ := strconv.Atoi(temp)
	// fmt.Println(res)
	return res
}

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	res := int64(0)
	for i, v := range nums {
		start := max(sort.SearchInts(nums, lower-v), i+1)
		end := max(sort.SearchInts(nums, upper-v+1), i+1)
		res += int64(end - start)
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
