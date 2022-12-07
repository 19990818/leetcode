package main

import "sort"

func minOperations2(nums1 []int, nums2 []int) int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	sum1, sum2 := sum(nums1), sum(nums2)
	if sum1 == sum2 {
		return 0
	}
	diff := sum2 - sum1
	if sum1 > sum2 {
		nums1, nums2 = nums2, nums1
		diff = sum1 - sum2
	}
	res := 0
	for left, right := 0, len(nums2)-1; left < len(nums1) || right >= 0; {
		if right == -1 || (left < len(nums1) && nums1[left]+nums2[right] < 7) {
			diff -= (6 - nums1[left])
			left++
		} else {
			diff -= (nums2[right] - 1)
			right--
		}
		res++
		if diff <= 0 {
			return res
		}
	}
	return -1
}
func sum(a []int) int {
	res := 0
	for _, v := range a {
		res += v
	}
	return res
}

func numberOfSubarrays(nums []int, k int) int {
	left, right := findKth(nums, 1), findKth(nums, k)
	if right == -1 {
		return 0
	}
	res := 0
	lc, rc := left+1, 1
	for right < len(nums) {
		for right < len(nums) && nums[right]%2 == 0 {
			right++
			rc++
		}
		res += lc * rc
		lc, rc = 1, 1
		left++
		for left < len(nums) && nums[left]%2 == 0 {
			left++
			lc++
		}
	}
	return res
}
func findKth(nums []int, k int) int {
	cnt := 0
	for idx, num := range nums {
		if num%2 == 1 {
			cnt++
		}
		if cnt == k {
			return idx
		}
	}
	return -1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	res1, res2 := inOrderTree(root1), inOrderTree(root2)
	return mergeSlice(res1, res2)
}
func inOrderTree(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		dfs(r.Left)
		res = append(res, r.Val)
		dfs(r.Right)
	}
	dfs(root)
	return res
}
func mergeSlice(a, b []int) []int {
	res := make([]int, 0)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	res = append(res, a[i:]...)
	res = append(res, b[j:]...)
	return res
}
