package main

func maxSumDivThree(nums []int) int {
	dp := make([]int, 3)
	res := 0
	for _, num := range nums {
		o := num % 3
		temp := append([]int{}, dp...)
		for i := 0; i < 3; i++ {
			if (i-o+3)%3 == 0 || temp[(i-o+3)%3] > 0 {
				dp[i] = max(dp[i], temp[(i-o+3)%3]+num)
			}
		}
		res = max(res, dp[0])
	}
	return res
}
