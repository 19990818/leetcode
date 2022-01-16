package main

import (
	"sort"
	"strings"
)

func capitalizeTitle(title string) string {
	titleArr := strings.Split(title, " ")
	ans := ""
	for index, val := range titleArr {
		if len(val) < 3 {
			ans += strings.ToLower(val)
		} else {
			temp := strings.ToLower(val)
			temp = getHead(temp)
			ans += temp
		}
		if index != len(titleArr)-1 {
			ans += " "
		}
	}
	return ans
}
func getHead(a string) string {
	ans := ""
	ans += string(a[0] + 'A' - 'a')
	ans += a[1:]
	return ans
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	ans := make([]int, 0)
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	left, right := 0, len(ans)-1
	res := 0
	for ; left < right; left, right = left+1, right-1 {
		if ans[left]+ans[right] > res {
			res = ans[left] + ans[right]
		}
	}
	return res
}

func longestPalindrome(words []string) int {
	ans := 0
	wordsMap := make(map[string]int)
	for _, val := range words {
		if _, ok1 := wordsMap[val]; ok1 {
			wordsMap[val]++
		} else {
			wordsMap[val] = 1
		}
	}
	flag := true
	flagMap := make(map[string]bool)
	for key, _ := range wordsMap {
		if _, ok := flagMap[reveseString(key)]; ok {
			continue
		}
		flagMap[key] = true
		t := min(wordsMap[key], wordsMap[reveseString(key)])
		if key == reveseString(key) {
			ans += (t / 2) * 4
		} else {
			ans += t * 4
		}

		if judgePanli(key) && (t%2 == 1) && flag {
			ans += 2
			flag = false
		}
	}
	return ans
}
func judgePanli(a string) bool {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		if a[left] != a[right] {
			return false
		}
	}
	return true
}
func reveseString(a string) string {
	ans := ""
	for i := len(a) - 1; i >= 0; i-- {
		ans += string(a[i])
	}
	return ans
}

func checkValid(matrix [][]int) bool {
	matrix2 := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		matrix2[i] = make([]int, len(matrix))
	}
	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {
			matrix2[j][i] = matrix[i][j]
		}
	}
	return checkFull(matrix2) && checkFull(matrix)
}
func checkFull(matrix [][]int) bool {
	n := len(matrix)
	for _, val := range matrix {
		sort.Ints(val)
		flag := make([]int, n)
		for _, val2 := range val {
			if val2 > n || val2 <= 0 {
				return false
			}
			flag[val2-1] = 1
		}
		for i := 0; i < n; i++ {
			if flag[i] == 0 {
				return false
			}
		}
	}
	return true
}

func minSwaps(nums []int) int {
	count := 0
	for _, val := range nums {
		if val == 1 {
			count++
		}
	}
	ans := count
	temp := count
	j := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			for ; j < count; j++ {
				if nums[j%len(nums)] == 1 {
					temp--
				}
			}
		} else {
			if nums[i-1] == 1 {
				temp++
			}
			if nums[j%len(nums)] == 1 {
				temp--
			}
			j++
		}
		ans = min(ans, temp)
	}
	return ans
}

func wordCount(startWords []string, targetWords []string) int {
	ans := 0
	startMap := make(map[string]int)
	for _, val := range startWords {
		temp := getMapKey(val)
		startMap[temp] = 1
	}
	for _, val := range targetWords {
		temp := getMapKey(val)
		for i := 0; i < len(temp); i++ {
			temp2 := ""
			temp2 = temp[0:i]
			temp2 += temp[i+1:]
			if _, ok := startMap[temp2]; ok {
				ans++
				break
			}
		}
	}
	return ans
}
func getMapKey(val string) string {
	tArr := make([]int, 26)
	for _, val2 := range val {
		tArr[val2-'a'] = 1
	}
	temp := ""
	for i := 0; i < 26; i++ {
		if tArr[i] == 1 {
			temp += string(i + 'a')
		}
	}
	return temp
}

func divideString(s string, k int, fill byte) []string {
	ans := make([]string, 0)
	for i := 0; i < len(s); i = i + k {
		if i+k > len(s) {
			temp := s[i:len(s)]
			for j := len(s); j < i+k; j++ {
				temp += string(fill)
			}
			ans = append(ans, temp)
		} else {
			ans = append(ans, s[i:i+k])
		}
	}
	return ans
}

func minMoves2(target int, maxDoubles int) int {
	ans := 0
	for i := target; i > 1; {
		if i%2 == 0 && maxDoubles > 0 {
			i = i / 2
			maxDoubles -= 1
			ans += 1
		} else {
			i--
			ans += 1
		}
	}
	return ans
}

func mostPoints(questions [][]int) int64 {
	dp := make([]int64, len(questions))
	dp[len(questions)-1] = int64(questions[len(questions)-1][0])
	for i := len(questions) - 2; i >= 0; i-- {
		if i+questions[i][1]+1 >= len(questions) {
			dp[i] = max64(int64(questions[i][0]), dp[i+1])
		} else {
			dp[i] = max64(int64(questions[i][0])+dp[i+questions[i][1]+1], dp[i+1])
		}
		//fmt.Println(dp[i])
	}
	return dp[0]
}
func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maxRunTime(n int, batteries []int) int64 {
	if len(batteries) < n {
		return 0
	}
	ans := int64(0)
	for len(batteries) >= n {
		sort.Ints(batteries)
		temp := make([]int, 0)
		for i := len(batteries) - 1; i >= len(batteries)-n; i-- {
			batteries[i] -= batteries[len(batteries)-n]
			ans += int64(batteries[len(batteries)-n])
			if batteries[i] != 0 {
				temp = append(temp, batteries[i])
			}
		}
		temp = append(temp, batteries[0:len(batteries)-n+1]...)
		batteries = temp
	}
	return ans
}
