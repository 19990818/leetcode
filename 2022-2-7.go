package main

import (
	"math"
	"sort"
)

func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) int {
	minute := targetSeconds / 60
	second := targetSeconds % 60
	//得到每一个按键 然后根据输入进行处理 前置0可以不进行处理
	ans := math.MaxInt64
	if minute < 100 && minute >= 0 && second >= 0 && second < 100 {
		stepArr := getArr(minute, second)
		reverseArr(stepArr)
		ans = getCost(stepArr, startAt, pushCost, moveCost)
	}

	if second < 40 && minute > 0 {
		minute2 := minute - 1
		second2 := second + 60
		stepArr2 := getArr(minute2, second2)
		reverseArr(stepArr2)
		ans = min(ans, getCost(stepArr2, startAt, pushCost, moveCost))
	}
	return ans
}
func reverseArr(nums []int) {
	for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
}
func getCost(stepArr []int, startAt, pushCost, moveCost int) int {
	ans1 := 0
	flag := 0
	for _, val := range stepArr {
		if val != 0 {
			flag = 1
		}
		if flag == 0 {
			continue
		}
		if val != startAt {
			ans1 += moveCost
		}
		startAt = val
		ans1 += pushCost
	}
	return ans1
}
func getArr(minute, second int) []int {
	stepArr := make([]int, 0)
	count := 2
	for count > 0 {
		stepArr = append(stepArr, second%10)
		second /= 10
		count--
	}
	for minute > 0 {
		stepArr = append(stepArr, minute%10)
		minute /= 10
	}
	return stepArr
}

func pivotArray(nums []int, pivot int) []int {
	smaller := make([]int, 0)
	equal := make([]int, 0)
	bigger := make([]int, 0)
	for _, val := range nums {
		if val < pivot {
			smaller = append(smaller, val)
		} else if val == pivot {
			equal = append(equal, val)
		} else {
			bigger = append(bigger, val)
		}
	}
	ans := make([]int, 0)
	ans = append(ans, smaller...)
	ans = append(ans, equal...)
	ans = append(ans, bigger...)
	return ans
}

func sortEvenOdd(nums []int) []int {
	even := make([]int, 0)
	odd := make([]int, 0)
	for index, val := range nums {
		if index%2 == 0 {
			even = append(even, val)
		} else {
			odd = append(odd, val)
		}
	}
	sort.Ints(even)
	sort.Ints(odd)
	ans := make([]int, 0)
	//fmt.Println(even,odd)
	i, j := len(odd)-1, 0
	for ; j < len(even) && i >= 0; i, j = i-1, j+1 {
		ans = append(ans, even[j], odd[i])
	}
	if i >= 0 {
		ans = append(ans, odd[i])
	}
	if j < len(even) {
		ans = append(ans, even[j])
	}
	return ans
}

func smallestNumber(num int64) int64 {
	numArr := make([]int, 0)
	flag := 0
	if num < 0 {
		flag = 1
		num = -num
	}
	for num > 0 {
		numArr = append(numArr, int(num%10))
		num /= 10
	}
	ans := int64(0)
	sort.Ints(numArr)
	if flag == 0 {
		tempIt := 0
		for tempIt < len(numArr) && numArr[tempIt] == 0 {
			tempIt++
		}
		if tempIt != 0 {
			if tempIt >= len(numArr) {
				return 0
			}
			numArr[0], numArr[tempIt] = numArr[tempIt], 0
		}
		for _, val := range numArr {
			ans = 10*ans + int64(val)
		}
	} else {
		for i := len(numArr) - 1; i >= 0; i-- {
			ans = 10*ans + int64(numArr[i])
		}
		ans = -ans
	}
	return ans
}
