package main

import "reflect"

func subarraySum(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			ans++
		}
	}
	var cur, pre int
	for i := len(nums) - 2; i >= 0; i-- {
		pre = nums[i]
		for j := i + 1; j < len(nums); j++ {
			cur = pre + nums[j]
			pre = cur
			if cur == k {
				ans++
			}
		}
	}
	return ans
}

func subarraySum2(nums []int, k int) int {
	sumM := make(map[int]int)
	sumM[0] = 1
	sum := 0
	ans := 0
	for _, val := range nums {
		sum += val
		sumM[sum]++
		if _, ok := sumM[sum-k]; ok {
			ans += sumM[sum-k]
			if k == 0 {
				ans -= 1
			}
		}
	}
	return ans
}

func arrayNesting(nums []int) int {
	flag := make(map[int]int)
	ans := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := flag[i]; !ok {
			cur := i
			count := 1
			flag[i] = 1
			for {
				cur = nums[cur]
				if flag[cur] != 0 {
					break
				}
				flag[cur] = 1
				count++
			}
			ans = max(ans, count)
		}
	}
	return ans
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	aS1, aS2 := make([]int, 26), make([]int, 26)
	for _, val := range s1 {
		aS1[val-'a']++
	}
	for i := 0; i < len(s1); i++ {
		aS2[s2[i]-'a']++
	}
	if reflect.DeepEqual(aS1, aS2) {
		return true
	}
	for j := len(s1); j < len(s2); j++ {
		aS2[s2[j-len(s1)]-'a']--
		aS2[s2[j]-'a']++
		if reflect.DeepEqual(aS1, aS2) {
			return true
		}
	}
	return false
}
