package main

import "sort"

func countOperations(num1 int, num2 int) int {
	count := 0
	for num1 != 0 && num2 != 0 {
		if num1 >= num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
		count++
	}
	return count
}

func minimumOperations(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		if nums[0] == nums[1] {
			return 1
		}
		return 0
	}
	oddMap := make(map[int]int)
	evenMap := make(map[int]int)
	oddMax1, oddMax2, evenMax1, evenMax2 := 0, 0, 0, 0
	for i := 0; i < len(nums); i += 2 {
		evenMap[nums[i]]++
	}
	for i := 1; i < len(nums); i += 2 {
		oddMap[nums[i]]++
	}
	evenNums, oddNums := make([]int, 0), make([]int, 0)
	for _, val := range evenMap {
		evenNums = append(evenNums, val)
	}
	for _, val := range oddMap {
		oddNums = append(oddNums, val)
	}
	sort.Ints(evenNums)
	sort.Ints(oddNums)
	evenMax1, evenMax2 = getMaxPre2(evenMap, evenNums)
	oddMax1, oddMax2 = getMaxPre2(oddMap, oddNums)
	maxCount := 0
	//fmt.Println(evenNums,oddNums)
	if evenMax1 == oddMax1 {
		maxCount = max(evenMap[evenMax1]+oddMap[oddMax2], evenMap[evenMax2]+oddMap[oddMax1])
	} else {
		maxCount = evenNums[len(evenNums)-1] + oddNums[len(oddNums)-1]
	}
	return len(nums) - maxCount
}
func getMaxPre2(evenMap map[int]int, evenNums []int) (int, int) {
	evenMax1, evenMax2 := 0, 0
	for key, val := range evenMap {
		if val == evenNums[len(evenNums)-1] {
			if evenMax1 == 0 {
				evenMax1 = key
			} else {
				evenMax2 = key
			}
		}
		if len(evenNums) > 1 && val == evenNums[len(evenNums)-2] {
			if evenMax2 == 0 {
				evenMax2 = key
			}
		}
	}
	return evenMax1, evenMax2
}

func minimumRemoval(beans []int) int64 {
	maxRemain := int64(0)
	sort.Ints(beans)
	for i := len(beans) - 1; i >= 0; i-- {
		maxRemain = max64(maxRemain, int64(beans[i]*(len(beans)-i)))
	}
	sum := int64(0)
	for _, val := range beans {
		sum += int64(val)
	}
	return sum - maxRemain
}

func targetIndices(nums []int, target int) []int {
	ans := make([]int, 0)
	sort.Ints(nums)
	for key, val := range nums {
		if val == target {
			ans = append(ans, key)
		}
	}
	return ans
}
