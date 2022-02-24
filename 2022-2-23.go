package main

import "math"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	nodeExistMap := make(map[int]int)
	nodeMap := make(map[int][]int)
	for _, val := range edges {
		nodeExistMap[val[0]]++
		nodeMap[val[0]] = append(nodeMap[val[0]], val[1])
		nodeExistMap[val[1]]++
		nodeMap[val[1]] = append(nodeMap[val[1]], val[0])
	}
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if nodeExistMap[i] == 1 {
			queue = append(queue, i)
		}
	}
	var ans []int
	for len(queue) > 0 {
		ans = make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[0]
			ans = append(ans, top)
			queue = queue[1:]
			for _, v := range nodeMap[top] {
				nodeExistMap[v]--
				if nodeExistMap[v] == 1 {
					queue = append(queue, v)
				}
			}
		}
	}
	return ans
}

func nthSuperUglyNumber(n int, primes []int) int {
	length := len(primes)
	pointers := make([]int, length)
	dp := make([]int, n+1)
	nums := make([]int, length)
	for i := range nums {
		nums[i] = 1
	}
	for i := 1; i <= n; i++ {
		minNum := math.MaxInt64
		for j := range nums {
			minNum = min(minNum, nums[j])
		}
		dp[i] = minNum
		for j := 0; j < length; j++ {
			if dp[i] == nums[j] {
				pointers[j]++
				nums[j] = dp[pointers[j]] * primes[j]
			}
		}
	}
	return dp[n]
}

func kthDistinct(arr []string, k int) string {
	m := make(map[string]int)
	for _, val := range arr {
		m[val]++
	}
	count := k
	for _, val := range arr {
		if m[val] == 1 {
			count--
		}
		if count == 0 {
			return val
		}
	}
	return ""
}

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
