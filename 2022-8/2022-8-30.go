package main

import (
	"sort"
	"strings"
)

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	//比较每个值 如果当前根节点的值
	if val > root.Val {
		res := new(TreeNode)
		res.Val = val
		res.Left = root
		return res
	}
	//实际上一直取右子树
	cur := root
	for cur.Right != nil && val < cur.Right.Val {
		cur = cur.Right
	}
	temp := new(TreeNode)
	temp.Val = val
	pre := cur.Right
	cur.Right = temp
	temp.Left = pre
	return root
}

func numComponents(head *ListNode, nums []int) int {
	ans := 0
	m := make(map[int]int)
	for _, val := range nums {
		m[val] = 1
	}
	start := 0
	for head != nil {
		if m[head.Val] == 1 {
			if start == 0 {
				ans++
				start = 1
			}
		} else {
			start = 0
		}
		head = head.Next
	}
	return ans
}

func minimumLengthEncoding(words []string) int {
	m := make(map[string]int)
	for _, val := range words {
		flag := 0
		for key := range m {
			if len(key) >= len(val) && key[len(key)-len(val):] == val {
				flag = 1
				continue
			}
			if len(key) < len(val) && key == val[len(val)-len(key):] {
				delete(m, key)
				flag = 1
				m[val] = 1
			}
		}
		if flag == 0 {
			m[val] = 1
		}
	}
	ans := 0
	for key := range m {
		ans += len(key) + 1
	}
	return ans
}

func minimumLengthEncoding2(words []string) int {
	var reverseString func(a string) string
	reverseString = func(a string) string {
		var res strings.Builder
		for i := len(a) - 1; i >= 0; i-- {
			res.WriteByte(a[i])
		}
		return res.String()
	}
	for i, val := range words {
		words[i] = reverseString(val)
	}
	sort.Strings(words)
	ans := 0
	for i := 0; i < len(words)-1; i++ {
		n := len(words[i])
		if len(words[i+1]) > len(words[i]) && words[i+1][len(words[i+1])-n:] == words[i] {
			continue
		}
		ans += n + 1
	}
	return ans + len(words[len(words)-1]) + 1
}
