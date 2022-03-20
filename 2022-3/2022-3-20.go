package main

func countHillValley(nums []int) int {
	ans := 0
	temp := make([]int, 0)
	temp = append(temp, nums[0])
	cur := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != cur {
			temp = append(temp, nums[i])
			cur = nums[i]
		}
	}
	for i := 1; i < len(temp)-1; i++ {
		if (temp[i]-temp[i-1])*(temp[i]-temp[i+1]) > 0 {
			ans++
		}
	}
	return ans
}

func countCollisions(directions string) int {
	ans := 0
	flag := false
	for _, val := range directions {
		if val == 'R' || val == 'S' {
			flag = true
		} else if flag {
			ans++
		}
	}
	flag = false
	for idx := len(directions) - 1; idx >= 0; idx-- {
		if directions[idx] == 'L' || directions[idx] == 'S' {
			flag = true
		} else if flag {
			ans++
		}
	}
	return ans
}

func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	// 01 背包问题 alicearrows表示体积 下标表示重量
	dp := make([]int, numArrows+1)
	dp[0] = 0

	for i := 0; i < len(aliceArrows); i++ {
		for j := numArrows; j > aliceArrows[i]; j-- {
			if dp[j] < dp[j-aliceArrows[i]-1]+i {
				dp[j] = dp[j-aliceArrows[i]-1] + i
			}
		}
	}
	target := dp[numArrows]
	var dfs func(numArrows int, aliceArrows []int, target int) []int
	dfs = func(numArrows int, aliceArrows []int, target int) []int {
		//fmt.Println(numArrows,target,aliceArrows)
		if numArrows < 1 {
			return []int{}
		}
		for idx, val := range aliceArrows {
			if idx == target && numArrows >= val+1 {
				return []int{idx}
			}
		}
		var ans []int
		for idx := len(aliceArrows) - 1; idx > 0; idx-- {
			ans = make([]int, 0)
			temp := make([]int, 0)
			ans = append(ans, idx)
			temp = append(temp, aliceArrows[0:idx]...)
			temp = append(temp, aliceArrows[idx+1:]...)
			subPath := dfs(numArrows-aliceArrows[idx]-1, temp, target-idx)
			if len(subPath) == 0 {
				ans = []int{}
			} else {
				ans = append(ans, subPath...)
				return ans
			}
		}
		//fmt.Println(ans)
		return ans
	}
	path := dfs(numArrows, aliceArrows, target)
	//fmt.Println(path)
	sum := numArrows
	ans := make([]int, len(aliceArrows))
	for _, val := range path {
		ans[val] = aliceArrows[val] + 1
		sum = sum - aliceArrows[val] - 1
	}
	ans[0] = sum
	//fmt.Println(dp[numArrows], path[numArrows])
	return ans
}

//使用枚举完成
func maximumBobPoints2(numArrows int, aliceArrows []int) []int {
	maxScore := -1
	ans := make([]int, len(aliceArrows))
	for i := 1; i < 1<<len(aliceArrows); i++ {
		arrow := 0
		score := 0
		temp := make([]int, len(aliceArrows))
		for j := 0; j < len(aliceArrows); j++ {
			if i>>j&1 == 1 {
				temp[j] = 1
				score += j
				arrow += aliceArrows[j] + 1
			}
			if arrow > numArrows {
				continue
			}
			if score > maxScore {
				maxScore = score
				temp[0] = numArrows - arrow
				ans = temp
			}
		}
	}
	return ans
}
