package main

func arithmeticTriplets(nums []int, diff int) int {
	m := make(map[int]int)
	for index, val := range nums {
		m[val] = index
	}
	ans := 0
	for i, val := range nums {
		idx1, ok1 := m[val-diff]
		idx2, ok2 := m[val+diff]
		if ok1 && ok2 && idx1 < i && idx2 > i {
			ans++
		}
	}
	return ans
}

func reachableNodes(n int, edges [][]int, restricted []int) int {
	out := make(map[int][]int)
	for _, val := range edges {
		out[val[0]] = append(out[val[0]], val[1])
		out[val[1]] = append(out[val[1]], val[0])
	}
	queue := []int{0}
	restrictedM := make(map[int]int)
	for _, val := range restricted {
		restrictedM[val] = 1
	}
	travel := make(map[int]int)
	ans := 0
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		ans++
		travel[cur] = 1
		for _, val := range out[cur] {
			if restrictedM[val] == 0 && travel[val] == 0 {
				travel[val] = 1
				queue = append(queue, val)
			}
		}
	}
	return ans
}

func validPartition(nums []int) bool {
	//这就是个很简单的动态规划
	n := len(nums)
	dp := make([]bool, n)
	dp[0] = false
	dp[1] = nums[0] == nums[1]
	if n == 2 {
		return dp[1]
	}
	var third func(i int) bool
	third = func(i int) bool {
		case1 := (nums[i] == nums[i-1] && nums[i] == nums[i-2])
		case2 := (nums[i] == nums[i-1]+1 && nums[i-1] == nums[i-2]+1)
		return case1 || case2
	}
	dp[2] = third(2)
	for i := 3; i < n; i++ {
		if nums[i] == nums[i-1] {
			dp[i] = dp[i] || dp[i-2]
		}
		if third(i) {
			dp[i] = dp[i] || dp[i-3]
		}
	}
	return dp[n-1]
}
