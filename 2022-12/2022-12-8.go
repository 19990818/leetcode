package main

import "math/bits"

func checkPowersOfThree(n int) bool {
	pow3 := make([]int, 16)
	pow3[0] = 1
	for i := 1; i < 16; i++ {
		pow3[i] = pow3[i-1] * 3
	}
	for i := 15; i >= 0; i-- {
		if n >= pow3[i] {
			n -= pow3[i]
			if n == 0 {
				return true
			}
		}
	}
	return false
}

func maxLength(arr []string) int {
	masks := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		mask := 0
		for _, c := range arr[i] {
			if (1<<int(c-'a'))&mask == 0 {
				mask |= 1 << int(c-'a')
			} else {
				mask = 0
				break
			}
		}
		masks = append(masks, mask)
	}
	var dfs func(pos, mask int)
	ans := 0
	dfs = func(pos, mask int) {
		if pos == len(masks) {
			ans = max(ans, bits.OnesCount(uint(mask)))
			return
		}
		if mask&masks[pos] == 0 {
			dfs(pos+1, mask|masks[pos])
		}
		dfs(pos+1, mask)
	}
	dfs(0, 0)
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := root.Val
	q := []*TreeNode{root}
	for len(q) > 0 {
		temp := q
		q = []*TreeNode{}
		tempRes := 0
		for len(temp) > 0 {
			cur := temp[0]
			temp = temp[1:]
			tempRes += cur.Val
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
		res = tempRes
	}
	return res
}
