package main

import "math"

func minimumScore(nums []int, edges [][]int) int {
	//分类讨论枚举切掉哪个边
	//利用xor性质进行o(1)计算，分类讨论需要用到节点之间的父母关系
	//使用时间戳记录每个节点的权重，之后可以o(1)得到节点间关系
	//记录以每个节点为根节点的子树的异或值，首先使用边讨论的时候注意是以孩子节点作为根节点
	//使用点讨论注意不能使用根节点，因为根节点不会有可以删除的边
	n := len(nums)
	g := make([][]int, n)
	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
		g[edge[1]] = append(g[edge[1]], edge[0])
	}
	t := 0
	xor := make([]int, n)
	in, out := make([]int, n), make([]int, n)
	var isParent func(x, y int) bool
	isParent = func(x, y int) bool {
		return in[x] <= in[y] && in[y] <= out[x]
	}
	var dfs func(x, pa int)
	dfs = func(x, pa int) {
		xor[x] = nums[x]
		t++
		in[x] = t
		for _, y := range g[x] {
			if y != pa {
				dfs(y, x)
				xor[x] ^= xor[y]
			}
		}
		out[x] = t
	}
	ans := math.MaxInt64
	//预处理使得x为y的父节点 减少讨论
	for i, edge := range edges {
		if isParent(edge[1], edge[0]) {
			edges[i][0], edges[i][1] = edges[i][1], edges[i][0]
		}
	}
	x, y, z := 0, 0, 0
	for i := 0; i < n-2; i++ {
		x1, y1 := edges[i][0], edges[i][1]
		for j := i + 1; j < n-1; j++ {
			x2, y2 := edges[j][0], edges[j][1]
			if isParent(y1, x2) {
				x = xor[y2]
				y = x ^ xor[y1]
				z = xor[0] ^ xor[y1]
			} else if isParent(y2, x1) {
				x = xor[y1]
				y = x ^ xor[y2]
				z = xor[0] ^ xor[y2]
			} else {
				x = xor[y1]
				y = xor[y2]
				z = xor[0] ^ x ^ y
			}
			maxTemp, minTemp := max(max(x, y), z), min(min(x, y), z)
			ans = min(ans, maxTemp-minTemp)
		}
	}
	return ans
}
