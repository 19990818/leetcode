package main

import (
	"math"
	"sort"
)

func closedIsland(grid [][]int) int {
	travel := make(map[int]int)
	res := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] == 0 && travel[i*len(grid[0])+j] == 0 && checkClosed(i, j, grid, travel) {
				res++
			}
		}
	}
	return res
}

func checkClosed(i, j int, grid [][]int, travel map[int]int) bool {
	m, n := len(grid), len(grid[0])
	tos := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	q := [][]int{{i, j}}
	travel[i*n+j] = 1
	res := true
	for len(q) > 0 {
		temp := q
		q = nil
		for _, v := range temp {
			if v[0] == m-1 || v[0] == 0 || v[1] == n-1 || v[1] == 0 {

				res = false
			}
			for _, to := range tos {
				x, y := v[0]+to[0], v[1]+to[1]
				if x < m && x >= 0 && y < n && y >= 0 && grid[x][y] == 0 && travel[x*n+y] == 0 {
					travel[x*n+y] = 1
					q = append(q, []int{x, y})
				}
			}
		}
	}
	return res
}

func distanceTraveled(mainTank int, additionalTank int) int {
	res := 0
	for mainTank >= 5 && additionalTank > 0 {
		res += 50
		mainTank -= 4
		additionalTank -= 1
	}
	return res + mainTank*10
}

func findValueOfPartition(nums []int) int {
	res := math.MaxInt64
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		res = min(res, nums[i]-nums[i-1])
	}
	return res
}

func specialPerm(nums []int) int {
	mod := int(1e9 + 7)
	m := (1 << len(nums)) - 1
	mem := make([][]int, 1<<len(nums))
	for i := range mem {
		mem[i] = make([]int, len(nums))
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == 0 {
			mem[i][j] = 1
			return mem[i][j]
		}
		if mem[i][j] != -1 {
			return mem[i][j]
		}
		res := 0
		for k, v := range nums {
			if i>>k&1 > 0 && (v%nums[j] == 0 || nums[j]%v == 0) {
				res = (res + dfs(i^(1<<k), k)) % mod
			}
		}
		mem[i][j] = res
		return res
	}
	res := 0
	for i := range nums {
		res = (res + dfs(m^(1<<i), i)) % mod
	}
	return res
}
