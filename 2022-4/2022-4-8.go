package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func findMinDifference(timePoints []string) int {
	var getMinute func(s string) int
	getMinute = func(s string) int {
		arr := strings.Split(s, ":")
		var ans int
		hour, _ := strconv.Atoi(arr[0])
		minute, _ := strconv.Atoi(arr[1])
		ans = hour*60 + minute
		return ans
	}
	minuteArr := make([]int, 0)
	for _, val := range timePoints {
		minuteArr = append(minuteArr, getMinute(val))
	}
	sort.Ints(minuteArr)
	// fmt.Println(minuteArr)
	res := math.MaxInt64
	for i := 1; i < len(minuteArr); i++ {
		res = min(res, minuteArr[i]-minuteArr[i-1])
		res = min(res, minuteArr[i-1]+1440-minuteArr[i])
	}
	res = min(res, minuteArr[0]+1440-minuteArr[len(minuteArr)-1])
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
