package main

import "reflect"

func minMutation(start string, end string, bank []string) int {
	//要么就是不同的个数要么是-1 没有这样的路径就是-1
	diff := 0
	for i := 0; i < len(start); i++ {
		if start[i] != end[i] {
			diff++
		}
	}
	queue := make([]string, 0)
	queue = append(queue, start)
	m := make(map[string]bool)
	m[start] = true
	ans := 0
	for len(queue) > 0 {
		temp := make([]string, 0)
		ans++
		for _, val := range queue {
			for _, val2 := range bank {
				if isOnlyOneByteDiffer(val, val2) {
					if m[val2] == false {
						m[val2] = true
						if val2 == end {
							return ans
						}
						temp = append(temp, val2)
					}
				}
			}
		}
		queue = temp
	}
	return -1
}
func isOnlyOneByteDiffer(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	ans := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			ans++
		}
	}
	return ans == 1
}

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	pFlag := make([]int, 26)
	for _, val := range p {
		pFlag[val-'a']++
	}
	ans := make([]int, 0)
	sFlag := make([]int, 26)
	for i := 0; i < len(p); i++ {
		sFlag[s[i]-'a']++
	}
	if reflect.DeepEqual(sFlag, pFlag) {
		ans = append(ans, 0)
	}
	for i := 1; i <= len(s)-len(p); i++ {
		sFlag[s[i-1]-'a']--
		sFlag[s[i-1+len(p)]-'a']++
		if reflect.DeepEqual(sFlag, pFlag) {
			ans = append(ans, i)
		}
	}
	return ans
}

//元素值可以再次作为下标得到值，重复对一个下标进行操作 说明此元素值出现两次
func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		idx := abs(nums[i])
		if nums[idx-1] > 0 {
			nums[idx-1] *= -1
		} else {
			res = append(res, idx)
		}
	}
	return res
}
