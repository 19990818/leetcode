package main

import (
	"bytes"
	"math"
	"strings"
)

func partitionDisjoint(nums []int) int {
	maxArr, minArr := make([]int, len(nums)), make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			maxArr[0] = nums[i]
		} else {
			maxArr[i] = max(maxArr[i-1], nums[i])
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if i == len(nums)-1 {
			minArr[i] = nums[i]
		} else {
			minArr[i] = min(minArr[i+1], nums[i])
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		if maxArr[i] <= minArr[i+1] {
			return i + 1
		}
	}
	return -1
}

func carPooling(trips [][]int, capacity int) bool {
	in, out := make(map[int]int), make(map[int]int)
	for _, val := range trips {
		in[val[1]] += val[0]
		out[val[2]] += val[0]
	}
	sum := 0
	for i := 0; i < 1001; i++ {
		sum = sum + in[i] - out[i]
		if sum > capacity {
			return false
		}
	}
	return true
}

func crackSafe(n int, k int) string {
	var dfs func(node int)
	travel := make(map[int]int)
	var ans strings.Builder
	highest := int(math.Pow10(n - 1))
	dfs = func(node int) {
		for i := 0; i < k; i++ {
			if travel[node*10+i] == 0 {
				//针对为什么要在dfs后面在进行加入
				//因为从开始节点到重新回到结束节点的过程中
				//我们可能并没有走过所有的边，因此需要在中途
				//dfs 将没有走过的边走过去，因为一旦回到原处
				//再重新进行遍历，将增大遍历距离
				travel[node*10+i] = 1
				dfs((node*10 + i) % highest)
				ans.WriteByte(byte('0' + i))
			}
		}
	}
	dfs(0)
	ans.WriteString(string(bytes.Repeat([]byte{'0'}, n-1)))
	return ans.String()
}
