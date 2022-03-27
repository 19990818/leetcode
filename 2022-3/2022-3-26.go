package main

import (
	"math"
	"sort"
)

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	stack := make([]int, 0)
	last := math.MinInt64
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < last {
			return true
		}
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			last = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func circularArrayLoop(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	n := len(nums)
	var next func(cur int) int
	next = func(cur int) int {
		return ((cur+nums[cur])%n + n) % n
	}
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			continue
		}
		slow, fast := i, next(i)
		for nums[slow]*nums[fast] > 0 && nums[slow]*nums[next(fast)] > 0 {
			if slow == fast {
				//循环只有一个元素 不符合条件
				if slow == next(slow) {
					break
				}
				return true
			}
			slow = next(slow)
			fast = next(next(fast))
		}
		add := i
		for nums[add]*nums[next(add)] > 0 {
			temp := add
			add = next(add)
			nums[temp] = 0
		}
	}
	return false
}

func minMoves2(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for _, val := range nums {
		sum += val
	}
	leftSum := 0
	ans := math.MaxInt32
	n := len(nums)
	for i := 0; i < len(nums); i++ {
		leftSum += nums[i]
		sum -= nums[i]
		temp := (2*(i+1)-n)*nums[i] + sum - leftSum
		if temp < ans {
			ans = temp
		}
	}
	return ans
}
