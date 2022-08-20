package main

import (
	"math"
	"sort"
)

func champagneTower(poured int, query_row int, query_glass int) float64 {
	// 简单动态规划
	dp := make([]float64, query_row+1)
	dp[0] = float64(poured)
	for i := 1; i <= query_row; i++ {
		temp := append([]float64{}, dp...)
		for j := 0; j < i+1; j++ {
			if j > 0 && j < i {
				dp[j] = (math.Max(temp[j-1]-1, 0) + math.Max(temp[j]-1, 0)) / 2
			} else if j == 0 {
				dp[j] = (math.Max(temp[j]-1, 0)) / 2
			} else {
				dp[j] = (math.Max(temp[j-1]-1, 0)) / 2
			}
		}
	}
	return math.Min(dp[query_glass], 1)
}

func eventualSafeNodes(graph [][]int) []int {
	// 实际上为拓扑排序找到安全节点
	n := len(graph)
	in, out := make([][]int, n), make([]int, n)
	queue := make([]int, 0)
	for idx, val := range graph {
		out[idx] = len(val)
		if out[idx] == 0 {
			queue = append(queue, idx)
		}
		for _, val2 := range val {
			in[val2] = append(in[val2], idx)
		}
	}
	ans := make([]int, 0)
	for len(queue) > 0 {
		cur := queue[0]
		ans = append(ans, cur)
		queue = queue[1:]
		for _, val := range in[cur] {
			out[val]--
			if out[val] == 0 {
				queue = append(queue, val)
			}
		}
	}
	sort.Ints(ans)
	return ans
}

func expressiveWords(s string, words []string) int {
	type pair struct {
		b   byte
		cnt int
	}
	var transferS func(s string) []pair
	transferS = func(s string) []pair {
		cur := s[0]
		cnt := 1
		res := make([]pair, 0)
		for i := 1; i <= len(s); i++ {
			if i == len(s) || s[i] != cur {
				res = append(res, pair{cur, cnt})
				cnt = 1
				if i < len(s) {
					cur = s[i]
				}

			} else {
				cnt++
			}
		}
		return res
	}
	target := transferS(s)
	ans := 0
	for _, val := range words {
		src := transferS(val)
		//fmt.Println(target,src)
		if len(src) == len(target) {
			ans++
			for idx := range src {
				if src[idx].b == target[idx].b && (src[idx].cnt == target[idx].cnt || target[idx].cnt >= 3 && target[idx].cnt > src[idx].cnt) {
					continue
				}
				ans--
				break
			}
		}
	}
	return ans
}
