package main

import (
	"sort"
)

func sortPeople(names []string, heights []int) []string {
	type Pair struct {
		name string
		h    int
	}
	temp := make([]Pair, 0)
	for i := 0; i < len(names); i++ {
		temp = append(temp, Pair{names[i], heights[i]})
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i].h > temp[j].h
	})
	ans := []string{}
	for _, val := range temp {
		ans = append(ans, val.name)
	}
	return ans
}

func longestSubarray(nums []int) int {
	//进行按位与说明最长连续最大值
	maxV := 0
	for _, val := range nums {
		maxV = max(maxV, val)
	}
	cnt := 0
	ans := 0
	for _, val := range nums {
		if val == maxV {
			cnt++
		} else {
			ans = max(ans, cnt)
			cnt = 0
		}
	}
	ans = max(ans, cnt)
	return ans
}

func goodIndices(nums []int, k int) []int {
	// 每个位置i后面对应的是非递减 前面非递增
	inc, dec := make([]int, len(nums)), make([]int, len(nums))
	i, j := 0, 1
	dec[0] = 1
	for j < len(nums) {
		if nums[j] <= nums[j-1] {
			//fmt.Println(j,i)
			dec[j] = j - i + 1
		} else {
			i = j
			dec[j] = 1
		}
		j++
	}
	i, j = len(nums)-1, len(nums)-2
	inc[i] = 1
	for j >= 0 {
		if nums[j] <= nums[j+1] {
			inc[j] = i - j + 1
		} else {
			i = j
			inc[j] = 1
		}
		j--
	}
	//fmt.Println(inc,dec)
	ans := make([]int, 0)
	for i := k; i < len(nums)-k; i++ {
		if inc[i+1] >= k && dec[i-1] >= k {
			ans = append(ans, i)
		}
	}
	return ans
}

func robotSim(commands []int, obstacles [][]int) int {
	type pair struct {
		x int
		y int
	}
	m := make(map[pair]int)
	for _, val := range obstacles {
		m[pair{val[0], val[1]}] = 1
	}
	ans := 0
	status := 0
	curx, cury := 0, 0
	for _, command := range commands {
		if command == -1 {
			status = (status + 1) % 4
		} else if command == -2 {
			status = (status + 3) % 4
		} else {
			switch status {
			case 0:
				for i := 0; i < command && m[pair{curx, cury + 1}] == 0; i++ {
					cury++
				}
			case 1:
				for i := 0; i < command && m[pair{curx + 1, cury}] == 0; i++ {
					curx++
				}
			case 2:
				for i := 0; i < command && m[pair{curx, cury - 1}] == 0; i++ {
					cury--
				}
			case 3:
				for i := 0; i < command && m[pair{curx - 1, cury}] == 0; i++ {
					curx--
				}
			}
			//fmt.Println(curx,cury)
			ans = max(ans, curx*curx+cury*cury)
		}
	}
	return ans
}
