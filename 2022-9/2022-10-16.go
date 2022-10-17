package main

func findMaxK(nums []int) int {
	ans := -1
	m := make(map[int]int)
	for _, val := range nums {
		m[val] = 1
	}
	for _, val := range nums {
		if val > 0 && m[-val] == 1 {
			ans = max(ans, val)
		}
	}
	return ans
}

func countDistinctIntegers(nums []int) int {
	m := make(map[int]int)
	for _, val := range nums {
		m[val] = 1
	}
	m2 := make(map[int]int)
	reverseInt := func(a int) int {
		res := 0
		for a > 0 {
			temp := a % 10
			a /= 10
			res = res*10 + temp
		}
		return res
	}
	for k := range m {
		m2[k] = 1
		m2[reverseInt(k)] = 1
	}
	return len(m2)
}

func sumOfNumberAndReverse(num int) bool {
	reverseInt := func(a int) int {
		res := 0
		for a > 0 {
			temp := a % 10
			a /= 10
			res = res*10 + temp
		}
		return res
	}
	for i := 0; i <= num; i++ {
		if i+reverseInt(i) == num {
			return true
		}
	}
	return false
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	dp := make([][]int64, len(nums))
	for i := range dp {
		dp[i] = make([]int64, 4)
	}
	if nums[0] == maxK && nums[0] == minK {
		dp[0][3] = 1
	} else if nums[0] == maxK {
		dp[0][2] = 1
	} else if nums[0] == minK {
		dp[0][1] = 1
	} else {
		if nums[0] < maxK && nums[0] > minK {
			dp[0][0] = 1
		}

	}
	for i := 1; i < len(nums); i++ {
		if nums[i] <= maxK && nums[i] >= minK {
			dp[i][3] = dp[i-1][3]
			dp[i][2] = dp[i-1][2]
			dp[i][1] = dp[i-1][1]
			dp[i][0] = dp[i-1][0]
		}
		if nums[i] == maxK && nums[i] == minK {
			dp[i][3] += dp[i-1][0]
			dp[i][3] += dp[i-1][1]
			dp[i][3] += dp[i-1][2]
			dp[i][3] += 1
			dp[i][1], dp[i][2], dp[i][0] = 0, 0, 0
		} else if nums[i] == maxK {
			dp[i][3] += dp[i-1][1]
			dp[i][2] += dp[i-1][0]
			dp[i][2] += 1
			dp[i][0], dp[i][1] = 0, 0
		} else if nums[i] == minK {
			dp[i][3] += dp[i-1][2]
			dp[i][1] += dp[i-1][0]
			dp[i][1] += 1
			dp[i][0], dp[i][2] = 0, 0
		} else if nums[i] < maxK && nums[i] > minK {
			dp[i][0] += 1
		} else {
			for j := 0; j < 4; j++ {
				dp[i][j] = 0
			}
		}
	}
	//fmt.Println(dp)
	ans := int64(0)
	for _, val := range dp {
		ans += val[3]
	}
	return ans
}
