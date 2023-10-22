package main

func minOperations0(nums []int, k int) int {
	res, cnt := 0, 0
	m := make(map[int]int)
	for i := len(nums) - 1; cnt < k && i >= 0; i-- {
		if m[nums[i]] == 0 && nums[i] <= k {
			cnt++
			m[nums[i]] = 1
		}
		res++
	}
	return res
}

func minOperations(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	res := 0
	for _, v := range m {
		if v < 2 {
			return -1
		}
		res += (v + 2) / 3
	}
	return res
}

func maxSubarrays(nums []int) int {
	// 只有得到0时候截断 与之后将结果减小
	sum := 0xfffff
	res := 0
	for _, v := range nums {
		sum &= v
		if sum == 0 {
			res++
			sum = 0xfffff
		}
	}
	return max(res, 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumTripletValue(nums []int) int64 {
	res, n := 0, len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				res = max(res, (nums[i]-nums[j])*nums[k])
			}
		}
	}
	return int64(res)
}

func maximumTripletValue2(nums []int) int64 {
	n := len(nums)
	pre := make([]int, n)
	suf := make([]int, n)
	pre[0] = 0
	for i := 1; i < n; i++ {
		pre[i] = max(pre[i-1], nums[i-1])
	}
	suf[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		suf[i] = max(suf[i+1], nums[i+1])
	}
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, (pre[i]-nums[i])*suf[i])
	}
	return int64(res)
}
