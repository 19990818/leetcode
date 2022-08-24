package main

import "math"

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cnt1 := 0
	queue := make([]int, 0)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 2 {
				queue = append(queue, i*n+j)
			}
			if grid[i][j] == 1 {
				cnt1++
			}
		}
	}
	ans := 0
	tos := [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	traval := make(map[int]int)
	for cnt1 > 0 {
		temp := make([]int, 0)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			x, y := cur/n, cur%n
			for _, to := range tos {
				if x+to[0] < m && x+to[0] >= 0 && y+to[1] < n && y+to[1] >= 0 && grid[x+to[0]][y+to[1]] == 1 && traval[(x+to[0])*n+y+to[1]] == 0 {
					around := (x+to[0])*n + y + to[1]
					traval[around] = 1
					cnt1--
					temp = append(temp, around)
				}
			}
		}
		if len(temp) == 0 {
			break
		}
		queue = temp
		ans++
	}
	if cnt1 > 0 {
		return -1
	}
	return ans
}

func largestSumOfAverages(nums []int, k int) float64 {
	n := len(nums)
	dp := make([][]float64, n)
	for i := range dp {
		dp[i] = make([]float64, k+1)
	}
	sum := make([]int, n)
	for i, val := range nums {
		if i == 0 {
			sum[i] = val
		} else {
			sum[i] = sum[i-1] + val
		}
	}
	for i := 0; i < n; i++ {
		dp[i][1] = float64(sum[i]) / float64(i+1)
	}
	for i := 0; i < n; i++ {
		for c := 2; c <= k && c <= i+1; c++ {
			for j := 0; j < i; j++ {
				dp[i][c] = math.Max(dp[i][c], dp[j][c-1]+float64(sum[i]-sum[j])/float64(i-j))
			}
		}
	}
	return dp[n-1][k]
}
