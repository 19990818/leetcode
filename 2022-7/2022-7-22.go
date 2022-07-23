package main

import "sort"

//得将数组按照开始升序排序 结束降序排序 进行贪心处理
//如果不对结束进行降序排列的话会出现更新导致cur和next一样的情况导致出错
//进行贪心的基础是从后面开始每次选最前面两个数字，这样可以最大限度重复使用
//讨论的时候主要利用其集合中最大值和最小值构成的线段
//每次一个区间的时候是否在这个线段当中 没有一个点 一个点 两个点进行讨论
func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return intervals[i][1] > intervals[j][1]
	})
	ans := 2
	cur := intervals[len(intervals)-1][0]
	next := intervals[len(intervals)-1][0] + 1
	for i := len(intervals) - 2; i >= 0; i-- {
		if intervals[i][1] >= next {
			continue
		} else if intervals[i][1] < cur {
			cur = intervals[i][0]
			next = intervals[i][0] + 1
			ans += 2
		} else {
			next = cur
			cur = intervals[i][0]
			ans++
		}
	}
	return ans
}
