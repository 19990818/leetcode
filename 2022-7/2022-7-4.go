package main

import "sort"

func countPaths(grid [][]int) int {
	res := 0
	mod := int(1e9 + 7)
	travel := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	var dfs func(curx, cury int) int
	dfs = func(curx, cury int) int {
		if dp[curx][cury] != 0 {
			return dp[curx][cury]
		}
		dp[curx][cury] = 1
		for _, val := range travel {
			x, y := curx+val[0], cury+val[1]
			if x < m && x >= 0 && y < n && y >= 0 && grid[x][y] > grid[curx][cury] {
				dp[curx][cury] = (dp[curx][cury] + dfs(x, y)) % mod
			}
		}
		return dp[curx][cury]
	}
	//以每个点开始的严格递增路径至少为1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//当当前位置没有计算的时候进行计算
			if dp[i][j] == 0 {
				dp[i][j] = dfs(i, j)
			}
		}
	}
	//fmt.Println(dp)
	for i := range grid {
		for j := range grid[i] {
			res = (res + dp[i][j]) % mod
		}
	}
	return res
}

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	deadM := make(map[string]int)
	for _, val := range deadends {
		deadM[val] = 1
	}
	startStr := "0000"
	var getNextStr func(s string) []string
	getNextStr = func(s string) []string {
		ans := make([]string, 0)
		for i := 0; i < 4; i++ {
			byteS := []byte(s)
			byteS[i] = byte((int(byteS[i]-'0')+1+10)%10 + '0')
			s1 := string(byteS)
			byteS[i] = byte((int(byteS[i]-'0')-2+10)%10 + '0')
			s2 := string(byteS)
			//fmt.Println(s1,s2)
			ans = append(ans, s1, s2)
		}
		return ans
	}
	if deadM[startStr] != 0 {
		return -1
	}
	queue := make([]string, 0)
	disM := make(map[string]int)
	queue = append(queue, startStr)
	travelM := make(map[string]int)
	cnt := 0
	for {
		temp := make([]string, 0)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			disM[cur] = cnt
			travelM[cur] = 1
			for _, val := range getNextStr(cur) {
				if _, ok := travelM[val]; !ok && deadM[val] == 0 {
					travelM[val] = 1
					temp = append(temp, val)
				}
			}
		}
		cnt++
		if len(temp) == 0 || disM[target] != 0 {
			break
		}
		queue = temp
	}
	if disM[target] == 0 {
		return -1
	}
	//fmt.Println(disM)
	return disM[target]
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	temp := make([][]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			temp = append(temp, []int{arr[i], arr[j]})
		}
	}
	sort.Sort(a(temp))
	return temp[k-1]
}

type a [][]int

func (this a) Len() int {
	return len(this)
}
func (this a) Less(i, j int) bool {
	return (this[i][0] * this[j][1]) < (this[j][0] * this[i][1])
}
func (this a) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
