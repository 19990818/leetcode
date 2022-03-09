package main

func wiggleMaxLength(nums []int) int {
	up, down := make([]int, len(nums)), make([]int, len(nums))
	up[0], down[0] = 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			up[i] = max(down[i-1]+1, up[i-1])
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			down[i] = max(up[i-1]+1, down[i-1])
			up[i] = up[i-1]
		} else {
			down[i] = down[i-1]
			up[i] = up[i-1]
		}
	}
	return max(up[len(nums)-1], down[len(nums)-1])
}

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, val := range nums {
			if i >= val {
				dp[i] += dp[i-val]
			}
		}
	}
	return dp[target]
}
