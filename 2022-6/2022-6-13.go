package main

import "strings"

func distinctNames(ideas []string) int64 {
	//反向尝试失败，因为反向考虑其中存在重复元素
	//后考虑正向情况，直接考虑什么情况可以进行交换
	//时间复杂度N2超时，并未合理使用hashmap
	//需要合理利用hashmap减少时间复杂度，因为只包含小写字母
	//实际上可用选择的交换并不多 只有26*25
	//因此可以判断每种交换的个数 相加即可
	swap := make([][]int, 26)
	for i := range swap {
		swap[i] = make([]int, 26)
	}
	hm := make(map[string]int)
	for _, val := range ideas {
		hm[val] = 1
	}
	for _, val := range ideas {
		tempB := int(val[0] - 'a')
		for i := 0; i < 26; i++ {
			if i == tempB {
				continue
			}
			var tempStr strings.Builder
			tempStr.WriteByte(byte('a' + i))
			tempStr.WriteString(val[1:])
			if hm[tempStr.String()] == 0 {
				swap[tempB][i]++
			}
		}
	}
	ans := int64(0)
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			//i,j为等效位置，其中每组元素可以和另外一组元素组合
			ans += int64(swap[i][j] * swap[j][i])
		}
	}
	return ans
}

func solveNQueens(n int) [][]string {
	ans := make([][]byte, n)
	res := make([][]string, 0)
	for i := range ans {
		ans[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			ans[i][j] = '.'
		}
	}
	var dfs func(sum, count int)
	dfs = func(sum, count int) {
		//fmt.Println(ans)
		if count == 0 {
			temp := make([]string, 0)
			for _, val := range ans {
				temp = append(temp, string(val))
			}
			res = append(res, temp)
			return
		}
		if sum >= n*n {
			return
		}
		row, col := sum/n, sum%n
		if statisfyColAndRow(row, col, ans) && statisfyDiagnal(row, col, ans) {
			ans[row][col] = 'Q'
			dfs(sum+1, count-1)
			ans[row][col] = '.'
		}
		dfs(sum+1, count)
	}
	dfs(0, n)
	return res
}
func statisfyColAndRow(row, col int, target [][]byte) bool {
	for i := 0; i < len(target); i++ {
		if target[i][col] == 'Q' || target[row][i] == 'Q' {
			return false
		}
	}
	return true
}
func statisfyDiagnal(row, col int, target [][]byte) bool {
	towards := [][]int{{-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
	n := len(target)
	for _, val := range towards {
		curCol, curRow := col+val[1], row+val[0]
		for curCol < n && curCol >= 0 && curRow >= 0 && curRow < n {
			if target[curRow][curCol] == 'Q' {
				return false
			}
			curCol += val[1]
			curRow += val[0]
		}
	}
	return true
}

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	ans := make([][]int, 0)
	var dfs func(curNode int, temp []int)
	dfs = func(curNode int, temp []int) {
		if curNode == n-1 {
			res := append([]int{}, temp...)
			ans = append(ans, res)
			return
		}
		for _, val := range graph[curNode] {
			temp = append(temp, val)
			dfs(val, temp)
			temp = temp[0 : len(temp)-1]
		}
	}
	temp := []int{0}
	dfs(0, temp)
	return ans
}
