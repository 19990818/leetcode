package main

import (
	"math"
)

func deleteAndEarn(nums []int) int {
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	l := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		l[i] = m[i+1] * (i + 1)
	}
	dp := make([]int, 10000)
	dp[0] = l[0]
	dp[1] = l[1]
	dp[2] = l[2] + l[0]
	for i := 3; i < len(l); i++ {
		dp[i] = max(dp[i-2], dp[i-3]) + l[i]
	}
	return max(dp[9999], dp[9998])
}

func networkDelayTime(times [][]int, n int, k int) int {
	mOut := make(map[int][][]int)
	for _, val := range times {
		mOut[val[0]] = append(mOut[val[0]], val[1:])
	}
	arrive := make([]int, n)
	for i := 0; i < n; i++ {
		arrive[i] = math.MaxInt64
	}

	flag := make([]int, n)
	ans := -1
	var dfs func(length, count int, arr [][]int)
	dfs = func(length, count int, arr [][]int) {
		if count == n {
			if ans == -1 {
				ans = length
			} else {
				ans = min(ans, length)
			}
		}
		for _, val := range arr {
			if flag[val[0]-1] == 0 {
				flag[val[0]-1] = 1
				dfs(length+val[1], count+1, mOut[val[0]])
			}
		}
	}
	flag[k-1] = 1
	dfs(0, 1, mOut[k])
	return ans
}

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	sum := 0
	i := 0
	for sum < target {
		i++
		sum += i
	}
	//fmt.Println(i)
	if (sum-target)%2 == 0 {
		return i
	}
	if (sum+i+1-target)%2 == 0 {
		return i + 1
	}
	return i + 2
}

func pyramidTransition(bottom string, allowed []string) bool {
	m := make(map[string][]byte)
	for _, val := range allowed {
		m[val[0:2]] = append(m[val[0:2]], val[2])
	}
	var build func(s1 []byte, pos int) bool
	build = func(s1 []byte, pos int) bool {
		if len(s1) == 1 {
			return true
		}
		if len(s1)-1 == pos {
			return build(s1[0:pos], 0)
		}
		judge := false
		for _, val := range m[string(s1[pos:pos+2])] {
			s1[pos] = val
			judge = judge || build(s1, pos+1)
		}
		return judge
	}
	return build([]byte(bottom), 0)
}
