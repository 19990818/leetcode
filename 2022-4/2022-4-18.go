package main

import "sort"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	m := len(r)
	ans := make([]bool, 0)
	for i := 0; i < m; i++ {
		if r[i] == l[i]+1 {
			ans = append(ans, true)
		} else {
			flag := true
			temp := append([]int{}, nums[l[i]:r[i]+1]...)
			sort.Ints(temp)
			for j := 1; j < len(temp)-1; j++ {
				if temp[j]-temp[j-1] != temp[j+1]-temp[j] {
					flag = false
					break
				}
			}
			ans = append(ans, flag)
		}
	}
	return ans
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	//需要一个数组用来记录每个人得到消息需要的时间
	time := make([]int, n)
	parentM := make(map[int][]int)
	for i := 0; i < len(manager); i++ {
		parentM[manager[i]] = append(parentM[manager[i]], i)
	}
	queue := make([]int, 0)
	queue = append(queue, headID)
	time[headID] = 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, val := range parentM[cur] {
			time[val] = time[cur] + informTime[cur]
			queue = append(queue, val)
		}
	}
	ans := 0
	for _, val := range time {
		ans = max(ans, val)
	}
	return ans
}

func numSubarrayProductLessThanK(nums []int, k int) int {
	//贪心 前缀后过大 需要对其进行缩小 当我们不符合条件时候 需要从左边减小乘积
	if k == 0 || k == 1 {
		return 0
	}
	ans := 0
	i := 0
	sum := 1
	for j := 0; j < len(nums); j++ {
		sum *= nums[j]
		for sum >= k {
			sum /= nums[i]
			i++
		}
		ans += j - i + 1
	}
	return ans
}
