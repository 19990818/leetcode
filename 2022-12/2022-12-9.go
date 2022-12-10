package main

import "sort"

func maxHeight(cuboids [][]int) int {
	for i := range cuboids {
		sort.Ints(cuboids[i])
	}
	n := len(cuboids)
	dp := make([]int, n)
	sort.Slice(cuboids, func(i, j int) bool {
		if cuboids[i][2] > cuboids[j][2] {
			return true
		}
		if cuboids[i][2] < cuboids[j][2] {
			return false
		}
		if cuboids[i][1] > cuboids[j][1] {
			return true
		}
		if cuboids[i][1] < cuboids[j][1] {
			return false
		}
		return cuboids[i][0] >= cuboids[j][0]
	})
	res := 0
	for i := range cuboids {
		dp[i] = cuboids[i][2]
		for j := 0; j < i; j++ {
			if cuboids[i][0] <= cuboids[j][0] && cuboids[i][1] <= cuboids[j][1] {
				dp[i] = max(dp[i], dp[j]+cuboids[i][2])
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func xorQueries(arr []int, queries [][]int) []int {
	xors := make([]int, len(arr)+1)
	for i := 0; i < len(arr); i++ {
		xors[i+1] = xors[i] ^ arr[i]
	}
	res := make([]int, len(queries))
	for i, v := range queries {
		res[i] = xors[v[0]] ^ xors[v[1]+1]
	}
	return res
}

func canReach(arr []int, start int) bool {
	m, n := make(map[int]int), len(arr)
	var canNext = func(cur, offset int) bool {
		return cur+offset < n && cur+offset >= 0 && m[cur+offset] == 0
	}
	q := []int{start}
	for len(q) > 0 {
		temp := q
		q = []int{}
		for len(temp) > 0 {
			cur := temp[0]
			temp = temp[1:]
			if arr[cur] == 0 {
				return true
			}
			if canNext(cur, arr[cur]) {
				m[cur+arr[cur]] = 1
				q = append(q, cur+arr[cur])
			}
			if canNext(cur, -arr[cur]) {
				m[cur-arr[cur]] = 1
				q = append(q, cur-arr[cur])
			}
		}
	}
	return false
}
