package main

import "strings"

func decodeString(s string) string {
	numsStack, stringStack := make([]int, 0), make([]string, 0)
	for i := 0; i < len(s); {
		if s[i] >= '0' && s[i] <= '9' {
			tempNum := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				tempNum = 10*tempNum + int(s[i]-'0')
				i++
			}
			numsStack = append(numsStack, tempNum)
		} else if s[i] == '[' {
			i++
			var tempBuf strings.Builder
			for i < len(s) && s[i] <= 'z' && s[i] >= 'a' {
				tempBuf.WriteByte(s[i])
				i++
			}
			stringStack = append(stringStack, tempBuf.String())
		} else if s[i] == ']' {
			var tempStr strings.Builder
			for j := 0; j < numsStack[len(numsStack)-1]; j++ {
				tempStr.WriteString(stringStack[len(stringStack)-1])
			}
			numsStack = numsStack[0 : len(numsStack)-1]
			stringStack = stringStack[0 : len(stringStack)-1]
			if len(stringStack) != 0 {
				stringStack[len(stringStack)-1] += tempStr.String()
			} else {
				stringStack = append(stringStack, tempStr.String())
			}
			i++
		} else {
			var tempStr strings.Builder
			for i < len(s) && s[i] <= 'z' && s[i] >= 'a' {
				tempStr.WriteByte(s[i])
				i++
			}
			if len(stringStack) == 0 {
				stringStack = append(stringStack, tempStr.String())
			} else {
				stringStack[len(stringStack)-1] += tempStr.String()
			}
		}
	}
	return stringStack[0]
}

func longestSubstring(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[rune]int)
	for _, val := range s {
		m[val]++
	}
	for i := 0; i < len(s); i++ {
		if m[rune(s[i])] < k {
			return max(longestSubstring(s[0:i], k), longestSubstring(s[i+1:], k))
		}
	}
	return len(s)
}

func maxRotateFunction(nums []int) int {
	ans := 0
	sum := 0
	start := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		start += nums[i] * i
	}
	ans = start
	for i := 1; i < len(nums); i++ {
		start = start - sum + len(nums)*nums[i-1]
		ans = max(ans, start)
	}
	return ans
}
