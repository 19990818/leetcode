package main

import (
	"sort"
	"strconv"
)

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	dp := make([]int, 10)
	for i := 0; i < 10; i++ {
		if i == 0 {
			dp[i] = 10
		} else if i == 1 {
			dp[i] = 9 * 9
		} else {
			dp[i] = dp[i-1] * (10 - i)
		}
	}
	ans := 0
	for i := 0; i < n && i < 10; i++ {
		ans += dp[i]
	}
	return ans
}

func canMeasureWater(jug1Capacity int, jug2Capacity int, targetCapacity int) bool {
	if targetCapacity > jug1Capacity+jug2Capacity {
		return false
	}
	for i := -jug2Capacity; i <= jug2Capacity; i++ {
		if (targetCapacity-i*jug1Capacity)%jug2Capacity == 0 {
			return true
		}
	}
	return false
}

//统计key之后的元素个数，得到最大元素个数大小
func mostFrequent(nums []int, key int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == key {
			m[nums[i+1]]++
		}
	}
	maxCount := 0
	ans := 0
	for k, val := range m {
		if val > maxCount {
			maxCount = val
			ans = k
		}
	}
	return ans
}

//排序问题，直接调用go sort 可知go sort的排序是稳定的
//需要自己定义接口实现，使用len,less,swap
func sortJumbled(mapping []int, nums []int) []int {
	a := arrNum{nums, mapping}
	sort.Sort(a)
	return a.arr
}
func fnum(mapping []int, num int) int {
	ans := 0
	s := strconv.Itoa(num)
	for i := 0; i < len(s); i++ {
		ans = ans*10 + mapping[s[i]-'0']
	}
	return ans
}

type arrNum struct {
	arr     []int
	mapping []int
}

func (m arrNum) Len() int {
	return len(m.arr)
}

func (m arrNum) Less(i, j int) bool {
	return fnum(m.mapping, m.arr[i]) < fnum(m.mapping, m.arr[j])
}

func (m arrNum) Swap(i, j int) {
	m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
}

//深度优先遍历得到祖先节点
//在之前的算法中判断是否元素存在，导致内存不足出错
//使用map标记节点是否被加入可有效减少内存消耗
func getAncestors(n int, edges [][]int) [][]int {
	m := make(map[int][]int)
	for _, val := range edges {
		m[val[1]] = append(m[val[1]], val[0])
	}
	ans := make([][]int, n)

	for i := 0; i < n; i++ {
		temp := make([]int, 0)
		tempArr := append([]int{}, m[i]...)
		visted := make(map[int]bool)
		for len(tempArr) > 0 {
			if !visted[tempArr[0]] {
				temp = append(temp, tempArr[0])
				visted[tempArr[0]] = true
			}
			for _, val := range m[tempArr[0]] {
				//fmt.Println(visted[val])
				if !visted[val] {
					tempArr = append(tempArr, val)
				}
			}
			tempArr = tempArr[1:]
			//fmt.Println(tempArr)
		}
		sort.Ints(temp)
		ans[i] = temp
	}
	return ans
}
