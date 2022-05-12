package main

import (
	"math"
	"sort"
)

func knightProbability(n int, k int, row int, column int) float64 {
	if k == 0 {
		return 1
	}
	total := math.Pow(8, float64(k))
	move := make([][]int, 0)
	x := []int{2, -2}
	y := []int{1, -1}
	flag := make([][][]float64, n)
	for i := 0; i < n; i++ {
		flag[i] = make([][]float64, n)
		for j := 0; j < n; j++ {
			flag[i][j] = make([]float64, k+1)
		}
	}
	for _, val := range x {
		for _, val2 := range y {
			move = append(move, []int{val, val2})
			move = append(move, []int{val2, val})
		}
	}
	in := float64(0)
	var dfs func(row, col int, k int) float64
	dfs = func(row, col, k int) float64 {
		if row < 0 || row > n-1 || col < 0 || col > n-1 {
			return 0
		}
		if k == 0 {
			return 1
		}
		ans := float64(0)
		for _, val := range move {
			row += val[0]
			col += val[1]
			if row < 0 || row > n-1 || col < 0 || col > n-1 {
				row -= val[0]
				col -= val[1]
				continue
			}
			if flag[row][col][k-1] == 0 {
				flag[row][col][k-1] = dfs(row, col, k-1) + 1
			}
			ans += flag[row][col][k-1] - 1
			row -= val[0]
			col -= val[1]
		}
		return ans
	}
	in = dfs(row, column, k)
	//fmt.Println(in,total)
	return in / total
}

type a struct {
	str string
	val int
}
type m []a

func topKFrequent(words []string, k int) []string {
	ma := make(map[string]int)
	for _, val := range words {
		ma[val]++
	}

	as := make([]a, 0)
	for key, val := range ma {
		as = append(as, a{key, val})
	}
	sort.Sort(m(as))
	ans := make([]string, 0)
	for i := 0; i < k; i++ {
		ans = append(ans, as[i].str)
	}
	return ans
}
func (this m) Len() int {
	return len(this)
}
func (this m) Less(i, j int) bool {
	if this[i].val < this[j].val {
		return true
	}
	if this[i].val > this[j].val {
		return false
	}
	return this[i].str < this[j].str
}
func (this m) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	flag := make([][]int, m)
	for i := 0; i < m; i++ {
		flag[i] = make([]int, n)
	}
	x := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var move func(row, col int) int
	var isvalid func(row, col int) bool
	isvalid = func(row, col int) bool {
		return row >= 0 && row < m && col >= 0 && col < n
	}
	move = func(row, col int) int {
		ans := 1
		flag[row][col] = 1
		for _, val := range x {
			if isvalid(val[0]+row, val[1]+col) && flag[val[0]+row][val[1]+col] == 0 && grid[val[0]+row][val[1]+col] == 1 {
				flag[val[0]+row][val[1]+col] = 1
				ans += move(val[0]+row, val[1]+col)
			}

		}
		return ans
	}
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && flag[i][j] == 0 {
				res = max(res, move(i, j))
			}
		}
	}
	return res
}
