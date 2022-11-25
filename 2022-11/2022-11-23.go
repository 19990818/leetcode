package main

func getMaximumGold(grid [][]int) int {
	tos := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	m, n := len(grid), len(grid[0])
	var dfs func(i, j int, ma map[int]int) int
	dfs = func(i, j int, ma map[int]int) int {
		ma[i*n+j] = 1
		temp := 0
		for _, to := range tos {
			if statisfy(i+to[0], j+to[1], m, n, ma, grid) {
				ma[(i+to[0])*n+j+to[1]] = 1
				temp = max(temp, dfs(i+to[0], j+to[1], ma))
				ma[(i+to[0])*n+j+to[1]] = 0
			}
		}
		return temp + grid[i][j]
	}
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != 0 {
				res = max(res, dfs(i, j, make(map[int]int)))
			}
		}
	}
	return res
}
func statisfy(i, j, m, n int, ma map[int]int, grid [][]int) bool {
	return i < m && i >= 0 && j < n && j >= 0 && ma[i*n+j] == 0 && grid[i][j] != 0
}
