package main

import "math/bits"

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	in := make(map[int]int)
	for _, val := range edges {
		in[val[1]]++
	}
	ans := make([]int, 0)
	for i := 0; i < n; i++ {
		if in[i] == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

func minOperations(nums []int) int {
	//看作每个数的操作，每个数的乘法取最大值，加法求和
	var cnt func(x int) (int, int)
	cnt = func(x int) (int, int) {
		res1 := bits.OnesCount(uint(x))
		res2 := bits.Len(uint(x))
		return res1, res2 - 1
	}
	ans := 0
	mulMax := 0
	for _, num := range nums {
		add, mul := cnt(num)
		//fmt.Println(num,add,mul)
		ans += add
		mulMax = max(mulMax, mul)
	}
	return ans + mulMax
}

func containsCycle(grid [][]byte) bool {
	tos := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	//m用来标记每个位置出现的时机 总是会在一个点进行相遇
	m, n := len(grid), len(grid[0])
	ma := make(map[int]int)
	var bfs func(cur int) bool
	bfs = func(cur int) bool {
		ma[cur] = 1
		s := 2
		queue := []int{cur}
		for {
			temp := make([]int, 0)
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				x, y := cur/n, cur%n
				for _, to := range tos {
					xp, yp := x+to[0], y+to[1]
					if xp < m && xp >= 0 && yp >= 0 && yp < n && grid[xp][yp] == grid[x][y] {
						next := xp*n + yp
						if ma[next] == 0 {
							ma[next] = s
							temp = append(temp, next)
						} else if s-ma[next] >= 4 {
							return true
						} else if ma[next] == s && 2*s >= 4 {
							return true
						}
					}
				}
			}
			s++
			if len(temp) == 0 {
				break
			}
			queue = temp
			//fmt.Println(queue)
		}
		return false
	}
	for i := range grid {
		for j := range grid[i] {
			if ma[i*n+j] == 0 {
				if bfs(i*n + j) {
					return true
				}
			}
		}
	}
	return false
}
