package main

func longestPalindromeSubseq(s string) int {
	length := len(s)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
		dp[i][i] = 1
	}
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	// fmt.Println(dp)
	return dp[0][length-1]
}

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, val := range coins {
		for j := 1; j < amount+1; j++ {
			if j >= val {
				dp[j] += dp[j-val]
			}
		}
	}
	return dp[amount]
}

func minBitFlips(start int, goal int) int {
	ans := 0
	mid := start ^ goal
	for i := 32; i >= 0; i-- {
		if mid>>i&1 == 1 {
			ans++
		}
	}
	return ans
}

func triangularSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	for {
		temp := make([]int, 0)
		for i := 1; i < len(nums); i++ {
			temp = append(temp, (nums[i]+nums[i-1])%10)
		}
		nums = temp
		//fmt.Println(nums,len(nums))
		if len(nums) == 1 {
			break
		}
	}
	return nums[0]
}

func numberOfWays(s string) int64 {
	var f func(s string, start byte) int64
	f = func(s string, start byte) int64 {
		i := 0
		for ; i < len(s); i++ {
			if s[i] == start {
				break
			}
		}
		count1 := int64(0)
		count2 := int64(0)
		ans := int64(0)
		cur := byte(start)
		step := 0
		for ; i < len(s); i++ {
			if s[i] == cur {
				count1++
				step = step | 1
				if step == 3 {
					ans += count2
				}
			} else {
				count2 += count1
				step = step | 2
			}
		}
		return ans
	}
	return f(s, '0') + f(s, '1')
}
