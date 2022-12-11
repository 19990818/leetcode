package main

import (
	"math"
	"sort"
	"strconv"
)

func maximumValue(strs []string) int {
	res := 0
	for _, v := range strs {
		num, err := strconv.Atoi(v)
		if err == nil {
			res = max(res, num)
		} else {
			res = max(res, len(v))
		}
	}
	return res
}

func maxStarSum(vals []int, edges [][]int, k int) int {
	n := len(vals)
	out := make([][]int, n)
	for _, v := range edges {
		out[v[0]] = append(out[v[0]], vals[v[1]])
		out[v[1]] = append(out[v[1]], vals[v[0]])
	}
	for i := range out {
		sort.Ints(out[i])
	}
	res := math.MinInt64
	for i := 0; i < n; i++ {
		temp := vals[i]
		res = max(res, temp)
		for j := 0; j < k && j < len(out[i]); j++ {
			temp += out[i][len(out[i])-1-j]
			res = max(res, temp)
		}
	}
	return res
}

func maxJump(stones []int) int {
	return max(getMaxDistance(stones, 1), getMaxDistance(stones, 2))
}
func getMaxDistance(stones []int, i int) int {
	start, end := stones[0], stones[len(stones)-1]
	// 过去走的路程 选择第一个开始
	res := 0
	for ; i < len(stones); i += 2 {
		res = max(res, stones[i]-start)
		start = stones[i]
	}
	res = max(res, end-start)
	return res
}

func deleteGreatestValue(grid [][]int) int {
	for i := range grid {
		sort.Ints(grid[i])
	}
	res := 0
	for j := 0; j < len(grid[0]); j++ {
		temp := 0
		for i := 0; i < len(grid); i++ {
			temp = max(temp, grid[i][j])
		}
		res += temp
	}
	return res
}

func longestSquareStreak(nums []int) int {
	sort.Ints(nums)
	m, travel := make(map[int]int), make(map[int]int)
	for _, v := range nums {
		m[v] = 1
	}
	res := -1
	for _, v := range nums {
		if travel[v] == 0 {
			cnt := 1
			for m[v*v] == 1 {
				v = v * v
				cnt++
				travel[v] = 1
			}
			if cnt > 1 && cnt > res {
				res = cnt
			}
		}
	}
	return res
}
