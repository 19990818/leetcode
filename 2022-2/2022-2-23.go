package main

import "math"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	nodeExistMap := make(map[int]int)
	nodeMap := make(map[int][]int)
	for _, val := range edges {
		nodeExistMap[val[0]]++
		nodeMap[val[0]] = append(nodeMap[val[0]], val[1])
		nodeExistMap[val[1]]++
		nodeMap[val[1]] = append(nodeMap[val[1]], val[0])
	}
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if nodeExistMap[i] == 1 {
			queue = append(queue, i)
		}
	}
	var ans []int
	for len(queue) > 0 {
		ans = make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[0]
			ans = append(ans, top)
			queue = queue[1:]
			for _, v := range nodeMap[top] {
				nodeExistMap[v]--
				if nodeExistMap[v] == 1 {
					queue = append(queue, v)
				}
			}
		}
	}
	return ans
}

func nthSuperUglyNumber(n int, primes []int) int {
	length := len(primes)
	pointers := make([]int, length)
	dp := make([]int, n+1)
	nums := make([]int, length)
	for i := range nums {
		nums[i] = 1
	}
	for i := 1; i <= n; i++ {
		minNum := math.MaxInt64
		for j := range nums {
			minNum = min(minNum, nums[j])
		}
		dp[i] = minNum
		for j := 0; j < length; j++ {
			if dp[i] == nums[j] {
				pointers[j]++
				nums[j] = dp[pointers[j]] * primes[j]
			}
		}
	}
	return dp[n]
}

func kthDistinct(arr []string, k int) string {
	m := make(map[string]int)
	for _, val := range arr {
		m[val]++
	}
	count := k
	for _, val := range arr {
		if m[val] == 1 {
			count--
		}
		if count == 0 {
			return val
		}
	}
	return ""
}
