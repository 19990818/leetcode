package main

func countElements(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	maxNum, minNum := nums[0], nums[0]
	for _, val := range nums {
		maxNum = max(maxNum, val)
		minNum = min(minNum, val)
	}
	ans := 0
	for _, val := range nums {
		if val != maxNum && val != minNum {
			ans++
		}
	}
	return ans
}

func rearrangeArray(nums []int) []int {
	pos, neg := make([]int, 0), make([]int, 0)
	ans := make([]int, 0)
	for _, val := range nums {
		if val > 0 {
			pos = append(pos, val)
		} else {
			neg = append(neg, val)
		}
	}
	for i := 0; i < len(pos); i++ {
		ans = append(ans, pos[i])
		ans = append(ans, neg[i])
	}
	return ans
}

func findLonely(nums []int) []int {
	m := make(map[int]int)
	ans := make([]int, 0)
	for _, val := range nums {
		m[val]++
	}
	for _, val := range nums {
		_, ok1 := m[val-1]
		_, ok2 := m[val+1]
		if (!ok1 && !ok2) && m[val] == 1 {
			ans = append(ans, val)
		}
	}
	return ans
}

func firstPalindrome(words []string) string {
	for _, val := range words {
		if isPandlim(val) {
			return val
		}
	}
	return ""
}
func isPandlim(s string) bool {
	for left, right := 0, len(s); left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}
