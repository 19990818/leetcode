package main

import "sort"

func findKDistantIndices(nums []int, key int, k int) []int {
	ans := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if abs(j-i) <= k && nums[j] == key {
				ans = append(ans, i)
				break
			}
		}
	}
	return ans
}

func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
	type cordinate struct {
		x int
		y int
	}
	m := make([][]cordinate, 0)
	for _, val := range artifacts {
		temp := make([]cordinate, 0)
		for i := val[0]; i <= val[2]; i++ {
			for j := val[1]; j <= val[3]; j++ {
				temp = append(temp, cordinate{i, j})
			}
		}
		m = append(m, temp)
	}
	arrMap := make(map[cordinate]int)
	for _, val := range dig {
		arrMap[cordinate{val[0], val[1]}] = 1
	}
	ans := 0
	for _, val := range m {
		flag := 1
		for _, val2 := range val {
			if _, ok := arrMap[val2]; !ok {
				flag = 0
				break
			}
		}
		ans += flag
	}
	return ans
}

func maximumTop(nums []int, k int) int {
	if len(nums) == 1 {
		if k%2 == 1 {
			return -1
		}
		return nums[0]
	}
	pending := make([]int, 0)
	for k > 1 && len(nums) != 0 {
		pending = append(pending, nums[0])
		nums = nums[1:]
		k--
	}
	sort.Ints(pending)
	if k == 1 && len(pending) == 0 {
		if len(nums) <= 1 {
			return -1
		}
		return nums[0]
	}
	if k == 1 && len(pending) != 0 {
		if len(nums) != 0 {
			if nums[0] > pending[len(pending)-1] {
				return nums[0]
			}
		}
		return pending[len(pending)-1]
	}
	//len(nums)==0 k>1
	return pending[len(pending)-1]
}

func minimumWeight(n int, edges [][]int, src1 int, src2 int, dest int) int64 {

}
