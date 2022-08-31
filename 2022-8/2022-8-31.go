package main

import "sort"

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	i, j := 0, 0
	for ; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[0 : len(stack)-1]
			j++
		}
	}
	return len(stack) == 0
}

func minIncrementForUnique(nums []int) int {
	m := make(map[int]int)
	ans := 0
	for _, val := range nums {
		cur := val
		for m[cur] == 1 {
			cur++
		}
		ans += cur - val
		m[cur] = 1
	}
	return ans
}

//本质上是一样的
func minIncrementForUnique2(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := range nums {
		if i < len(nums)-1 && nums[i+1] <= nums[i] {
			nums[i+1] = nums[i] + 1
			ans += nums[i+1] - nums[i]
		}
	}
	return ans
}

func bagOfTokensScore(tokens []int, power int) int {
	if len(tokens) == 0 {
		return 0
	}
	sort.Ints(tokens)
	if power < tokens[0] {
		return 0
	}
	ans := 0
	cnt := 0
	j := len(tokens) - 1
	for i := 0; i <= j; i++ {
		if power >= tokens[i] {
			power -= tokens[i]
			cnt++
			ans = max(ans, cnt)
		} else {
			cnt--
			power += tokens[j]
			j--
			i--
		}
	}
	return ans
}
