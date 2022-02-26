package main

import (
	"sort"
)

func wiggleSort(nums []int) {
	sort.Ints(nums)
	ans := make([]int, 0)
	for i, j := (len(nums)+1)/2-1, len(nums)-1; i >= 0; i, j = i-1, j-1 {
		if i >= 0 {
			ans = append(ans, nums[i])
		}
		if j >= (len(nums)+1)/2 {
			ans = append(ans, nums[j])
		}
	}
	copy(nums, ans)
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, even := head, head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
