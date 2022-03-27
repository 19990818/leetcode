package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	sum1 := make(map[int]int)
	sum2 := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum1[nums1[i]+nums2[j]]++
		}
	}
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			sum2[nums3[i]+nums4[j]]++
		}
	}
	ans := 0
	for key, val := range sum1 {
		ans += val * sum2[-key]
	}
	return ans
}
