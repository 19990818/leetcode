package main

import "math"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	dp := make([][]int, k+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt64
		}
	}
	dp[0][src] = 0
	ans := math.MaxInt64
	for t := 1; t <= k+1; t++ {
		for _, val := range flights {
			i, j, dis := val[0], val[1], val[2]
			var temp int
			if dp[t-1][i] == math.MaxInt64 {
				temp = math.MaxInt64
			} else {
				temp = dp[t-1][i] + dis
			}
			dp[t][j] = min(dp[t][j], temp)
		}
		ans = min(ans, dp[t][dst])
	}
	if ans == math.MaxInt64 {
		ans = -1
	}
	return ans
}

//二分图使用染色法 隔一个染上不同颜色
func isBipartite(graph [][]int) bool {
	m := make(map[int]int)
	valid := true
	var dfs func(color int, node int)
	dfs = func(color, node int) {
		color2 := 1
		if color == 1 {
			color2 = 2
		}
		for _, val := range graph[node] {
			if m[val] == 0 {
				m[val] = color2
				dfs(color2, val)
			} else if m[val] != color2 {
				valid = false
				return
			}
		}
	}
	for i := range graph {
		if m[i] == 0 {
			dfs(1, i)
		}
	}
	return valid
}

func preimageSizeFZF(k int) int {
	//5的个数决定0的个数
	//k==0 (0,1,2,3,4) k=1(5,6,7,8,9)
	//k==2 (10,11,12,13,14) k=3(15,16,17,18,19)
	//k==4 (20,21,22,23,24) k==6(25,26,27,28,29)
	//遇到可确定上下界限的 并且有可以确定左右条件的
	//就可以使用二分查找
	right := int(5 * (1e9))
	left := 0
	var countFive func(num int) int
	countFive = func(num int) int {
		count := 0
		for num > 0 {
			count += num / 5
			num /= 5
		}
		return count
	}
	for left < right {
		mid := (right-left)>>1 + left
		count := countFive(mid)
		if count == k {
			return 5
		} else if count < k {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return 0
}
