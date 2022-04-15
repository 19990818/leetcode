package main

import "sort"

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	ans := 0
	var binaryfind func(nums []int, target int) int
	binaryfind = func(nums []int, target int) int {
		left, right := 0, len(nums)
		if target > nums[len(nums)-1] {
			return len(nums) - 1
		}
		if target <= nums[0] {
			return -1
		}
		for left < right {
			mid := (right-left)>>1 + left
			if nums[mid] >= target {
				right = mid - 1
			} else {
				left = mid
				if left == right-1 {
					if target > nums[right] {
						return right
					}
					break
				}
			}
			//fmt.Println(left,right)
		}
		return left
	}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			k := binaryfind(nums[j+1:], nums[i]+nums[j])
			//fmt.Println(k)
			ans += k + 1

		}
	}
	return ans
}

func leastInterval(tasks []byte, n int) int {
	m := make([]int, 26)
	for _, val := range tasks {
		m[val-'A']++
	}
	maxNum := 0
	for _, val := range m {
		maxNum = max(maxNum, val)
	}
	maxCount := 0
	for _, val := range m {
		if val == maxNum {
			maxCount++
		}
	}
	return max((maxNum-1)*(n+1)+maxCount, len(tasks))
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	order := make([][]*TreeNode, 0)
	if depth == 1 {
		res := new(TreeNode)
		res.Val = val
		res.Left = root
		return res
	}
	queue := append([]*TreeNode{}, root)
	for len(queue) > 0 {
		temp := make([]*TreeNode, 0)
		order = append(order, queue)
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				temp = append(temp, cur.Left)
			}
			if cur.Right != nil {
				temp = append(temp, cur.Right)
			}
		}
		queue = temp
	}
	for _, val2 := range order[depth-2] {
		left, right := val2.Left, val2.Right
		nodeLeft, nodeRight := new(TreeNode), new(TreeNode)
		nodeLeft.Val, nodeRight.Val = val, val
		val2.Left, val2.Right = nodeLeft, nodeRight
		nodeLeft.Left = left
		nodeRight.Right = right
	}
	return root
}
