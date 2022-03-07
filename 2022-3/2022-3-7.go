package main

import (
	"sort"
	"strconv"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([][]int, len(nums))
	dp[0] = nums[0:1]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 {
				if len(dp[i]) < len(dp[j])+1 {
					dp[i] = append([]int{}, dp[j]...)
				}
			}
		}
		dp[i] = append(dp[i], nums[i])
	}
	var ans []int
	for _, val := range dp {
		if len(val) > len(ans) {
			ans = val
		}
	}
	//fmt.Println(dp)
	return ans
}

const mod = 1337

func superPow(a int, b []int) int {
	ans := 1
	for i := len(b) - 1; i >= 0; i-- {
		ans = ans * pow(a, b[i]) % mod
		a = pow(a, 10)
	}
	return ans
}
func pow(x, n int) int {
	ans := 1
	for ; n > 0; n /= 2 {
		if n&1 > 0 {
			ans = ans * x % mod
		}
		x = x * x % mod
	}
	return ans
}

func findNthDigit(n int) int {
	dp := make([]int, 10)
	dp[0] = 9
	for i := 1; i < 10; i++ {
		dp[i] = dp[i-1] * 10 / i * (i + 1)
	}
	i := 0
	for ; n >= dp[i]; i++ {
		n -= dp[i]
	}
	if n == 0 {
		return 9
	}
	//fmt.Println(dp)
	base := pow(10, i) - 1
	base += n / (i + 1)
	//fmt.Println(i,base)
	s := strconv.Itoa(base)
	if n%(i+1) == 0 {
		return int(s[len(s)-1] - '0')
	}
	s2 := strconv.Itoa(base + 1)
	return int(s2[n%(i+1)-1] - '0')
}
