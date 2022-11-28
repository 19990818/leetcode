package main

func pivotInteger(n int) int {
	sum := n * (n + 1) / 2
	subSum := 0
	for i := 1; i <= n; i++ {
		subSum += i
		if subSum == sum-subSum+i {
			return i
		}
	}
	return -1
}

func appendCharacters(s string, t string) int {
	i := 0
	for j := range s {
		if s[j] == t[i] {
			i++
		}
		if i >= len(t) {
			break
		}
	}
	return len(t) - i
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
	stack := make([]int, 0)
	for head != nil {
		for len(stack) > 0 && head.Val > stack[len(stack)-1] {
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, head.Val)
		head = head.Next
	}
	res := new(ListNode)
	cur := res
	for i, val := range stack {
		cur.Val = val
		if i < len(stack)-1 {
			cur.Next = new(ListNode)
		}

		cur = cur.Next
	}
	return res
}

func countSubarrays(nums []int, k int) int {
	// 我们首先找到k的位置
	posk := 0
	for i, num := range nums {
		if num == k {
			posk = i
			break
		}
	}
	diff := make([]int, len(nums))
	diff[posk] = 0
	res := 1
	leftM := make(map[int]int)
	rightM := make(map[int]int)
	for left := posk - 1; left >= 0; left-- {
		if nums[left] > k {
			diff[left] = diff[left+1] + 1
		} else {
			diff[left] = diff[left+1] - 1
		}
		if diff[left] == 0||diff[left]==1 {
			res++
		}
		leftM[diff[left]]++
	}
	for right := posk + 1; right < len(nums); right++ {
		if nums[right] > k {
			diff[right] = diff[right-1] + 1
		} else {
			diff[right] = diff[right-1] - 1
		}
		if diff[right] == 0||diff[right]==1 {
			res++
		}
		rightM[diff[right]]++
	}
	for k := range leftM {
		res += leftM[k] * rightM[-k]
		res += leftM[k] * rightM[-k+1]
	}
	return res
}
