package main

import "strings"

func countTime(time string) int {
	if !strings.Contains(time, "?") {
		return 1
	}
	ans := 1
	if time[0] == '?' {
		if time[1] <= '3' && time[1] >= '0' {
			ans *= 3
		} else if time[1] == '?' {
			ans = 24
		} else {
			ans *= 2
		}
	} else if time[1] == '?' {
		if time[0] == '2' {
			ans *= 4
		} else {
			ans *= 10
		}
	}
	if time[3] == '?' {
		ans *= 6
	}
	if time[4] == '?' {
		ans *= 10
	}
	return ans
}
func productQueries(n int, queries [][]int) []int {
	powers := make([]int, 0)
	for i := 0; i < 32; i++ {
		if n>>i&1 != 0 {
			powers = append(powers, 1<<i)
		}
	}
	mod := int(1e9 + 7)
	ans := make([]int, 0)
	for _, val := range queries {
		temp := 1
		for i := val[0]; i <= val[1]; i++ {
			temp = (temp * powers[i]) % mod
		}
		ans = append(ans, temp)
	}
	return ans
}

func minimizeArrayValue(nums []int) int {
	ans := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		ans = max(ans, (sum+i)/(i+1))
	}
	return ans
}