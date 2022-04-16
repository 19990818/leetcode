package main

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	//递归
	if head == nil {
		return true
	}
	if root == nil {
		return false
	}
	var dfs func(h *ListNode, r *TreeNode) bool
	dfs = func(h *ListNode, r *TreeNode) bool {
		if h == nil {
			return true
		}
		if r == nil {
			return false
		}
		if h.Val != r.Val {
			return false
		}
		return dfs(h.Next, r.Left) || dfs(h.Next, r.Right)
	}
	if dfs(head, root) {
		return true
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	stack = append(stack, 0)
	ans := make([]int, len(temperatures))
	for i := 1; i < len(temperatures); i++ {
		//出现高温
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			cur := stack[len(stack)-1]
			ans[cur] = i - cur
			stack = stack[0 : len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

func kClosest(points [][]int, k int) [][]int {
	type a struct {
		dis   int
		point []int
	}
	distance := make([]a, 0)
	for _, val := range points {
		distance = append(distance, a{val[0]*val[0] + val[1]*val[1], val})
	}
	var getK func(nums []a, k int)
	getK = func(nums []a, k int) {
		temp := nums[0]
		pior := nums[0].dis
		left, right := 0, len(nums)-1
		for left < right {
			for left < right && nums[right].dis >= pior {
				right--
			}
			nums[left] = nums[right]
			for left < right && nums[left].dis <= pior {
				left++
			}
			nums[right] = nums[left]
		}
		nums[left] = temp
		//fmt.Println("nums",nums,left)
		if left == k-1 {
			return
		}
		if left < k {
			getK(nums[left+1:], k-left-1)
			return
		}
		getK(nums[0:left], k)
	}
	getK(distance, k)
	// fmt.Println(distance)
	res := make([][]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, distance[i].point)
	}
	return res
}

func findClosestNumber(nums []int) int {
	a, b := math.MinInt32, math.MaxInt32
	for _, val := range nums {
		if val < 0 {
			a = max(a, val)
		} else if val > 0 {
			b = min(b, val)
		} else {
			return 0
		}
	}
	if b >= -a {
		return b
	}
	return a
}

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	if total < cost1 && total < cost2 {
		return 1
	}
	count1 := total / cost1
	ans := int64(0)
	for i := 0; i <= count1; i++ {
		cost2Sum := total - i*cost1
		ans += int64(cost2Sum)/int64(cost2) + 1
	}
	return ans
}

type ATM struct {
	m []int
}

func ConstructorATM() ATM {
	return ATM{make([]int, 5)}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := range this.m {
		this.m[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	ans := make([]int, 0)
	temp := append([]int{}, this.m...)
	money := []int{20, 50, 100, 200, 500}
	for i := 4; i >= 0; i-- {
		if amount >= money[i]*temp[i] {
			amount -= money[i] * temp[i]
			temp[i] = 0
		} else {
			t := amount / money[i]
			temp[i] -= t
			amount -= t * money[i]

		}
		//fmt.Println(temp[i])
	}
	//fmt.Println(this.m,temp)
	if amount == 0 {
		for i := range this.m {
			ans = append(ans, this.m[i]-temp[i])
		}
		this.m = temp
		return ans
	}
	return []int{-1}
}
