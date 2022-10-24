package main

import (
	"sort"
	"strconv"
	"strings"
)

func haveConflict(event1 []string, event2 []string) bool {
	job1 := event1[1]
	f := func(s string) int {
		arr := strings.Split(s, ":")
		h, _ := strconv.Atoi(arr[0])
		m, _ := strconv.Atoi(arr[1])
		return h*60 + m
	}
	return !(f(job1) < f(event2[0]) || f(event2[1]) < f(event1[0]))
}

func subarrayGCD(nums []int, k int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	gcd := func(a, b int) int {
		for {
			a, b = b, a%b
			if b == 0 {
				return a
			}
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		dp[i][i] = nums[i]
		if dp[i][i] == k {
			ans++
		}
		for j := i + 1; j < n; j++ {
			dp[i][j] = gcd(dp[i][j-1], nums[j])
			//fmt.Println(i,j,dp[i][j])
			if dp[i][j] == k {
				ans++
			}
		}
	}
	return ans
}

func minCost(nums []int, cost []int) int64 {
	costSum, n := 0, len(nums)
	pairs := make([][]int, n)
	for i := range nums {
		pairs[i] = []int{nums[i], cost[i]}
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0]
	})
	costSs := make([]int, len(nums))
	for i, val := range pairs {
		costSum += val[1]
		if i > 0 {
			costSs[i] = costSs[i-1] + val[1]
		}
	}
	dp := make([]int64, n)
	for i := 1; i < n; i++ {
		dp[0] += int64(pairs[i][0]-pairs[i-1][0]) * int64(cost[i])
	}
	ans := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1] - (int64(costSum-costSs[i]) * int64(pairs[i][0]-pairs[i-1][0]))
		dp[i] += int64(costSs[i]) * int64(pairs[i][0]-pairs[i-1][0])
		if dp[i] < ans {
			//fmt.Println(i)
			ans = dp[i]
		}
	}
	//fmt.Println(dp)
	return ans
}
func makeSimilar(nums []int, target []int) int64 {
	numodd, numeven := make([]int, 0), make([]int, 0)
	for _, num := range nums {
		if num%2 == 0 {
			numodd = append(numodd, num)
		} else {
			numeven = append(numeven, num)
		}
	}
	targetodd, targeteven := make([]int, 0), make([]int, 0)
	for _, t := range target {
		if t%2 == 0 {
			targetodd = append(targetodd, t)
		} else {
			targeteven = append(targeteven, t)
		}
	}
	sort.Ints(numodd)
	sort.Ints(numeven)
	sort.Ints(targetodd)
	sort.Ints(targeteven)
	ans := int64(0)
	for i := range numodd {
		ans += abs(numodd[i] - targetodd[i])
	}
	for i := range targeteven {
		ans += abs(numeven[i] - targeteven[i])
	}
	return ans / 4
}
func abs(a int) int64 {
	if a < 0 {
		return int64(-a)
	}
	return int64(a)
}
