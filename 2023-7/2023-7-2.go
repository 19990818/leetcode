package main

import (
	"math/bits"
	"strconv"
)

func countBeautifulPairs(nums []int) int {
	res := 0
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if gcd(int(strconv.Itoa(nums[i])[0]-'0'), nums[j]%10) == 1 {
				res++
			}
		}
	}
	return res
}
func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func makeTheIntegerZero(num1 int, num2 int) int {
	for i := 0; num1 > 0; i++ {
		num1 -= num2
		if bits.OnesCount(uint(num1)) <= i && num1 >= i {
			return i
		}
	}
	return -1
}

func numberOfGoodSubarraySplits(nums []int) int {
	mod := int(1e9 + 7)
	res := 1
	flag := false
	cnt := 0
	for _, v := range nums {
		if v == 1 {
			if flag {
				res = (res * cnt) % mod
			}
			cnt = 1
			flag = true
		} else {
			cnt++
		}
	}
	if !flag {
		return 0
	}
	return res
}

func longestAlternatingSubarray(nums []int, threshold int) int {
	res := 0
	cnt := 0
	start := 0
	flag := false
	for i, v := range nums {
		if v%2 == 0 && v <= threshold {
			if (i-start)%2 == 0 {
				cnt++
			} else {
				start = i
				cnt = 1
			}
			flag = true
		} else if flag && (i-start)%2 == 1 && v <= threshold {
			cnt++
		} else {
			flag = false
			cnt = 0
		}
		res = max(res, cnt)
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
