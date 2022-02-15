package main

import "strconv"

func getHint(secret string, guess string) string {
	bulls := 0
	cows := 0
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		}
	}
	secretMap := make(map[rune]int)
	guessMap := make(map[rune]int)
	for _, val := range secret {
		secretMap[val]++
	}
	for _, val := range guess {
		guessMap[val]++
	}
	for _, val := range secret {
		secretMap[val]++
	}
	for key := range secretMap {
		cows += min(secretMap[key], guessMap[key])
	}
	cows -= bulls
	ans := ""
	ans += strconv.Itoa(bulls)
	ans += "A"
	ans += strconv.Itoa(cows)
	ans += "B"
	return ans
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	result := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			} else {
				dp[i] = max(dp[i], 1)
			}
		}
		result = max(result, dp[i])
	}
	//fmt.Println(dp)
	return result
}

func findFinalValue(nums []int, original int) int {
	mNums := make(map[int]int)
	for _, val := range nums {
		mNums[val] = 1
	}
	_, ok := mNums[original]
	for ok {
		original *= 2
		_, ok = mNums[original]
	}
	return original
}
