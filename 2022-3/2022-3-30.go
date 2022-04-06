package main

func magicalString(n int) int {
	if n == 1 {
		return 1
	}
	ans := 0
	s := make([]int, n)
	s[0] = 1
	j := 0
	cur := 1
	for i := 0; i < n; {
		count := s[j]
		if s[j] == 0 {
			count = cur
		}
		//fmt.Println(count)
		for i < n && count > 0 {
			if cur == 1 {
				ans++
			}
			s[i] = cur
			i++
			count--
		}
		cur ^= 3
		j++
	}
	//fmt.Println(s)
	return ans
}

func PredictTheWinner(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return PredictTheWinnerHelp(nums, 0, 0, sum, 1) == 1
}
func PredictTheWinnerHelp(nums []int, player1, player2, sum, flag int) int {
	if 2*player1 >= sum {
		return 1
	}
	if 2*player2 > sum {
		return 2
	}
	if flag == 1 {
		if PredictTheWinnerHelp(nums[1:], player1+nums[0], player2, sum, 2) == 1 {
			return 1
		}
		if PredictTheWinnerHelp(nums[0:len(nums)-1], player1+nums[len(nums)-1], player2, sum, 2) == 1 {
			return 1
		}
		return 2
	}
	if PredictTheWinnerHelp(nums[1:], player1, player2+nums[0], sum, 1) == 1 && PredictTheWinnerHelp(nums[0:len(nums)-1], player1, player2+nums[len(nums)-1], sum, 1) == 1 {
		return 1
	}
	return 2
}

func PredictTheWinnerDp(nums []int) bool {
	length := len(nums)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
		dp[i][i] = nums[i]
	}
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	return dp[0][length-1] >= 0
}

func findSubsequences(nums []int) [][]int {
	ans := make([][]int, 0)
	var dfs func(nums []int, cur []int)
	dfs = func(nums, cur []int) {
		if len(nums) == 0 {
			return
		}
		m2 := make(map[int]int)
		for i := 0; i < len(nums); i++ {
			if m2[nums[i]] != 0 {
				continue
			}
			m2[nums[i]] = 1
			if len(cur) == 0 {
				cur = append(cur, nums[i])
			} else {
				if nums[i] >= cur[len(cur)-1] {
					cur = append(cur, nums[i])
				} else {
					continue
				}
			}
			if len(cur) > 1 {
				temp := append([]int{}, cur...)
				ans = append(ans, temp)
			}
			dfs(nums[i+1:], cur)
			cur = cur[0 : len(cur)-1]
		}
	}
	dfs(nums, []int{})
	return ans
}
