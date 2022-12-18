package main

func canChoose(groups [][]int, nums []int) bool {
	i, j := 0, 0
	for i < len(groups) && j < len(nums) {
		temp := j
		for k := 0; k < len(groups[i]) && j < len(nums); k++ {
			if groups[i][k] == nums[j] {
				j++
			}
		}
		if j-temp == len(groups[i]) {
			i++
		} else {
			j = temp + 1
		}
	}
	return i == len(groups)
}
