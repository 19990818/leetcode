package main

func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dpWidth, dpHeight := make([][]int, m+1), make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dpWidth[i] = make([]int, n+1)
		dpHeight[i] = make([]int, n+1)
	}
	ans := 0
	for i := 1; i <= m; i++ {

	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dpWidth[i][j] = dpWidth[i][j-1] + 1
				dpHeight[i][j] = dpHeight[i-1][j] + 1
				minWidth := dpWidth[i][j]
				//取得一个最长的宽度和高度 然后定下一个
				//讨论另外一个，遍历所有高度，然后取得这些高度中最窄的宽度
				//计算
				for k := 1; k <= dpHeight[i][j]; k++ {
					minWidth = min(minWidth, dpWidth[i][j-k+1])
					ans = max(ans, minWidth*dpHeight[i][j])
				}
			}
		}
	}
	return ans
}
