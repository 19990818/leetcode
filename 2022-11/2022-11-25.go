package main

import "math"

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	// 首先我们先构造图
	n := len(edges) + 1
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	// 因为是双向的实际上我们可能会走回来，所以要把之前的记录下来，防止重复
	// 因为不存在环 所以我们两个点之间只有一条路径，因此只需要记录不往回走即可‘
	// 首先定义bob的路径，因为bob实际上只有一条路径可走，我们将bob到达每个点
	// 的时间记录下来，针对无法到达的点时间设置为最大
	// 当处理alice时,alice到达的时间如果比bob小 得到所有amout = 对半分
	bobTime := make([]int, n)
	for i := range bobTime {
		bobTime[i] = math.MaxInt64
	}
	var dfsBob func(c, pre, t int) bool
	dfsBob = func(c, pre, t int) bool {
		if c == 0 {
			bobTime[c] = t
			return true
		}
		for _, y := range g[c] {
			if y != pre && dfsBob(y, c, t+1) {
				bobTime[c] = t
				return true
			}
		}
		return false
	}
	dfsBob(bob, -1, 0)
	res := math.MinInt64
	g[0] = append(g[0], -1)
	var dfsAlice func(c, pre, t, total int)
	dfsAlice = func(c, pre, t, total int) {
		if t < bobTime[c] {
			total += amount[c]
		} else if t == bobTime[c] {
			total += amount[c] / 2
		}
		if len(g[c]) == 1 {
			res = max(res, total)
			return
		}
		for _, y := range g[c] {
			if y != pre {
				dfsAlice(y, c, t+1, total)
			}
		}
	}
	dfsAlice(0, -1, 0, 0)
	return res
}
