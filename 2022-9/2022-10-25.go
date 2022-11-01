package main

import (
	"sort"
	"strings"
)

func shortestBridge(grid [][]int) int {
	//其中有一个到达另一个即可
	travel := make(map[int]int)
	m, n := len(grid), len(grid[0])
	tos := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	cnt := 1
	bfs := func(cur int) {
		queue := []int{cur}
		for {
			temp := make([]int, 0)
			for len(queue) > 0 {
				now := queue[0]
				queue = queue[1:]
				x, y := now/n, now%n
				for _, to := range tos {
					if x+to[0] < m && x+to[0] >= 0 && y+to[1] < n && y+to[1] >= 0 &&
						travel[(x+to[0])*n+y+to[1]] == 0 {
						travel[(x+to[0])*n+to[1]] = cnt
					}
					temp = append(temp, (x+to[0])*n+y+to[1])
				}
			}
			if len(temp) == 0 {
				break
			}
			queue = temp
		}

	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 && travel[i*n+j] == 0 {
				travel[i*n+j] = cnt
				bfs(i*n + j)
				cnt++
			}
		}
	}
	part1 := make([]int, 0)
	for k, v := range travel {
		if v == 1 {
			part1 = append(part1, k)
		}
	}
	res := 0
	for {
		temp := make([]int, 0)
		for len(part1) > 0 {
			cur := part1[0]
			x, y := cur/n, cur%n
			part1 = part1[1:]
			for _, to := range tos {
				if x+to[0] < m && x+to[0] >= 0 && y+to[1] < n && y+to[1] >= 0 {
					if grid[x+to[0]][y+to[1]] == 0 {
						temp = append(temp, (x+to[0])*n+y+to[1])
					} else if travel[(x+to[0])*n+y+to[1]] == 2 {
						return res
					}
				}
			}
		}
		res++
		if len(temp) == 0 {
			break
		}
		part1 = temp
	}
	return -1
}

func oddString(words []string) string {
	diff := make(map[string]int)
	find := make(map[string]string)
	for _, word := range words {
		temp := make([]string, 0)
		for i := 1; i < len(word); i++ {
			temp = append(temp, string(word[i]-word[i-1]))
		}
		key := strings.Join(temp, ",")
		diff[key]++
		find[key] = word
	}
	for k, v := range diff {
		if v == 1 {
			return find[k]
		}
	}
	return ""
}

func destroyTargets(nums []int, space int) int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num%space]++
	}
	maxCnt := 0
	for _, v := range cnt {
		if v > maxCnt {
			maxCnt = v
		}
	}
	maxM := make(map[int]int)
	for k, v := range cnt {
		if v == maxCnt {
			maxM[k] = 1
		}
	}
	sort.Ints(nums)
	for _, num := range nums {
		if maxM[num%space] == 1 {
			return num
		}
	}
	return -1
}

func twoEditWords(queries []string, dictionary []string) []string {
	diff := func(a, b string) bool {
		if a == b {
			return true
		}
		cnt := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				cnt++
			}
			if cnt > 2 {
				return false
			}
		}
		return true
	}
	ans := make([]string, 0)
	for _, query := range queries {
		for _, dic := range dictionary {
			if diff(query, dic) {
				ans = append(ans, query)
				break
			}
		}
	}
	return ans
}

func secondGreaterElement(nums []int) []int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	dp[n-1][0], dp[n-1][1] = -1, -1
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			dp[i][0] = i + 1
			dp[i][1] = dp[i+1][0]
		} else {
			start := i + 1
			for {
				start = dp[start][0]
				if start == -1 {
					dp[i][0], dp[i][1] = -1, -1
					break
				}
				if nums[start] > nums[i] {
					dp[i][0] = start
					dp[i][1] = dp[start][0]
				}
			}
		}
	}
	ans := make([]int, 0)
	for _, v := range dp {
		ans = append(ans, v[1])
	}
	return ans
}
