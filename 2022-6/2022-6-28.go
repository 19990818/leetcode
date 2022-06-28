package main

import "sort"

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	nums := make([]int, 0)
	numsR := make([]int, 0)
	sum := 0
	for i := range nums1 {
		nums = append(nums, nums1[i]-nums2[i])
		numsR = append(numsR, nums2[i]-nums1[i])
		sum += (nums1[i] + nums2[i])
	}
	var dymicPlan func(nums []int) int
	dymicPlan = func(nums []int) int {
		n := len(nums1)
		dp := make([][][]int, n)
		for i := range dp {
			dp[i] = make([][]int, 2)
			for j := range dp[i] {
				dp[i][j] = make([]int, 2)
			}
		}
		dp[0][0][0] = nums[0]
		dp[0][1][1] = -nums[0]
		for count := 0; count < 2; count++ {
			for i := 1; i < n; i++ {
				if count > 0 {
					dp[i][count][0] = max(dp[i-1][count][0], dp[i-1][count][1]) + nums[i]
					dp[i][count][1] = max(dp[i-1][count-1][0], dp[i-1][count][1]) - nums[i]
				} else {
					dp[i][count][0] = dp[i-1][count][0] + nums[i]
				}
			}
		}
		sub := max(dp[n-1][1][0], dp[n-1][1][1])
		sub = max(dp[n-1][0][0], sub)
		return sub
	}
	return (sum + max(dymicPlan(nums), dymicPlan(numsR))) / 2
}

func distinctSequences(n int) int {
	//统计以某个数结尾的数字的个数 将所有的相加得到答案
	//动态规划解决
	//二维存在重复的 在位置i处理，我们会将i-2同位置的所有选择去掉
	//而我们在i-1处理时，会处理i-3同位置的所有选择，被删除的选择我们要进行回补操作
	//这样就能够保证在处理位置i时候，子问题只有i-2同位置，与i-1没有关系
	//注意3,6之间的关系 因为3,6不是互质关系，首先一直对6的map中存在3导致错误
	mod := int(1e9 + 7)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 6)
	}
	for i := 0; i < 6; i++ {
		dp[0][i] = 1
	}
	m := make(map[int][]int)
	m[0] = []int{2, 3, 4, 5, 6}
	m[1] = []int{1, 3, 5}
	m[3] = []int{1, 3, 5}
	m[5] = []int{1, 5}
	m[2] = []int{1, 2, 4, 5}
	m[4] = []int{1, 2, 3, 4, 6}
	for i := 1; i < n; i++ {
		if i > 2 {
			for k := 0; k < 6; k++ {
				dp[i-1][k] = (dp[i-1][k] + dp[i-3][k]) % mod
			}
		}
		for j := 0; j < 6; j++ {
			for _, val := range m[j] {
				dp[i][j] = (dp[i][j] + dp[i-1][val-1]) % mod
				if i > 1 {
					dp[i][j] = (dp[i][j] - dp[i-2][j] + mod) % mod
				}
			}
		}
	}
	sum := 0
	for i := 0; i < 6; i++ {
		sum = (sum + dp[n-1][i]) % mod
	}
	return sum
}

func maximumProduct(nums []int) int {
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}
	sort.Ints(nums)
	n := len(nums)
	temp1 := nums[n-1] * nums[n-2] * nums[n-3]
	temp2 := nums[0] * nums[1] * nums[n-1]
	return max(temp1, temp2)
}
