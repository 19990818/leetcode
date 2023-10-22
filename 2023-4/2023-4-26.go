package main

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	sum := make([][]int, len(nums))
	for i := range sum {
		sum[i] = make([]int, len(nums))
	}
	for i := 0; i < len(nums); i++ {
		sum[i][i] = nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum[i][j] = sum[i][j-1] + nums[j]
		}
	}
	res := 0
	for i := 0; i < len(nums)-firstLen; i++ {
		for j := 0; j < i && j < len(nums)-secondLen; j++ {
			res = max(res, sum[i][i+firstLen]+sum[j][j+secondLen])
		}
		for j := i + firstLen; j < len(nums)-secondLen; j++ {
			res = max(res, sum[i][i+firstLen]+sum[j][j+secondLen])
		}
	}
	return res
}
