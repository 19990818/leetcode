package main

import (
	"math"
	"sort"
	"strings"
)

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	//fmt.Println(nums)
	left, right := 0, len(nums)-1
	mid := (right-left)>>1 + left
	for left < mid && nums[left] == nums[mid] {
		left++
	}
	for mid < right && nums[right] == nums[mid] {
		right--
	}
	if left >= mid && right <= mid {
		return nums[mid]
	}
	// fmt.Println(left,right)
	if left >= mid || nums[mid] > nums[left] {
		//左边是递增的
		return min(nums[left], findMin(nums[mid+1:]))
	}
	//右边是递增的
	return min(nums[mid], findMin(nums[0:mid]))
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	//基数排序和桶排序可以完成o(n)复杂度
	n := len(nums)
	minNum, maxNum := math.MaxInt64, math.MinInt64
	for _, val := range nums {
		minNum = min(minNum, val)
		maxNum = max(maxNum, val)
	}
	type pair struct{ mi, ma int }
	d := max(1, (maxNum-minNum)/(n-1))
	buckets := make([]pair, (maxNum-minNum)/d+1)
	for i := range buckets {
		buckets[i] = pair{-1, -1}
	}
	for _, val := range nums {
		pos := (val - minNum) / d
		if buckets[pos].ma == -1 {
			buckets[pos].ma = val
			buckets[pos].mi = val
		} else {
			buckets[pos].ma = max(buckets[pos].ma, val)
			buckets[pos].mi = min(buckets[pos].mi, val)
		}
	}
	pre := -1
	ans := 0
	for idx, val := range buckets {
		if val.mi == -1 {
			continue
		}
		if pre != -1 {
			ans = max(ans, val.mi-buckets[pre].ma)
		}
		pre = idx
	}
	return ans
}

func maximumGap2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	//基数排序
	n := len(nums)
	maxNum := math.MinInt64
	for _, val := range nums {
		maxNum = max(maxNum, val)
	}
	cur := 1
	for cur >= maxNum {
		cnt := make([]int, 10)
		for _, val := range nums {
			cnt[(val/cur)%10]++
		}
		for i := 1; i < 10; i++ {
			cnt[i] += cnt[i-1]
		}
		temp := make([]int, n)
		for i := n - 1; i >= 0; i-- {
			pos := (nums[i] / cur) % 10
			temp[cnt[pos]-1] = nums[i]
			cnt[pos]--
		}
		nums = temp
		cur *= 10
	}
	ans := 0
	for i := 1; i < n; i++ {
		ans = max(ans, nums[i]-nums[i-1])
	}
	return ans
}

func reorderedPowerOf2(n int) bool {
	dest := make([]int, 0)
	for i := 0; i < 32; i++ {
		dest = append(dest, 1<<i)
	}
	var getSingeNum func(a int) string
	getSingeNum = func(a int) string {
		nums := make([]int, 0)
		for a > 0 {
			nums = append(nums, a%10)
			a /= 10
		}
		sort.Ints(nums)
		var res strings.Builder
		for _, val := range nums {
			res.WriteByte(byte(val + '0'))
		}
		return res.String()
	}
	destStr := make(map[string]int)
	for _, val := range dest {
		destStr[getSingeNum(val)] = 1
	}
	return destStr[getSingeNum(n)] == 1
}
