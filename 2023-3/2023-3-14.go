package main

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	m, n := len(rowSum), len(colSum)
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = min(rowSum[i], colSum[j])
			rowSum[i] -= res[i][j]
			colSum[j] -= res[i][j]
		}
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximalNetworkRank(n int, roads [][]int) int {
	m := make([]int, n)
	t := make(map[int]int)
	for _, v := range roads {
		m[v[0]]++
		m[v[1]]++
		t[v[0]*n+v[1]] = 1
		t[v[1]*n+v[0]] = 1
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if m[i*n+j] == 1 {
				res = max(res, m[i]+m[j]-1)
			} else {
				res = max(res, m[i]+m[j])
			}
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
