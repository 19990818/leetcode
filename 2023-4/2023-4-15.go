package main

import "sort"

func gardenNoAdj(n int, paths [][]int) []int {
	nears := make([][]int, n+1)
	for _, path := range paths {
		nears[path[0]] = append(nears[path[0]], path[1])
		nears[path[1]] = append(nears[path[1]], path[0])
	}
	ans := make([]int, n)
	for i := 1; i <= n; i++ {
		for c := 1; c <= 4; c++ {
			flag := true
			for _, v := range nears[i] {
				if ans[v-1] == c {
					flag = false
				}
			}
			if flag {
				ans[i-1] = c
			}
		}
	}
	return ans
}

func findMatrix(nums []int) [][]int {
	res := make([][]int, 0)
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for len(m) > 0 {
		temp := make([]int, 0)
		for k := range m {
			if m[k] > 0 {
				temp = append(temp, k)
				m[k]--
				if m[k] == 0 {
					delete(m, k)
				}
			}
		}
		res = append(res, temp)
	}
	return res
}

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	diff := make([]int, len(reward1))
	res := 0
	for i := range reward1 {
		res += reward2[i]
		diff[i] = reward1[i] - reward2[i]
	}
	sort.Ints(diff)
	for i := len(diff) - 1; i >= len(diff)-k; i-- {
		res += diff[i]
	}
	return res
}
