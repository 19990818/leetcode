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

func productExceptSelf(nums []int) []int {
	sum := 1
	sum2 := 1
	count := 0
	ans := make([]int, 0)
	for _, val := range nums {
		if val == 0 {
			sum2 *= 1
			count++
		} else {
			sum2 *= val
		}
		sum *= val
	}
	if count > 1 {
		sum2 = 0
	}
	flag1 := 0
	if sum < 0 {
		flag1 = 1
		sum = -sum
	}
	for _, val := range nums {
		t := 0
		temp := sum
		flag2 := 0
		if val < 0 {
			flag2 = 1
			val = -val
		}
		if val == 0 {
			ans = append(ans, sum2)
			continue
		}
		for i := 31; i >= 0; i-- {
			//fmt.Println(val<<i,temp)
			if temp >= (val << i) {
				temp -= val << i
				t += 1 << i
			}
		}
		if flag1^flag2 == 1 {
			t = -t
		}
		ans = append(ans, t)
	}
	return ans
}

func singleNumber3(nums []int) []int {
	num1, num2 := 0, 0
	xorm := 0
	for _, val := range nums {
		xorm ^= val
	}
	//得到第一位两者不相同的标志位 用于分为两类
	pos := xorm & (-xorm)
	for _, val := range nums {
		if val&pos != 0 {
			num1 ^= val
		} else {
			num2 ^= val
		}
	}
	return []int{num1, num2}
}
