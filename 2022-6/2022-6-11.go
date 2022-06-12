package main

import (
	"sort"
)

func minFlipsMonoIncr(s string) int {
	count1sum, count0sum := 0, 0
	n := len(s)
	for _, val := range s {
		if val == '1' {
			count1sum++
		} else {
			count0sum++
		}
	}
	ans := n - count1sum
	left1sum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			left1sum++
		}
		right1sum := count1sum - left1sum
		ans = min(ans, left1sum+n-i-1-right1sum)
	}
	return ans
}

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	lowwer, higher := false, false
	digit, special := false, false
	specialStr := "!@#$%^&*()-+"
	m := make(map[byte]int)
	for i := range specialStr {
		m[specialStr[i]] = 1
	}
	for i := 0; i < len(password); i++ {
		if password[i] <= 'z' && password[i] >= 'a' {
			lowwer = true
		}
		if password[i] <= 'Z' && password[i] >= 'A' {
			higher = true
		}
		if password[i] <= '9' && password[i] >= '0' {
			digit = true
		}
		if m[password[i]] == 1 {
			special = true
		}
		if i != len(password)-1 {
			if password[i] == password[i+1] {
				return false
			}
		}
	}
	return lowwer && higher && digit && special
}

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	ans := make([]int, 0)
	for _, val := range spells {
		left, right := 0, len(potions)
		for left < right {
			mid := (right-left)>>1 + left
			if int64(val*potions[mid]) >= success {
				right = mid
			} else {
				left = mid + 1
			}
		}
		ans = append(ans, len(potions)-left)
	}
	return ans
}

func countSubarrays(nums []int, k int64) int64 {
	// 01下标的选择，通常如果需要i-1项 我们直接多分配一个
	// 当绕不明白的时候，将逻辑拆分开
	sum := make([]int64, len(nums)+1)
	for i, val := range nums {
		sum[i+1] = sum[i] + int64(val)
	}
	i, j := 1, 1
	ans := int64(0)
	for ; i <= len(nums); i++ {
		for j <= i && (sum[i]-sum[j-1])*int64(i-j+1) >= k {
			j++
		}
		ans += int64(i - j + 1)
	}
	return ans
}
