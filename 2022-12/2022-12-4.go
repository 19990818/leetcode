package main

import (
	"math"
	"sort"
	"strings"
)

func isCircularSentence(sentence string) bool {
	arr := strings.Split(sentence, " ")
	for i := 0; i < len(arr); i++ {
		if arr[(i+1)%len(arr)][0] != arr[i][len(arr[i])-1] {
			return false
		}
	}
	return true
}

func dividePlayers(skill []int) int64 {
	sort.Ints(skill)
	ans := int64(0)
	temp := skill[0] + skill[len(skill)-1]
	for i, j := 0, len(skill)-1; i < j; i, j = i+1, j-1 {
		if skill[i]+skill[j] != temp {
			return -1
		}
		ans += int64(skill[i] * skill[j])
	}
	return ans
}

func minScore(n int, roads [][]int) int {
	g := make([][][]int, n+1)
	for _, val := range roads {
		g[val[0]] = append(g[val[0]], []int{val[1], val[2]})
		g[val[1]] = append(g[val[1]], []int{val[0], val[2]})
	}
	travel := make([]int, n+1)
	q := []int{1}
	travel[1] = 1
	res := math.MaxInt64
	for len(q) > 0 {
		temp := q
		q = []int{}
		for len(temp) > 0 {
			cur := temp[0]
			temp = temp[1:]
			for _, out := range g[cur] {
				if travel[out[0]] == 0 {
					q = append(q, out[0])
					travel[out[0]] = 1
				}
				res = min(res, out[1])
			}
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
