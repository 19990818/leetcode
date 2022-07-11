package main

import (
	"math"
	"sort"
)

func maxSlidingWindow(nums []int, k int) []int {
	// 单调栈 新来的大的会干掉之前的小的
	ans := make([]int, 0)
	stack := make([]int, 0)
	for i, val := range nums {
		//当窗口存在元素 元素下标小于等于i-k是过期的
		if len(stack) > 0 && i-k >= stack[0] {
			stack = stack[1:]
		}
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= val {
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, i)
		if i >= k-1 {
			ans = append(ans, nums[stack[0]])
		}
	}
	if len(ans) == 0 {
		ans = append(ans, nums[stack[0]])
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}
	if root.Val == 2 {
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	}
	return evaluateTree(root.Right) && evaluateTree(root.Left)
}

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)
	i, j := 0, 0
	//遍历每辆车子的顾客 j为最后没有上车的时间
	cnt := 0
	for i < len(buses) {
		cnt = 0
		for j < len(passengers) && passengers[j] <= buses[i] && cnt < capacity {
			j++
			cnt++
		}
		if j >= len(passengers) {
			break
		}
		i++
	}
	m := make(map[int]int)
	for _, val := range passengers {
		m[val] = 1
	}
	if cnt < capacity {
		for res := buses[len(buses)-1]; res > 0; res-- {
			if m[res] == 1 {
				continue
			}
			return res
		}
		return 0
	}
	for res := passengers[j-1] - 1; res > 0; res-- {
		if m[res] == 1 {
			continue
		}
		return res
	}
	return 0
}

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	//总是跳进自己误认为正确的地方无法自拔
	sub := make([]int, 0)
	sum := 0
	for i := range nums1 {
		if nums1[i] > nums2[i] {
			sub = append(sub, nums1[i]-nums2[i])
		} else {
			sub = append(sub, nums2[i]-nums1[i])
		}
		sum += sub[i]
	}
	if sum <= k1+k2 {
		return 0
	}
	sort.Ints(sub)
	n := len(sub)
	ans := int64(0)
	//使用优先队列超时 没有充分利用性质 时间消耗太大
	cnt := k1 + k2
	var pow2 func(a int) int64
	pow2 = func(a int) int64 {
		return int64(a * a)
	}
	//我们会将后面的都减到一样大
	pre := 0
	//fmt.Println(sub)
	for i := 0; i < n; i++ {
		if x := int(math.Ceil(float64(sum-cnt-pre) / float64(n-i))); x <= sub[i] {
			cnt2 := cnt - sum + pre + (n-i)*x
			for j := i; j < n; j++ {
				sub[j] = x
				if cnt2 > 0 {
					sub[j]--
					cnt2--
				}
			}
			break
		}
		pre += sub[i]
	}
	//fmt.Println(sub)
	for _, val := range sub {
		ans += pow2(val)
	}
	return ans
}
