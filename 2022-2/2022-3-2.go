package main

import (
	"reflect"
	"strconv"
	"strings"
)

func transformArray(arr []int) []int {
	flag := 1
	for flag == 1 {
		temp := make([]int, 0)
		temp = append(temp, arr...)
		flag = 0
		for i := 1; i < len(arr)-1; i++ {
			if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
				temp[i]--
				flag = 1
			}
			if arr[i] < arr[i-1] && arr[i] < arr[i+1] {
				temp[i]++
				flag = 1
			}
		}
		arr = temp
	}
	return arr
}

func toHexspeak(num string) string {
	numDigit, _ := strconv.Atoi(num)
	ans := make([]rune, 0)
	for numDigit > 0 {
		rel := numDigit % 16
		if rel > 1 && rel < 10 {
			return "ERROR"
		}
		if rel == 1 {
			ans = append(ans, 'I')
		} else if rel == 0 {
			ans = append(ans, 'O')
		} else {
			ans = append(ans, rune('A'+(rel-10)))
		}
		numDigit /= 16
	}
	for left, right := 0, len(ans)-1; left < right; left, right = left+1, right-1 {
		ans[left], ans[right] = ans[right], ans[left]
	}
	var ansS strings.Builder
	for _, val := range ans {
		ansS.WriteRune(val)
	}
	return ansS.String()
}

func countElements(arr []int) int {
	m := make(map[int]int)
	for _, val := range arr {
		m[val] = 1
	}
	ans := 0
	for _, val := range arr {
		if _, ok := m[val+1]; ok {
			ans++
		}
	}
	return ans
}

func stringShift(s string, shift [][]int) string {
	ans := 0
	for _, val := range shift {
		if val[0] == 0 {
			ans += val[1]
		} else {
			ans -= val[1]
		}
	}
	ans = ans % len(s)
	//fmt.Println(ans)
	if ans < 0 {
		return s[len(s)+ans:] + s[0:len(s)+ans]
	}
	if ans == 0 {
		return s
	}
	return s[ans:] + s[0:ans]
}

func getLonelyNodes(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0)
	if root.Left != nil {
		if root.Right == nil {
			ans = append(ans, root.Left.Val)
		}
		ans = append(ans, getLonelyNodes(root.Left)...)
	}
	if root.Right != nil {
		if root.Left == nil {
			ans = append(ans, root.Right.Val)
		}
		ans = append(ans, getLonelyNodes(root.Right)...)
	}
	return ans
}

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	cur := head
	count := 0
	temp := head
	for cur != nil {
		count++
		if count%(m+n) == m {
			temp = cur
		}
		cur = cur.Next
		if count%(m+n) == 0 {
			temp.Next = cur
		}
	}
	if count%(m+n) > m {
		temp.Next = nil
	}
	return head
}

func largestSubarray(nums []int, k int) []int {
	ans := make([]int, k)
	for i := 0; i < len(nums)-k+1; i++ {
		ans = compareArr(ans, nums[i:i+k])
	}
	return ans
}
func compareArr(a, b []int) []int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return b
		}
		if a[i] > b[i] {
			return a
		}
	}
	if len(a) > len(b) {
		return a
	}
	return b
}

func badSensor(sensor1 []int, sensor2 []int) int {
	for i := 0; i < len(sensor1)-1; i++ {
		if sensor1[i] != sensor2[i] {
			if reflect.DeepEqual(sensor2[i:len(sensor2)-1], sensor1[i+1:]) && reflect.DeepEqual(sensor1[i:len(sensor1)-1], sensor2[i+1:]) {
				return -1
			}
			if reflect.DeepEqual(sensor1[i:len(sensor1)-1], sensor2[i+1:]) {
				return 1
			}
			if reflect.DeepEqual(sensor2[i:len(sensor2)-1], sensor1[i+1:]) {
				return 2
			}
		}
	}
	return -1
}

func isDecomposable(s string) bool {
	if len(s)%3 != 2 {
		return false
	}
	temp := s[0]
	count := 1
	flag := 0
	for i := 1; i < len(s); i++ {
		if s[i] == temp {
			count++
		} else {
			temp = s[i]
			if count%3 != 0 {
				if count%3 == 2 && flag == 0 {
					flag = 1
				} else {
					return false
				}
			}
			count = 1
		}
	}
	if count%3 != 0 {
		if count%3 == 2 && flag == 0 {
			flag = 1
		} else {
			return false
		}
	}
	return flag == 1
}
