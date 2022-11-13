package main

import "sort"

func distinctAverages(nums []int) int {
	sort.Ints(nums)
	m := make(map[float64]int)
	for i := 0; i < len(nums)/2; i++ {
		temp := float64(nums[i]+nums[len(nums)-1-i]) / float64(2)
		m[temp] = 1
	}
	return len(m)
}

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1
	mod := int(1e9 + 7)
	for i := 1; i <= high; i++ {
		if i-zero >= 0 {
			dp[i] = (dp[i] + dp[i-zero]) % mod
		}
		if i-one >= 0 {
			dp[i] = (dp[i-one] + dp[i]) % mod
		}
	}
	res := 0
	for i := low; i <= high; i++ {
		res = (res + dp[i]) % mod
	}
	return res
}

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	// 我们同时将两人进行移动
	in := make(map[int][]int)
	for _, val := range edges {
		in[val[0]] = append(in[val[0]], val[1])
		in[val[1]] = append(in[val[1]], val[0])
	}
	a := bfs(0, bob, in, amount)
	res := 0
	for _, val := range a {
		res = max(res, val)
	}
	return res
}
func bfs(sa, sb int, in map[int][]int, amount []int) []int {
	// 搞一个数组记录每条路径得到的分数
	res := make([]int, 0)
	t1 := make(map[int]int)
	q1 := []int{sa}
	t1[sa] = 1
	cnt1 := []int{amount[sa]}
	meetP := meet(sa, sb, in)
	flag := make(map[int]int)
	flag[meetP] = 1
	for {
		temp := make([]int, 0)
		cnt1Temp := make([]int, 0)
		for len(q1) > 0 {
			cur1 := q1[0]
			q1 = q1[1:]
			c := cnt1[0]
			cnt1 = cnt1[1:]
			if len(in[cur1]) == 0 {
				res = append(res, c)
			}
			for _, val := range in[cur1] {
				if t1[val] == 0 {
					t1[val] = 1
					flag[val] = flag[cur1]
					if flag[val] == 1 {
						if val == meetP {
							c += amount[meetP] / 2
						}
					} else if flag[val] == 0 {
						c += amount[val]
					} else if val == sb {
						flag[val] = 0
					}
					cnt1Temp = append(cnt1Temp, c)
					temp = append(temp, val)
				}
			}
		}
		if len(temp) == 0 {
			break
		}
		cnt1 = cnt1Temp
		q1 = temp
	}
	return res
}
func meet(sa, sb int, in map[int][]int) int {
	t1, t2 := make(map[int]int), make(map[int]int)
	q1, q2 := []int{sa}, []int{sb}
	t1[sa], t2[sb] = 1, 1
	meet := -1
	for {
		temp1, temp2 := make([]int, 0), make([]int, 0)
		meetM := make(map[int]int)
		for len(q1) > 0 {
			cur1 := q1[0]
			cur2 := q2[0]
			for _, val := range in[cur1] {
				if t1[val] == 0 {
					t1[val] = 1
					temp1 = append(temp1, val)
					meetM[val] = 1
				}
			}
			for _, val := range in[cur2] {
				if t2[val] == 0 {
					t2[val] = 1
					temp2 = append(temp2, val)
					if meetM[val] == 1 {
						return meet
					}
				}
			}
		}
		if len(temp1) == 0 {
			break
		}
		q1, q2 = temp1, temp2
	}
	return -1
}
