package main

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	m := make(map[int]int)
	for _, val := range items1 {
		m[val[0]] += val[1]
	}
	for _, val := range items2 {
		m[val[0]] += val[1]
	}
	ans := make([][]int, 0)
	for k, v := range m {
		ans = append(ans, []int{k, v})
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return ans
}

func countBadPairs(nums []int) int64 {
	m := make(map[int]int)
	n := len(nums)
	for idx, val := range nums {
		m[val-idx]++
	}
	wellPairs := int64(0)
	for _, val := range m {
		wellPairs += int64(val*(val-1)) / 2
	}
	return int64(n*(n-1)/2) - wellPairs
}

func taskSchedulerII(tasks []int, space int) int64 {
	m := make(map[int]int64)
	cnt := int64(1)
	for _, val := range tasks {
		if _, ok := m[val]; !ok {
			m[val] = cnt
			cnt++
		} else {
			cnt = max64(m[val]+int64(space+1), cnt)
			m[val] = cnt
			cnt++
		}
	}
	return cnt - 1
}
func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
