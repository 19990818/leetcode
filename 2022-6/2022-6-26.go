package main

import (
	"math"
	"math/rand"
	"sort"
	"strings"
)

func minCost(costs [][]int) int {
	n := len(costs)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 3)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			dp[i+1][j] = min(dp[i][(j+1)%3]+costs[i][(j+1)%3], dp[i][(j+2)%3]+costs[i][(j+2)%3])
		}
	}
	ans := math.MaxInt64
	for j := 0; j < 3; j++ {
		ans = min(ans, dp[n][j])
	}
	return ans
}

type SolutionBlack struct {
	m     map[int]int
	total int
}

func ConstructorBlack(n int, blacklist []int) SolutionBlack {
	sort.Ints(blacklist)
	exist := make(map[int]int)
	m := make(map[int]int)
	index := n - 1
	for _, val := range blacklist {
		exist[val] = 1
		m[val] = val
	}
	//fmt.Println(exist)
	for _, val := range blacklist {
		if val >= index {
			break
		}
		for _, ok := exist[index]; index > val; {
			_, ok = exist[index]
			if !ok {
				break
			}
			index--
		}
		if index < val {
			break
		}
		//fmt.Println(val,index)
		m[val] = index
		exist[index] = 1
		index--
	}
	return SolutionBlack{m, n}
}

func (this *SolutionBlack) PickNBlack() int {
	index := rand.Intn(this.total - len(this.m))
	res := index
	if _, ok := this.m[res]; ok {
		res = this.m[index]
	}
	return res
}

func countAsterisks(s string) int {
	strArr := strings.Split(s, "|")
	res := 0
	for i, val := range strArr {
		if i%2 == 0 {
			res += strings.Count(val, "*")
		}
	}
	return res
}

func countPairs(n int, edges [][]int) int64 {
	out := make([][]int, n)
	for _, val := range edges {
		out[val[0]] = append(out[val[0]], val[1])
		out[val[1]] = append(out[val[1]], val[0])
	}
	parts := make([]int, 0)
	travel := make(map[int]int)
	var bfs func(num int) int
	bfs = func(num int) int {
		queue := make([]int, 0)
		queue = append(queue, num)
		ans := 1
		travel[num] = 1
		for {
			temp := make([]int, 0)
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				for _, val := range out[cur] {
					if travel[val] == 0 {
						travel[val] = 1
						ans++
						temp = append(temp, val)
					}
				}
			}
			if len(temp) == 0 {
				break
			}
			queue = temp
		}
		return ans
	}
	for i := 0; i < n; i++ {
		if travel[i] == 0 {
			parts = append(parts, bfs(i))
		}
	}
	//fmt.Println(parts)
	res := int64(0)
	if len(parts) == 0 {
		return res
	}
	for i := 0; i < len(parts); i++ {
		res += int64(parts[i] * (n - parts[i]))
	}
	return res / 2
}
