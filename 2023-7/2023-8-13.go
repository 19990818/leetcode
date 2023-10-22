package main

import (
	"sort"
	"strconv"
)

func maxSum(nums []int) int {
	m := make(map[int][]int)
	sort.Ints(nums)
	for _, v := range nums {
		s := strconv.Itoa(v)
		m[get(s)] = append(m[get(s)], v)
	}
	res := -1
	for _, v := range m {
		if len(v) > 1 {
			res = max(res, v[len(v)-1]+v[len(v)-2])
		}
	}
	return res
}
func get(a string) int {
	res := 0
	for _, v := range a {
		if int(v-'0') > res {
			res = int(v - '0')
		}
	}
	return res
}

func doubleIt(head *ListNode) *ListNode {
	nums := make([]int, 0)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	c := 0
	var res *ListNode
	for i := len(nums) - 1; i >= 0; i-- {
		t := new(ListNode)
		t.Val = (c + 2*nums[i]) % 10
		c = (c + 2*nums[i]) / 10
		t.Next = res
		res = t
	}
	if c > 0 {
		t := new(ListNode)
		t.Val = c
		t.Next = res
		res = t
	}
	return res
}
