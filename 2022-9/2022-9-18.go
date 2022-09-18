package main

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return 2 * n
}

func longestContinuousSubstring(s string) int {
	ans := 1
	cur := s[0]
	cnt := 1
	for i := 1; i < len(s); i++ {
		if s[i] != cur+1 {
			ans = max(ans, cnt)
			cnt = 1
		} else {
			cnt++
		}
		cur = s[i]
	}
	ans = max(ans, cnt)
	return ans
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	queue := []*TreeNode{root}
	cnt := 0
	for {
		temp := make([]*TreeNode, 0)
		cnt++
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.Right != nil {
				temp = append(temp, cur.Right)
			}
			if cur.Left != nil {
				temp = append(temp, cur.Left)
			}
		}
		if cnt%2 == 1 {
			for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
				temp[i].Val, temp[j].Val = temp[j].Val, temp[i].Val
			}
		}
		if len(temp) == 0 {
			break
		}
		queue = temp
	}
	return root
}
