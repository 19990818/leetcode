package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func rearrangeCharacters(s string, target string) int {
	m := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, val := range target {
		m2[val]++
	}
	for _, val := range s {
		m[val]++
	}
	ans := math.MaxInt64
	for key, val := range m2 {
		ans = min(ans, m[key]/val)
	}
	return ans
}

func discountPrices(sentence string, discount int) string {
	split := strings.Fields(sentence)
	for idx, val := range split {
		if len(val) > 0 && val[0] == '$' {
			temp, err := strconv.Atoi(val[1:])
			if err == nil {
				succ := fmt.Sprintf("$%.2f", float64(temp)*float64(100-discount)/100)
				split[idx] = succ
			}
		}
	}
	return strings.Join(split, " ")
}

func totalSteps(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	type kV struct {
		v     int
		count int
	}
	//关键为等效替代 中间不符合递增的数字需要被删除 那么记录下删除不合适的数字最大次数为所求
	ans := 0
	stack := make([]kV, 0)
	stack = append(stack, kV{nums[0], 0})
	for i := 1; i < len(nums); i++ {
		maxT := 0
		for len(stack) > 0 && nums[i] >= stack[len(stack)-1].v {
			maxT = max(maxT, stack[len(stack)-1].count)
			stack = stack[0 : len(stack)-1]
		}
		if len(stack) > 0 {
			maxT = maxT + 1
		}
		ans = max(ans, maxT)
		stack = append(stack, kV{nums[i], maxT})
	}
	return ans
}
