package main

func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	toX := []int{-1, 0, 1}
	toY := []int{-1, 0, 1}
	ans := make([][]int, n-2)
	for i := range ans {
		ans[i] = make([]int, n-2)
	}
	var findMax func(i, j int) int
	findMax = func(i, j int) int {
		ans := 0
		for _, x := range toX {
			for _, y := range toY {
				ans = max(ans, grid[i+x][j+y])
			}
		}
		return ans
	}
	for i := 0; i < n-2; i++ {
		for j := 0; j < n-2; j++ {
			ans[i][j] = findMax(i+1, j+1)
		}
	}
	return ans
}

func edgeScore(edges []int) int {
	m := make(map[int]int)
	n := len(edges)
	for i, v := range edges {
		m[v] += i
	}
	ans := 0
	cnt := 0
	for i := 0; i < n; i++ {
		if m[i] > cnt {
			ans = i
			cnt = m[i]
		}
	}
	return ans
}

func smallestNumber(pattern string) string {
	flag := make([]int, 9)
	var dfs func(l int)
	temp := []byte{}
	var ans string
	dfs = func(l int) {
		//fmt.Println(l,string(temp),len(pattern))
		if l == len(pattern)+1 {
			if ans == "" || ans > string(temp) {
				ans = string(temp)
			}
			return
		}
		if len(ans) == len(pattern)+1 {
			return
		}
		for i := 1; i <= 9; i++ {
			if flag[i-1] == 0 {
				if l > 0 && pattern[l-1] == 'D' && i > int(temp[len(temp)-1]-'0') {
					continue
				}
				if l > 0 && pattern[l-1] == 'I' && i < int(temp[len(temp)-1]-'0') {
					continue
				}
				flag[i-1] = 1
				temp = append(temp, byte(i+'0'))
				dfs(l + 1)
				flag[i-1] = 0
				temp = temp[0 : len(temp)-1]
			}
		}
	}
	dfs(0)
	return ans
}
