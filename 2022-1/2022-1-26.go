package main

func nthUglyNumber(n int) int {
	nums2 := make([]int, 0)
	nums3 := make([]int, 0)
	nums5 := make([]int, 0)
	if n == 1 {
		return 1
	}
	temp := 1
	for n > 1 {
		nums2 = append(nums2, temp*2)
		nums3 = append(nums3, temp*3)
		nums5 = append(nums5, temp*5)
		flag1, flag2, flag3 := 0, 0, 0
		if nums2[0] <= nums3[0] && nums2[0] <= nums5[0] {
			temp = nums2[0]
			flag1 = 1
		}
		if nums3[0] <= nums2[0] && nums3[0] <= nums5[0] {
			temp = nums3[0]
			flag2 = 1
		}
		if nums5[0] <= nums2[0] && nums5[0] <= nums3[0] {
			temp = nums5[0]
			flag3 = 1
		}
		if flag1 == 1 {
			nums2 = nums2[1:]
		}
		if flag2 == 1 {
			nums3 = nums3[1:]
		}
		if flag3 == 1 {
			nums5 = nums5[1:]
		}
		n--
	}
	return temp
}
