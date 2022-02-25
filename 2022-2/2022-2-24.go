package main

func maxProduct3(words []string) int {
	hashMap := make([]int, len(words))
	for i, val := range words {
		for _, c := range val {
			hashMap[i] |= 1 << (c - 'a')
		}
	}
	ans := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if hashMap[i]&hashMap[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for _, val := range coins {
		if val <= amount {
			dp[val] = 1
		}

	}
	for i := 1; i <= amount; i++ {
		for _, val := range coins {
			if i >= val {
				if dp[i] != 0 {
					if dp[i-val] != 0 {
						dp[i] = min(dp[i], dp[i-val]+1)
					}

				} else {
					if dp[i-val] != 0 {
						dp[i] = dp[i-val] + 1
					}
				}
			}
		}
	}
	if dp[amount] == 0 && amount != 0 {
		return -1
	}
	return dp[amount]
}

func countVowelSubstrings(word string) int {
	ans := 0
	for i := 0; i < len(word)-4; i++ {
		for j := i + 4; j < len(word); j++ {
			if isVowelString(word[i : j+1]) {
				ans++
			}
		}
	}
	return ans
}
func isVowelString(s string) bool {
	m := make(map[rune]int)
	for _, val := range "aeiou" {
		m[val] = 0
	}
	for _, val := range s {
		if _, ok := m[val]; !ok {
			return false
		} else {
			m[val] = 1
		}
	}
	count := 0
	for _, val := range m {
		if val == 1 {
			count++
		}
	}
	return count == 5
}
