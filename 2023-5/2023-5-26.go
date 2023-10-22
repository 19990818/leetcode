package main

func shortestPathBinaryMatrix(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 500
		}
	}
	if grid[0][0] != 0 {
		return -1
	}
	dp[0][0] = 1
	q := [][]int{[]int{0, 0}}
	travel := make(map[int]int)
	travel[0] = 1
	for len(q) > 0 {
		temp := q
		q = nil
		for _, t := range temp {
			i, j := t[0], t[1]
			if grid[i][j] == 0 {
				for x := -1; x < 2; x++ {
					for y := -1; y < 2; y++ {
						if x != 0 || y != 0 {
							dp[i][j] = min(dp[i][j], dp[min(max(i+x, 0), m-1)][min(max(j+y, 0), n-1)]+1)
							if i+x < m && i+x >= 0 && j+y < n && j+y >= 0 && travel[(i+x)*n+j+y] == 0 {
								q = append(q, []int{i + x, j + y})
								travel[(i+x)*n+j+y] = 1
							}
						}
					}
				}
			}
		}
	}
	//fmt.Println(dp)
	if dp[m-1][n-1] >= 500 {
		return -1
	}
	return dp[m-1][n-1]
}
