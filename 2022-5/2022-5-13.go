package main

import (
	"container/heap"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%k != 0 {
		return false
	}
	n := sum / k
	ans := make([][]int, k)
	sumC := make([]int, k)
	sort.Ints(nums)
	var dfs func(nums []int, ans [][]int) bool
	dfs = func(nums []int, ans [][]int) bool {
		for index := len(nums) - 1; index >= 0; index-- {
			flag := false
			val := nums[index]
			for i := 0; i < k; i++ {
				if sumC[i]+val > n {
					continue
				}
				if i > 0 && sumC[i] == sumC[i-1] {
					continue
				}
				ans[i] = append(ans[i], val)
				sumC[i] += val
				flag = flag || dfs(nums[0:index], ans)
				ans[i] = ans[i][0 : len(ans[i])-1]
				sumC[i] -= val
			}
			if flag == false {
				return false
			}
		}
		return true
	}
	return dfs(nums, ans)
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{val, nil, nil}
	}
	var travel func(root *TreeNode, val int)
	travel = func(root *TreeNode, val int) {
		if root.Val == val {
			return
		}
		if root.Val > val {
			if root.Left == nil {
				root.Left = &TreeNode{val, nil, nil}
				return
			}
			travel(root.Left, val)
		} else {
			if root.Right == nil {
				root.Right = &TreeNode{val, nil, nil}
				return
			}
			travel(root.Right, val)
		}
	}
	travel(root, val)
	return root
}

type KthLargest struct {
	h digit
	k int
}

func ConstructorK(k int, nums []int) KthLargest {
	var temp digit
	heap.Init(&temp)
	KL := KthLargest{temp, k}
	for _, val := range nums {
		KL.Add(val)
	}
	return KL
}

func (this *KthLargest) Add(val int) int {
	if len(this.h) >= this.k {
		if val > this.h[0] {
			this.h[0] = val
			heap.Fix(&this.h, 0)
		}
	} else {
		heap.Push(&this.h, val)
	}
	return this.h[0]
}

type digit []int

func (this digit) Len() int {
	return len(this)
}

func (this digit) Less(i, j int) bool {
	return this[i] > this[j]
}

func (this digit) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *digit) Pop() interface{} {
	old := *this
	num := old[len(old)-1]
	*this = old[0 : len(old)-1]
	return num
}

func (this *digit) Push(x interface{}) {
	*this = append(*this, x.(int))
}
