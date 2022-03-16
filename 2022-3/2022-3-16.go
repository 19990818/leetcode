package main

import "sort"

type node [][]int

func (m node) Len() int {
	return len(m)
}

func (m node) Less(i, j int) bool {
	return m[i][0] < m[j][0]
}

func (m node) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func reconstructQueue(people [][]int) [][]int {
	type info struct {
		curNum  int
		realNum int
		hight   int
	}
	queue := make([][]int, 0)
	ans := make([][]int, 0)
	temp := make([]info, 0)
	for _, val := range people {
		if val[1] == 0 {
			queue = append(queue, val)
		} else {
			temp = append(temp, info{val[1], val[1], val[0]})
		}
	}
	for len(queue) > 0 {
		sort.Sort(node(queue))
		cur := queue[0]
		queue = queue[1:]
		ans = append(ans, cur)
		subTemp := make([]info, 0)
		for _, val := range temp {
			if val.hight <= cur[0] {
				val.curNum--
				if val.curNum == 0 {
					queue = append(queue, []int{val.hight, val.realNum})
				} else {
					subTemp = append(subTemp, val)
				}
			} else {
				subTemp = append(subTemp, val)
			}
		}
		temp = subTemp
	}
	return ans
}

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	ans := 0
	dp := 0
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp++
			ans += dp
		} else {
			dp = 0
		}
	}
	return ans
}

func canPartition(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make(map[int]bool)
	dp[0] = true
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		temp := make(map[int]bool)
		for key, val := range dp {
			temp[key] = val
		}
		for key := range temp {
			//fmt.Println(key,nums[i])
			if key+nums[i] < target {
				dp[key+nums[i]] = true
			} else if key+nums[i] == target {
				return true
			}
		}
	}
	return false
}

func canPartition2(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[target] == 1
}
