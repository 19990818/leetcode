package main

import "math"

func calculateTax(brackets [][]int, income int) float64 {
	ans := float64(0)
	i := 0
	for income > 0 {
		if i == 0 {
			ans += float64(min(income, brackets[i][0])*brackets[i][1]) / 100
			income -= min(income, brackets[i][0])
		} else {
			ans += float64(min(income, brackets[i][0]-brackets[i-1][0])*brackets[i][1]) / 100
			income -= min(income, brackets[i][0]-brackets[i-1][0])
		}
		i++
	}
	return ans
}

func minPathCost(grid [][]int, moveCost [][]int) int {
	//记录到达每个位置使用的值
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if dp[i][j] == 0 {
					dp[i][j] = dp[i-1][k] + moveCost[grid[i-1][k]][j] + grid[i][j]
				}
				dp[i][j] = min(dp[i][j], dp[i-1][k]+moveCost[grid[i-1][k]][j]+grid[i][j])
			}
		}
	}
	//fmt.Println(dp)
	ans := math.MaxInt64
	for j := 0; j < n; j++ {
		ans = min(ans, dp[m-1][j])
	}
	return ans
}

func distributeCookies(cookies []int, k int) int {
	//小于平均数的可以分cookies
	sum := 0
	for _, val := range cookies {
		sum += val
	}
	avg := sum / k
	ans := math.MaxInt64
	var dfs func(cur []int, i int)
	dfs = func(cur []int, i int) {
		if i == len(cookies) {
			//fmt.Println(cur)
			tempMax := 0
			for _, val := range cur {
				if tempMax < val {
					tempMax = val
				}
			}
			ans = min(ans, tempMax)
			return
		}
		queue := make([]int, 0)
		for index, val := range cur {
			if val < avg {
				queue = append(queue, index)
			}
		}
		for _, val := range queue {
			cur[val] += cookies[i]
			dfs(cur, i+1)
			cur[val] -= cookies[i]
		}
	}
	nums := make([]int, k)
	dfs(nums, 0)
	return ans
}

// func distinctNames(ideas []string) int64 {
// 	m := make(map[string][]byte)
// 	m2 := make(map[byte]int)
// 	for _, val := range ideas {
// 		m[val[1:]] = append(m[val[1:]], val[0])
// 		m2[val[0]]++
// 	}
// 	ans := int64(0)
// 	ans = int64(len(ideas)) * int64(len(ideas)-1)
// 	for _, val := range m {
// 		for _, b := range val {
// 			if m2[b] > 0 {
// 				ans -= int64(m2[b] - 1)
// 			}
// 		}
// 	}
// 	return ans
// }

func findAndReplacePattern(words []string, pattern string) []string {
	ans := make([]string, 0)
	for _, val := range words {
		m1 := make(map[byte]byte)
		m2 := make(map[byte]byte)
		flag := 0
		for i := 0; i < len(val); i++ {
			if _, ok := m1[val[i]]; !ok {
				m1[val[i]] = pattern[i]
			} else if pattern[i] != m1[val[i]] {
				flag = 1
				break
			}
			if _, ok := m2[pattern[i]]; !ok {
				m2[pattern[i]] = val[i]
			} else if val[i] != m2[pattern[i]] {
				flag = 1
				break
			}
		}
		if flag == 0 {
			ans = append(ans, val)
		}
	}
	return ans
}
