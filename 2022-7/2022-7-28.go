package main

import "sort"

func letterCombinations(digits string) []string {
	m := map[byte]string{'2': "abc", '3': "def", '4': "ghi", '5': "jkl", '6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz"}
	ans := make([]string, 0)
	var dfs func(idx int)
	temp := make([]byte, 0)
	dfs = func(idx int) {
		if idx == len(digits) {
			if len(temp) != 0 {
				ans = append(ans, string(temp))
			}
			return
		}
		for _, val := range m[digits[idx]] {
			temp = append(temp, byte(val))
			dfs(idx + 1)
			temp = temp[0 : len(temp)-1]
		}
	}
	dfs(0)
	return ans
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//快慢指针
	res := head
	slow, fast := head, head
	for n > 0 {
		n--
		fast = fast.Next
	}
	if fast == nil {
		return res.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return res
}

func isValid(s string) bool {
	//栈思想
	in := map[rune]int{'(': 1, '{': 1, '[': 1}
	out := map[rune]rune{')': '(', '}': '{', ']': '['}
	stack := make([]rune, 0)
	for _, val := range s {
		if in[val] == 1 {
			stack = append(stack, val)
		} else {
			if len(stack) == 0 || out[val] != stack[len(stack)-1] {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}
	return len(stack) == 0
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := new(ListNode)
	ans := res
	cur1, cur2 := list1, list2
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			res.Next = &ListNode{Val: cur1.Val}
			cur1 = cur1.Next
		} else {
			res.Next = &ListNode{Val: cur2.Val}
			cur2 = cur2.Next
		}
		res = res.Next
	}
	if cur1 != nil {
		res.Next = cur1
	}
	if cur2 != nil {
		res.Next = cur2
	}
	return ans.Next
}

func generateParenthesis2(n int) []string {
	//dfs生成括号 left计数左括号 cur计数右括号 右括号数不能大于左括号
	left := 0
	ans := make([]string, 0)
	m := make(map[string]int)
	temp := []byte{}
	var dfs func(cur int)
	dfs = func(cur int) {
		if left > n || cur > n {
			return
		}
		if cur == n && cur == left {
			if m[string(temp)] == 0 {
				m[string(temp)] = 1
				ans = append(ans, string(temp))
			}
			return
		}
		left++
		temp = append(temp, '(')
		dfs(cur)
		temp = temp[0 : len(temp)-1]
		left--
		if left > cur {
			temp = append(temp, ')')
			dfs(cur + 1)
			temp = temp[0 : len(temp)-1]
		}
	}
	dfs(0)
	return ans
}

func nextPermutation(nums []int) {
	//从后到前找到第一个会比后面的元素小的进行交换才可以得到紧邻的最小值
	right := len(nums) - 1
	stack := []int{right}
	pos := -1
	for i := right - 1; i >= 0; i-- {
		if nums[i] <= nums[stack[len(stack)-1]] {
			for len(stack) > 0 && nums[i] < nums[stack[len(stack)-1]] {
				pos = stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
			}
			if pos != -1 {
				nums[i], nums[pos] = nums[pos], nums[i]
				sort.Ints(nums[i+1:])
				return
			}
		}
		stack = append(stack, i)
	}
	sort.Ints(nums)
}

func search(nums []int, target int) int {
	//少考虑了一种情况 一个元素不存在比较的递增序列
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)>>1 + left
		//说明左边是递增的
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[left] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[right] > nums[mid] {
			//右边递增
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			left = mid + 1
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	//两次二分一次最后出现 一次第一次出现
	var binarySearch func(flag int) int
	binarySearch = func(flag int) int {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := (right-left)>>1 + left
			if nums[mid] < target {
				left = mid + 1
			} else if nums[mid] > target {
				right = mid - 1
			} else {
				if flag == 1 {
					if mid+1 >= len(nums) || nums[mid+1] != target {
						return mid
					}
					left = mid + 1
				} else {
					if mid-1 < 0 || nums[mid-1] != target {
						return mid
					}
					right = mid - 1
				}
			}
		}
		if len(nums) == 0 || left >= len(nums) || nums[left] != target {
			return -1
		}
		return left
	}
	return []int{binarySearch(-1), binarySearch(1)}
}

func lengthOfLIS(nums []int) int {
	//动态规划 以i结束的最长的串为多少
	dp := make([]int, len(nums))
	dp[0] = 1
	ans := 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	//fmt.Println(dp)
	return ans
}
