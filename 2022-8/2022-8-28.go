package main

import (
	"sort"
	"strings"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	sum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		sum[i+1] = nums[i] + sum[i]
	}
	ans := make([]int, len(queries))
	for idx, val := range queries {
		for i := 0; i < len(nums)+1; i++ {
			if i == len(nums) || val >= sum[i] && val < sum[i+1] {
				ans[idx] = i
				break
			}
		}
	}
	return ans
}

func removeStars(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			stack = append(stack, s[i])
		} else {
			stack = stack[0 : len(stack)-1]
		}
	}
	return string(stack)
}

func garbageCollection(garbage []string, travel []int) int {
	//统计哪些位置有g,p,m
	g, p, m := 0, 0, 0
	cnt := 0
	for i, val := range garbage {
		if strings.Contains(val, "G") {
			g = i
		}
		if strings.Contains(val, "P") {
			p = i
		}
		if strings.Contains(val, "M") {
			m = i
		}
		cnt += len(val)
	}
	gs, ps, ms := 0, 0, 0
	for i, val := range travel {
		if i < g {
			gs += val
		}
		if i < p {
			ps += val
		}
		if i < m {
			ms += val
		}
	}
	return gs + ps + ms + cnt
}

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	ans := make([][]int, k)
	for i := range ans {
		ans[i] = make([]int, k)
	}
	rgi, rgo := make([][]int, k), make([]int, k)
	for _, val := range rowConditions {
		rgi[val[1]-1] = append(rgi[val[1]-1], val[0]-1)
		rgo[val[0]-1]++
	}
	cgi, cgo := make([][]int, k), make([]int, k)
	for _, val := range colConditions {
		cgi[val[1]-1] = append(cgi[val[1]-1], val[0]-1)
		cgo[val[0]-1]++
	}
	var topsort func(input [][]int, out []int) []int
	topsort = func(input [][]int, out []int) []int {
		queue := make([]int, 0)
		for i, val := range out {
			if val == 0 {
				queue = append(queue, i)
			}
		}
		ans := make([]int, 0)
		for len(queue) > 0 {
			cur := queue[0]
			ans = append(ans, cur)
			queue = queue[1:]
			for _, val := range input[cur] {
				out[val]--
				if out[val] == 0 {
					queue = append(queue, val)
				}
			}
		}
		if len(ans) < k {
			return []int{}
		}
		return ans
	}
	rr, cr := topsort(rgi, rgo), topsort(cgi, cgo)
	if len(rr) == 0 || len(cr) == 0 {
		return [][]int{}
	}
	//表示每个数字的r,c值
	mr, mc := make(map[int]int), make(map[int]int)
	for i := range rr {
		mr[rr[i]+1] = k - 1 - i
		mc[cr[i]+1] = k - 1 - i
	}
	for i := 1; i <= k; i++ {
		ans[mr[i]][mc[i]] = i
	}
	return ans
}
